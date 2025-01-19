package html

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"github.com/stroiman/go-dom/browser/dom"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type domParser struct{}

// Parses a HTML or XML from an [io.Reader] instance. The parsed nodes will
// have reference the window, e.g. letting events bubble to the window itself.
// The document pointer will be replaced by the created document.
//
// The document is updated using a pointer rather than returned as a value, as
// parseing process can e.g. execute script tags that require the document to
// be set on the window _before_ the script is executed.
func (p domParser) ParseReader(window Window, document *dom.Document, reader io.Reader) error {
	*document = newHTMLDocument(window)
	return parseIntoDocument(window, *document, reader)
}

func (p domParser) ParseFragment(
	ownerDocument dom.Document,
	reader io.Reader,
) (dom.DocumentFragment, error) {
	nodes, err := html.ParseFragment(reader, &html.Node{
		Type:     html.ElementNode,
		Data:     "body",
		DataAtom: atom.Body,
	})
	result := ownerDocument.CreateDocumentFragment()
	if err == nil {
		for _, child := range nodes {
			element := createElementFromNode(nil, ownerDocument, nil, child)
			result.AppendChild(element)
		}
	}
	return result, err
}

func NewDOMParser() domParser { return domParser{} }

type ElementSteps interface {
	AppendChild(parent dom.Node, child dom.Node) dom.Node
	Connected(w Window, n dom.Element)
}

type BaseRules struct{}

func (r BaseRules) AppendChild(parent dom.Node, child dom.Node) dom.Node {
	res, err := parent.AppendChild(child)
	if err != nil {
		panic(err)
	}
	return res
}

func (r BaseRules) Connected(w Window, n dom.Element) {}

type ScriptElementRules struct{ BaseRules }

func (r ScriptElementRules) Connected(win Window, node dom.Element) {
	var script string
	src, hasSrc := node.GetAttribute("src")
	slog.Debug("Process script tag")
	if !hasSrc {
		b := strings.Builder{}
		for _, child := range node.ChildNodes().All() {
			switch n := child.(type) {
			case dom.TextNode:
				b.WriteString(n.Text())
			}
		}
		script = b.String()
	} else {
		window := win.(*window)
		resp, err := window.httpClient.Get(src)
		if err != nil {
			panic(err)
		}
		if resp.StatusCode != 200 {
			panic("Bad response")
		}

		buf := bytes.NewBuffer([]byte{})
		buf.ReadFrom(resp.Body)
		script = string(buf.Bytes())

	}
	// TODO: Propagate error for better error handling
	err := win.Run(script)
	if err != nil {
		slog.Error("Error loading script")
	} else {
		slog.Debug("Script loaded/executed")
	}
}

type TemplateElementRules struct{ BaseRules }

func (TemplateElementRules) AppendChild(parent dom.Node, child dom.Node) dom.Node {
	template, ok := child.(HTMLTemplateElement)
	if !ok {
		panic("Parser error, applying tepmlate rules to non-template element")
	}
	parent.AppendChild(child)
	return template.Content()
}

var ElementMap = map[atom.Atom]ElementSteps{
	atom.Script:   ScriptElementRules{},
	atom.Template: TemplateElementRules{},
}

func parseIntoDocument(w Window, doc dom.Document, r io.Reader) error {
	node, err := html.Parse(r)
	if err != nil {
		return err
	}
	iterate(w, doc, doc, node)
	return nil
}

func cloneNode(n *html.Node) *html.Node {
	return &html.Node{
		Type:      n.Type,
		Data:      n.Data,
		DataAtom:  n.DataAtom,
		Namespace: n.Namespace,
		Attr:      n.Attr,
	}
}

// convertNS converts the namespace URI from x/net/html to the _right_
// namespace.
// SVG elements have namespace "svg"
func convertNS(ns string) string {
	switch ns {
	case "svg":
		return "http://www.w3.org/2000/svg"
	default:
		return ns
	}
}

func createElementFromNode(
	w Window,
	d dom.Document,
	parent dom.Node,
	source *html.Node,
) dom.Element {
	rules := ElementMap[source.DataAtom]
	if rules == nil {
		rules = BaseRules{}
	}
	var newNode dom.Node
	var newElm dom.Element
	if source.Namespace == "" {
		newElm = d.CreateElement(source.Data)
	} else {
		newElm = d.CreateElementNS(convertNS(source.Namespace), source.Data)
	}
	for _, a := range source.Attr {
		newElm.SetAttribute(a.Key, a.Val)
	}
	newNode = newElm
	if parent != nil {
		newNode = rules.AppendChild(parent, newElm)
	}
	iterate(w, d, newNode, source)
	// ?
	if rules != nil {
		if newElm.IsConnected() {
			if w != nil {
				rules.Connected(w, newElm)
			}
		}
	}
	return newElm
}

func iterate(w Window, d dom.Document, dest dom.Node, source *html.Node) {
	for child := range source.ChildNodes() {
		switch child.Type {
		case html.ElementNode:
			createElementFromNode(w, d, dest, child)
		case html.TextNode:
			dest.AppendChild(dom.NewTextNode(child.Data))
		case html.DoctypeNode:
			dest.AppendChild(dom.NewDocumentType(child.Data))
		case html.CommentNode:
			dest.AppendChild(dom.NewCommentNode(child.Data))
		default:
			panic(fmt.Sprintf("Node not yet supported: %v", child.Type))
		}
	}
}
