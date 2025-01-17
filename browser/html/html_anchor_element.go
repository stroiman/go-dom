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
	result := e.HTMLElement.Click()
	if href, found := e.GetAttribute("href"); found && result {
		e.getHTMLDocument().getWindow().Navigate(href)
	}
	return result
}
