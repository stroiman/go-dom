package browser

import (
	netURL "net/url"
)

type URL interface {
	GetHash() string
	GetHost() string
	GetHostname() string
	GetHref() string
	// TODO
	// SetHref(href string)
	GetOrigin() string
	GetPathname() string
	GetPort() string
	GetProtocol() string
	GetSearch() string
}

type url struct {
	// Cannot be named `url` conflicts with `net/url` import in other files in
	// same package.
	url *netURL.URL
}

func NewUrl() URL {
	return &url{}
}

func NewURLFromNetURL(u *netURL.URL) URL {
	return url{u}
}

func (l url) GetHash() string {
	if l.url.Fragment == "" {
		return ""
	}
	return "#" + l.url.Fragment
}

func (l url) GetHost() string { return l.url.Host }

func (l url) GetHostname() string {
	return l.url.Hostname()
}

func (l url) GetHref() string { return l.url.String() }

func (l url) GetOrigin() string { return l.url.Scheme + "://" + l.url.Host }

func (l url) GetPathname() string { return l.url.Path }

func (l url) GetProtocol() string { return l.url.Scheme + ":" }

func (l url) GetPort() string { return l.url.Port() }

func (l url) GetSearch() string {
	if l.url.RawQuery != "" {
		return "?" + l.url.RawQuery
	} else {
		return ""
	}
}
