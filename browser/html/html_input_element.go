package html

import (
	"strings"
)

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
	if t == "" {
		return "text"
	}
	return strings.ToLower(t)
}

func (e *htmlInputElement) SetType(val string) {
	e.SetAttribute("type", val)
}

func (e *htmlInputElement) Click() bool {
	ok := e.htmlElement.Click()
	if ok {
		if e.Type() == "submit" {
			e.trySubmitForm()
		}
	}
	return ok
}

func (e *htmlInputElement) trySubmitForm() {
	var form HTMLFormElement
	parent := e.Parent()
	for {
		if parent == nil {
			break
		}
		if f, ok := parent.(HTMLFormElement); ok {
			form = f
			break
		}
		parent = parent.Parent()
	}
	if form != nil {
		form.RequestSubmit(e)
	}
}
