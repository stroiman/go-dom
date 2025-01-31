package html

import (
	"strings"

	. "github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/internal/dom"
)

type HTMLTemplateElement interface {
	HTMLElement
	Content() DocumentFragment
}

type htmlTemplateElement struct {
	*htmlElement
	content DocumentFragment
}

func NewHTMLTemplateElement(ownerDocument HTMLDocument) HTMLTemplateElement {
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
