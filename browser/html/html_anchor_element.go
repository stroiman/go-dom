package html

import "github.com/stroiman/go-dom/browser/dom"

func (e *htmlAnchorElement) Click() bool {
	result := e.HTMLElement.Click()
	if href, found := e.GetAttribute("href"); found && result {
		window := e.getHTMLDocument().getWindow()
		newUrl := window.resolveHref(href)
		window.Navigate(newUrl.Href())
	}
	return result
}
func (e *htmlAnchorElement) Href() string {
	return dom.ParseURL(e.getWindow().History().window.baseLocation).Href()
}
func (e *htmlAnchorElement) SetHref(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Origin() string {
	return dom.ParseURL(e.getWindow().History().window.baseLocation).Origin()
}
func (e *htmlAnchorElement) Protocol() string {
	return dom.ParseURL(e.getWindow().History().window.baseLocation).Protocol()
}
func (e *htmlAnchorElement) SetProtocol(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Username() string   { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetUsername(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Password() string   { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) SetPassword(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Host() string {
	return dom.ParseURL(e.getWindow().History().window.baseLocation).Host()
}
func (e *htmlAnchorElement) SetHost(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Hostname() string {
	return dom.ParseURL(e.getWindow().History().window.baseLocation).Hostname()
}
func (e *htmlAnchorElement) SetHostname(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Port() string {
	return dom.ParseURL(e.getWindow().History().window.baseLocation).Port()
}
func (e *htmlAnchorElement) SetPort(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Pathname() string {
	return dom.ParseURL(e.getWindow().History().window.baseLocation).Pathname()
}
func (e *htmlAnchorElement) SetPathname(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Search() string {
	return dom.ParseURL(e.getWindow().History().window.baseLocation).Search()
}
func (e *htmlAnchorElement) SetSearch(string) { panic("Not implemented, sorry") }
func (e *htmlAnchorElement) Hash() string {
	return dom.ParseURL(e.getWindow().History().window.baseLocation).Hash()
}
func (e *htmlAnchorElement) SetHash(string) { panic("Not implemented, sorry") }
