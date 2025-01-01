package html

type htmlAnchorElement struct {
	HTMLElement
}

func NewHTMLAnchorElement(ownerDocument HTMLDocument) HTMLElement {
	result := &htmlAnchorElement{
		NewHTMLElement("a", ownerDocument),
	}
	result.SetSelf(result)
	return result
}

func (e *htmlAnchorElement) Click() bool {
	// result := e.HTMLElement.Click()
	if href := e.GetAttribute("href"); href != "" {
		e.getHTMLDocument().getWindow().Navigate(e.GetAttribute("href"))
	}
	return true
}
