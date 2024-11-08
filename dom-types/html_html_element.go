package dom_types

import (
	"fmt"

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
	fmt.Printf("NAME %s %v", child.NodeName(), e == e.OwningDocument.DocumentElement())
	if child.NodeName() == "BODY" && e == e.OwningDocument.DocumentElement() {
		fmt.Println("SET")
		e.OwningDocument.SetBody(child)
	}
	return child
}
