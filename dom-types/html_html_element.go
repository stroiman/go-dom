package dom_types

type HTMLHtmlElement struct {
	HTMLElement
}

func NewHTMLHtmlElement() HTMLHtmlElement {
	return HTMLHtmlElement{NewHTMLElement("HTML")}
}
