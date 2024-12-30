package dom

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type DOMParser interface {
	// Parses a HTML or XML from an [io.Reader] instance. The parsed nodes will
	// have reference the window, e.g. letting events bubble to the window itself.
	// The document pointer will be replaced by the created document.
	//
	// The document is updated using a pointer rather than returned as a value, as
	// parseing process can e.g. execute script tags that require the document to
	// be set on the window _before_ the script is executed.
	ParseReader(window Window, document *Document, reader io.Reader) error
}

// TODO: Remove
type domParser struct{}

func (p domParser) ParseReader(window Window, document *Document, reader io.Reader) error {
	*document = NewDocument(window)
	return parseIntoDocument(window, *document, reader)
}

func NewDOMParser() DOMParser { return domParser{} }

type ElementSteps interface {
	AppendChild(parent Node, child Node) Node
	Connected(w Window, n Element)
}

type BaseRules struct{}

func (r BaseRules) AppendChild(parent Node, child Node) Node {
	return parent.AppendChild(child)
}

func (r BaseRules) Connected(w Window, n Element) {}

type ScriptElementRules struct{ BaseRules }

func (r ScriptElementRules) Connected(win Window, node Element) {
	var script string
	src := node.GetAttribute("src")
	slog.Debug("Process script tag")
	if src == "" {
		b := strings.Builder{}
		for _, child := range node.ChildNodes().All() {
			switch n := child.(type) {
			case TextNode:
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

func (TemplateElementRules) AppendChild(parent Node, child Node) Node {
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

func parseIntoDocument(w Window, doc Document, r io.Reader) error {
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

func createElementFromNode(w Window, d Document, parent Node, source *html.Node) Element {
	rules := ElementMap[source.DataAtom]
	if rules == nil {
		rules = BaseRules{}
	}
	var newNode Node
	newElm := d.CreateElement(source.Data)
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
		if newElm.Connected() {
			if w != nil {
				rules.Connected(w, newElm)
			}
		}
	}
	return newElm
}

func iterate(w Window, d Document, dest Node, source *html.Node) {
	for child := range source.ChildNodes() {
		switch child.Type {
		case html.ElementNode:
			createElementFromNode(w, d, dest, child)
		case html.TextNode:
			NodeHelper{dest}.AppendChild(NewTextNode(cloneNode(child), child.Data))
		default:
			panic(fmt.Sprintf("Node not yet supported: %v", child))
		}
	}
}
