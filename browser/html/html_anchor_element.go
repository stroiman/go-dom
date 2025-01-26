package html

func (e *htmlAnchorElement) Click() bool {
	result := e.HTMLElement.Click()
	if href, found := e.GetAttribute("href"); found && result {
		window := e.getHTMLDocument().getWindow()
		newUrl := window.resolveHref(href)
		window.Navigate(newUrl.Href())
	}
	return result
}
