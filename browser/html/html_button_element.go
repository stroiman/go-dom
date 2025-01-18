package html

type htmlButtonElement struct{ *htmlElement }

func NewHTMLButtonElement(ownerDocument HTMLDocument) HTMLElement {
	result := &htmlButtonElement{newHTMLElement("button", ownerDocument)}
	result.SetSelf(result)
	return result
}

func (e *htmlButtonElement) Click() bool {
	ok := e.htmlElement.Click()
	if ok {
		e.trySubmitForm()
	}
	return ok
}
func (e *htmlButtonElement) trySubmitForm() {
	t, _ := e.GetAttribute("type")
	if t != "submit" {
		return
	}
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
		form.Submit()
	}
}
