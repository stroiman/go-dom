package browser

import (
	"golang.org/x/net/html"
)

type HTMLElement interface {
	Element
}

func NewHTMLElement(node *html.Node) Element {
	return newElementFromNode(node)
}

func NewHTMLUnknownElement(node *html.Node) Element {
	return NewHTMLElement(node)
}

func NewHTMLHtmlElement(node *html.Node) Element {
	return NewHTMLElement(node)
}
