package dom

import (
	"strings"

	"golang.org/x/net/html"
)

type HTMLElement interface {
	Element
}

type htmlElement struct {
	*element
}

func NewHTMLElement(tagName string, ownerDocument Document) HTMLElement {
	result := &htmlElement{newElement(tagName, ownerDocument)}
	result.SetSelf(result)
	return result
}

func newHTMLElement(tagName string, ownerDocument Document) *htmlElement {
	result := &htmlElement{newElement(tagName, ownerDocument)}
	result.SetSelf(result)
	return result
}

type HTMLTemplateElement interface {
	HTMLElement
	Content() DocumentFragment
}

type htmlTemplateElement struct {
	*htmlElement
	content DocumentFragment
}

func NewHTMLTemplateElement(ownerDocument Document) HTMLTemplateElement {
	result := &htmlTemplateElement{
		newHTMLElement("template", ownerDocument),
		NewDocumentFragment(ownerDocument),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlTemplateElement) Content() DocumentFragment { return e.content }

func (e *htmlTemplateElement) RenderChildren(builder *strings.Builder) {
	if renderer, ok := e.content.(ChildrenRenderer); ok {
		renderer.RenderChildren(builder)
	}
}

func (e *htmlTemplateElement) createHtmlNode() *html.Node {
	node := e.htmlElement.createHtmlNode()
	for _, child := range e.content.nodes() {
		node.AppendChild(NodeIterator{child}.toHtmlNode(nil))
	}
	return node
}
