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

func (e *htmlAnchorElement) SetHref(string)     { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetProtocol(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Username() string   { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetUsername(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Password() string   { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetPassword(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetHost(string)     { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetHostname(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetPort(string)     { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetPathname(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetSearch(string)   { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetHash(string)     { panic("Not implemented, sorry") }
