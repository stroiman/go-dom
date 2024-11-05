package dom_types

type HTMLElement struct {
	Element
}

func NewHTMLElement(tagName string) HTMLElement { return HTMLElement{NewElement(tagName)} }
