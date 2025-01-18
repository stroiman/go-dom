package html

import "strings"

type HTMLInputElement interface {
	HTMLElement
	Type() string
	SetType(value string)
}

type htmlInputElement struct{ *htmlElement }

func NewHTMLInputElement(ownerDocument HTMLDocument) HTMLInputElement {
	result := &htmlInputElement{newHTMLElement("input", ownerDocument)}
	result.SetSelf(result)
	return result
}

func (e *htmlInputElement) Type() string {
	t, _ := e.GetAttribute("type")
	return strings.ToLower(t)
}
func (e *htmlInputElement) SetType(val string) {
	e.SetAttribute("type", val)
}
