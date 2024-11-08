package dom_types

import (
	"github.com/stroiman/go-dom/interfaces"
)

type HTMLHtmlElement struct {
	*HTMLElement
	OwningDocument interfaces.Document
}

func NewHTMLHtmlElement(doc interfaces.Document) *HTMLHtmlElement {
	return &HTMLHtmlElement{NewHTMLElement("HTML"), doc}
}

func (e *HTMLHtmlElement) Append(child interfaces.Element) interfaces.Element {
	if child.NodeName() == "BODY" && e == e.OwningDocument.DocumentElement() {
		e.OwningDocument.SetBody(child)
	}
	return child
}
