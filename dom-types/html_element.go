package dom_types

import (
	"strings"

	"github.com/stroiman/go-dom/interfaces"
)

func NewHTMLElement(tagName string) *Element {
	return NewElement(strings.ToUpper(tagName))
}

func NewHTMLUnknownElement(tagName string) *Element {
	return NewHTMLElement(strings.ToUpper(tagName))
}

func NewHTMLHtmlElement(doc interfaces.Document) *Element {
	return NewHTMLElement("HTML")
}
