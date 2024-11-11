package browser

import (
	"golang.org/x/net/html"
)

func NewHTMLElement(node *html.Node) Element {
	return NewElement(node.Data, node)
}

func NewHTMLUnknownElement(node *html.Node) Element {
	return NewHTMLElement(node)
}

func NewHTMLHtmlElement(node *html.Node) Element {
	return NewHTMLElement(node)
}
