package dom_types

import "strings"

type HTMLElement struct {
	Element
}

func NewHTMLElement(tagName string) HTMLElement {
	return HTMLElement{NewElement(strings.ToUpper(tagName))}
}
