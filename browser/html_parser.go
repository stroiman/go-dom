package browser

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

func parseStream(r io.Reader) Document {
	node, err := html.Parse(r)
	if err != nil {
		panic(err)
	}
	response := NewDocument()
	iterate(response, response, node)
	return response

}

func iterate(d Document, dest Node, source *html.Node) {
	for child := range source.ChildNodes() {
		switch child.Type {
		case html.ElementNode:
			newElm := d.wrapElement(child)
			dest.AppendChild(newElm)
			iterate(d, newElm, child)
		}
	}
}

func ParseHtmlStream(s io.Reader) Document {
	return parseStream(s)
}

func ParseHtmlString(s string) Document {
	return ParseHtmlStream(strings.NewReader(s))
}
