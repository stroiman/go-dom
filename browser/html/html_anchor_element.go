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
		window := e.getHTMLDocument().getWindow()
		newUrl := window.resolveHref(href)
		window.Navigate(newUrl.Href())
	}
	return result
}
