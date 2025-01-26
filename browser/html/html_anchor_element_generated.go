// This file is generated. Do not edit.

package html

type HTMLAnchorElement interface {
	HTMLElement
	Target() string
	SetTarget(string)
	Download() string
	SetDownload(string)
	Ping() string
	SetPing(string)
	Rel() string
	SetRel(string)
	RelList() string
	Hreflang() string
	SetHreflang(string)
	Type() string
	SetType(string)
	Text() string
	SetText(string)
	ReferrerPolicy() string
	SetReferrerPolicy(string)
	Href() string
	SetHref(string)
	Origin() string
	Protocol() string
	SetProtocol(string)
	Username() string
	SetUsername(string)
	Password() string
	SetPassword(string)
	Host() string
	SetHost(string)
	Hostname() string
	SetHostname(string)
	Port() string
	SetPort(string)
	Pathname() string
	SetPathname(string)
	Search() string
	SetSearch(string)
	Hash() string
	SetHash(string)
}

type htmlAnchorElement struct {
	HTMLElement
}

func NewHTMLAnchorElement(ownerDoc HTMLDocument) HTMLAnchorElement {
	result := &htmlAnchorElement{NewHTMLElement("a", ownerDoc)}
	result.SetSelf(result)
	return result
}

func (e *htmlAnchorElement) Target() string {
	result, _ := e.GetAttribute("target")
	return result
}

func (e *htmlAnchorElement) SetTarget(val string) {
	e.SetAttribute("target", val)
}
func (e *htmlAnchorElement) Download() string {
	result, _ := e.GetAttribute("download")
	return result
}

func (e *htmlAnchorElement) SetDownload(val string) {
	e.SetAttribute("download", val)
}
func (e *htmlAnchorElement) Ping() string {
	result, _ := e.GetAttribute("ping")
	return result
}

func (e *htmlAnchorElement) SetPing(val string) {
	e.SetAttribute("ping", val)
}
func (e *htmlAnchorElement) Rel() string {
	result, _ := e.GetAttribute("rel")
	return result
}

func (e *htmlAnchorElement) SetRel(val string) {
	e.SetAttribute("rel", val)
}
func (e *htmlAnchorElement) RelList() string {
	result, _ := e.GetAttribute("relList")
	return result
}
func (e *htmlAnchorElement) Hreflang() string {
	result, _ := e.GetAttribute("hreflang")
	return result
}

func (e *htmlAnchorElement) SetHreflang(val string) {
	e.SetAttribute("hreflang", val)
}
func (e *htmlAnchorElement) Type() string {
	result, _ := e.GetAttribute("type")
	return result
}

func (e *htmlAnchorElement) SetType(val string) {
	e.SetAttribute("type", val)
}
func (e *htmlAnchorElement) Text() string {
	result, _ := e.GetAttribute("text")
	return result
}

func (e *htmlAnchorElement) SetText(val string) {
	e.SetAttribute("text", val)
}
func (e *htmlAnchorElement) ReferrerPolicy() string {
	result, _ := e.GetAttribute("referrerPolicy")
	return result
}

func (e *htmlAnchorElement) SetReferrerPolicy(val string) {
	e.SetAttribute("referrerPolicy", val)
}
