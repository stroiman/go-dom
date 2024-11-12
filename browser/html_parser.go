package browser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/net/html/atom"
)

type ElementSteps interface {
	Connected(w Window, n Node)
}

type ScriptElementRules struct{}

func (r ScriptElementRules) Connected(win Window, node Node) {
	b := strings.Builder{}
	for child := range node.wrappedNode().ChildNodes() {
		if child.Type == html.TextNode {
			b.WriteString(child.Data)
		}
	}
	win.Eval(b.String())
}

var ElementMap = map[atom.Atom]ElementSteps{
	atom.Script: ScriptElementRules{},
}

func parseStream(w *window, r io.Reader) Document {
	node, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	doc := newDocument(node)
	if w != nil {
		w.document = doc
	}
	iterate(w, doc, doc, node)
	return doc
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

func iterate(w Window, d Document, dest Node, source *html.Node) {
	for child := range source.ChildNodes() {
		switch child.Type {
		case html.ElementNode:
			rules := ElementMap[child.DataAtom]
			newElm := d.createElement(cloneNode(child))
			dest.AppendChild(newElm)
			iterate(w, d, newElm, child)
			// ?
			// if newElm.Connected() {
			if rules != nil {
				if w != nil {
					rules.Connected(w, newElm)
				}
			}
			// }
		default:
			clone := newNode(cloneNode(child))
			dest.AppendChild(&clone)
		}
	}
}

func ParseHtmlStream(s io.Reader) Document {
	return parseStream(nil, s)
}

func ParseHtmlString(s string) Document {
	return ParseHtmlStream(strings.NewReader(s))
}
