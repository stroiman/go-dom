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

func parseStream(w *window, doc Document, r io.Reader) Document {
	node, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	if doc == nil {
		doc = NewDocument()
	}
	if w != nil {
		w.document = doc
	}
	iterate(w, doc, doc, node)
	return doc
}

func iterate(w Window, d Document, dest Node, source *html.Node) {
	for child := range source.ChildNodes() {
		switch child.Type {
		case html.ElementNode:
			rules := ElementMap[child.DataAtom]
			newElm := d.wrapElement(child)
			dest.AppendChild(newElm)
			// if newElm.Connected() {
			if rules != nil {
				if w != nil {
					rules.Connected(w, newElm)
				}
			}
			// }
			iterate(w, d, newElm, child)
		}
	}
}

func ParseHtmlStream(s io.Reader) Document {
	return parseStream(nil, nil, s)
}

func ParseHtmlString(s string) Document {
	return ParseHtmlStream(strings.NewReader(s))
}
