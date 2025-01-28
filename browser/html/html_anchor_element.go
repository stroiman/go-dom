package html

import (
	"github.com/stroiman/go-dom/browser/dom"
)

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

func (e *htmlAnchorElement) setUrl(f func(dom.URL, string), val string) {
	f(e.URL, val)
	e.HTMLElement.SetAttribute("href", e.URL.Href())
}

func (e *htmlAnchorElement) SetHref(val string) {
	e.setUrl(dom.URL.SetHref, val)
}

func (e *htmlAnchorElement) SetProtocol(val string) { e.setUrl(dom.URL.SetProtocol, val) }
func (e *htmlAnchorElement) SetUsername(val string) { e.setUrl(dom.URL.SetUsername, val) }
func (e *htmlAnchorElement) SetPassword(val string) { e.setUrl(dom.URL.SetPassword, val) }
func (e *htmlAnchorElement) SetHost(val string)     { e.setUrl(dom.URL.SetHost, val) }
func (e *htmlAnchorElement) SetHostname(val string) { e.setUrl(dom.URL.SetHostname, val) }
func (e *htmlAnchorElement) SetPort(val string)     { e.setUrl(dom.URL.SetPort, val) }
func (e *htmlAnchorElement) SetPathname(val string) { e.setUrl(dom.URL.SetPathname, val) }
func (e *htmlAnchorElement) SetSearch(val string)   { e.setUrl(dom.URL.SetSearch, val) }
func (e *htmlAnchorElement) SetHash(val string)     { e.setUrl(dom.URL.SetHash, val) }
