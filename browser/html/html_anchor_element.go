package html

import "github.com/stroiman/go-dom/browser/dom"

type htmlAnchorElement struct {
	HTMLElement
	dom.URL
}

func NewHTMLAnchorElement(ownerDoc HTMLDocument) HTMLAnchorElement {
	result := &htmlAnchorElement{
		HTMLElement: NewHTMLElement("a", ownerDoc),
		URL:         dom.ParseURL(ownerDoc.getWindow().History().window.baseLocation),
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

func (e *htmlAnchorElement) SetAttribute(name string, val string) {
	win := e.getWindow().History().window
	e.HTMLElement.SetAttribute(name, val)
	if name == "href" {
		e.URL = dom.ParseURLBase(val, win.baseLocation)
	}
}
