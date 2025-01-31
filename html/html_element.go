package html

import (
	"strings"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/dom"
)

type HTMLElement interface {
	Element
	Renderer
	ChildrenRenderer
	getHTMLDocument() HTMLDocument
	getWindow() Window
}

type htmlElement struct {
	Element
	Renderer
	ChildrenRenderer
	htmlDocument HTMLDocument
}

func NewHTMLElement(tagName string, ownerDocument HTMLDocument) HTMLElement {
	return newHTMLElement(tagName, ownerDocument)
}

func newHTMLElement(tagName string, ownerDocument HTMLDocument) *htmlElement {
	element := NewElement(tagName, ownerDocument)
	renderer, _ := element.(Renderer)
	childrenRenderer, _ := element.(ChildrenRenderer)
	result := &htmlElement{element, renderer, childrenRenderer, ownerDocument}
	result.SetSelf(result)
	return result
}

func (e *htmlElement) getHTMLDocument() HTMLDocument { return e.htmlDocument }

func (e *htmlElement) getWindow() Window { return e.getHTMLDocument().getWindow() }

func (e *htmlElement) TagName() string {
	return strings.ToUpper(e.Element.TagName())
}

func (e *htmlElement) CloneNode(deep bool) dom.Node {
	res := e.htmlDocument.CreateElement(e.TagName())
	for a := range e.Attributes().All() {
		res.SetAttribute(a.Name(), a.Value())
	}
	return res
}
