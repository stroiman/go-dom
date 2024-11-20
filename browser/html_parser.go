package browser

import (
	"bytes"
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
	var script string
	src := node.GetAttribute("src")
	if src == "" {
		b := strings.Builder{}
		for child := range node.wrappedNode().ChildNodes() {
			if child.Type == html.TextNode {
				b.WriteString(child.Data)
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
		fmt.Println("Error loading script", src, err)
	} else {
		fmt.Println("Script loaded/executed")
	}

}

var ElementMap = map[atom.Atom]ElementSteps{
	atom.Script: ScriptElementRules{},
}

func parseStream(w *window, r io.Reader) Document {
	node, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	doc := newDocument(cloneNode(node))
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
		case html.TextNode:
			dest.AppendChild(NewTextNode(cloneNode(child), child.Data))
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
