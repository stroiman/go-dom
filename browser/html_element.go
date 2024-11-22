package browser

import (
	"golang.org/x/net/html"
)

type HTMLElement interface {
	Element
}

func NewHTMLElement(node *html.Node, ownerDocument Document) Element {
	return newElementFromNode(node, ownerDocument)
}

func NewHTMLUnknownElement(node *html.Node, ownerDocument Document) Element {
	return NewHTMLElement(node, ownerDocument)
}

func NewHTMLHtmlElement(node *html.Node, ownerDocument Document) Element {
	return NewHTMLElement(node, ownerDocument)
}
