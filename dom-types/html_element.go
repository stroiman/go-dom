package dom_types

import (
	"strings"
)

func NewHTMLElement(tagName string) Element {
	return NewElement(strings.ToUpper(tagName))
}

func NewHTMLUnknownElement(tagName string) Element {
	return NewHTMLElement(strings.ToUpper(tagName))
}

func NewHTMLHtmlElement(doc *document) Element {
	return NewHTMLElement("HTML")
}
