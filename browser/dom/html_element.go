package dom

import (
	"fmt"
	"strings"

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

type HTMLTemplateElement interface {
	HTMLElement
	Content() DocumentFragment
}

type htmlTemplateElement struct {
	element
	content DocumentFragment
}

func NewHTMLTemplateElement(node *html.Node, ownerDocument Document) HTMLTemplateElement {
	return &htmlTemplateElement{
		*(newElementFromNode(node, ownerDocument)),
		NewDocumentFragment(ownerDocument),
	}
}

func (e *htmlTemplateElement) Content() DocumentFragment { return e.content }

func (e *htmlTemplateElement) createHtmlNode() *html.Node {
	fmt.Println("FOO")
	node := e.element.createHtmlNode()
	for _, child := range e.content.nodes() {
		node.AppendChild(NodeIterator{child}.toHtmlNode(nil))
	}
	return node
}

func (e *htmlTemplateElement) OuterHTML() string {
	writer := &strings.Builder{}
	html.Render(writer, toHtmlNode(e))
	return string(writer.String())
}

func (e *htmlTemplateElement) InnerHTML() string {
	writer := &strings.Builder{}
	for _, child := range e.ChildNodes().All() {
		html.Render(writer, toHtmlNode(child))
	}
	return string(writer.String())
}
