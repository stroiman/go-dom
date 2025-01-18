package html

import "strings"

type HTMLButtonElement interface {
	HTMLElement
	Type() string
	SetType(val string)
}

type htmlButtonElement struct{ *htmlElement }

func NewHTMLButtonElement(ownerDocument HTMLDocument) HTMLButtonElement {
	result := &htmlButtonElement{newHTMLElement("button", ownerDocument)}
	result.SetSelf(result)
	return result
}

func (e *htmlButtonElement) Click() bool {
	ok := e.htmlElement.Click()
	if ok {
		if e.Type() == "submit" {
			e.trySubmitForm()
		}
	}
	return ok
}

func (e *htmlButtonElement) trySubmitForm() {
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

func (e *htmlButtonElement) Type() string {
	t, _ := e.GetAttribute("type")
	l := strings.ToLower(t)
	switch l {
	case "button":
		return l
	case "reset":
		return l
	case "submit":
		return l
	case "menu":
		return l
	}
	return "submit"
}

func (e *htmlButtonElement) SetType(val string) {
	e.SetAttribute("type", val)
}
