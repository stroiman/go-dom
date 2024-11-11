package browser

import (
	"fmt"
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

func parseStream(w Window, doc Document, r io.Reader) Document {
	node, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	if doc == nil {
		doc = NewDocument()
	}
	iterate(w, doc, doc, node)
	return doc
}

func iterate(w Window, d Document, dest Node, source *html.Node) {
	fmt.Println("Before", d.ChildNodes())
	for child := range source.ChildNodes() {
		switch child.Type {
		case html.ElementNode:
			rules := ElementMap[child.DataAtom]
			fmt.Println("Type", child.Data, rules)
			newElm := d.wrapElement(child)
			dest.AppendChild(newElm)
			fmt.Println("Dest child nodes", dest.ChildNodes())
			// if newElm.Connected() {
			if rules != nil {
				fmt.Println("Connected!")
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
