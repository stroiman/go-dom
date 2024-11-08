package dom_types

import "strings"

type HTMLElement struct {
	*Element
}

type HTMLUnknownElement struct {
	*HTMLElement
}

func NewHTMLElement(tagName string) *HTMLElement {
	return &HTMLElement{NewElement(strings.ToUpper(tagName))}
}

func NewHTMLUnknownElement(tagName string) *HTMLUnknownElement {
	return &HTMLUnknownElement{NewHTMLElement(strings.ToUpper(tagName))}
}
