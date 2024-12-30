package dom

import (
	"strings"

	"golang.org/x/net/html"
)

type HTMLElement interface {
	Element
}

func NewHTMLElement(tagName string, ownerDocument Document) Element {
	return NewElement(tagName, ownerDocument)
}

type HTMLTemplateElement interface {
	HTMLElement
	Content() DocumentFragment
}

type htmlTemplateElement struct {
	HTMLElement
	content DocumentFragment
}

func NewHTMLTemplateElement(ownerDocument Document) HTMLTemplateElement {
	return &htmlTemplateElement{
		NewHTMLElement("template", ownerDocument),
		NewDocumentFragment(ownerDocument),
	}
}

func (e *htmlTemplateElement) Content() DocumentFragment { return e.content }

func (e *htmlTemplateElement) createHtmlNode() *html.Node {
	node := e.HTMLElement.createHtmlNode()
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
