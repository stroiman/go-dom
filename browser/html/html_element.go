package html

import (
	"strings"

	. "github.com/stroiman/go-dom/browser/dom"
)

type HTMLElement interface {
	Element
	getHTMLDocument() HTMLDocument
}

type htmlElement struct {
	Element
	Renderer
	ChildrenRenderer
	htmlDocument HTMLDocument
}

func NewHTMLElement(tagName string, ownerDocument HTMLDocument) HTMLElement {
	element := NewElement(tagName, ownerDocument)
	renderer, _ := element.(Renderer)
	childrenRenderer, _ := element.(ChildrenRenderer)
	result := &htmlElement{element, renderer, childrenRenderer, ownerDocument}
	result.SetSelf(result)
	return result
}

func (e *htmlElement) getHTMLDocument() HTMLDocument { return e.htmlDocument }

type HTMLTemplateElement interface {
	HTMLElement
	Content() DocumentFragment
}

type htmlTemplateElement struct {
	HTMLElement
	content DocumentFragment
}

func NewHTMLTemplateElement(ownerDocument HTMLDocument) HTMLTemplateElement {
	result := &htmlTemplateElement{
		NewHTMLElement("template", ownerDocument),
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
