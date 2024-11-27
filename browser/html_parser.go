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
	Connected(w Window, n Element)
}

type ScriptElementRules struct{}

func (r ScriptElementRules) Connected(win Window, node Element) {
	var script string
	src := node.GetAttribute("src")
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
	doc := NewDocument(w)
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

func createElementFromNode(w Window, d Document, parent Node, source *html.Node) Element {
	rules := ElementMap[source.DataAtom]
	newElm := d.createElement(cloneNode(source))
	if parent != nil {
		parent.AppendChild(newElm)
	}
	iterate(w, d, newElm, source)
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

func ParseHtmlStream(s io.Reader) Document {
	return parseStream(nil, s)
}

func ParseHtmlString(s string) Document {
	return ParseHtmlStream(strings.NewReader(s))
}
