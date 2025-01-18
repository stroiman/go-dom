package html

type htmlButtonElement struct{ *htmlElement }

func NewHTMLButtonElement(ownerDocument HTMLDocument) HTMLElement {
	result := &htmlButtonElement{newHTMLElement("button", ownerDocument)}
	result.SetSelf(result)
	return result
}

func (e *htmlButtonElement) Click() bool {
	var form HTMLFormElement
	for parent := e.Parent(); ; parent = parent.Parent() {
		if f, ok := parent.(HTMLFormElement); ok {
			form = f
			break
		}
	}
	ok := e.htmlElement.Click()
	if ok {
		form.Submit()
	}
	return ok
}
