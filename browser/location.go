package browser

import (
	"net/url"
)

/*
TODO:
ancestorOrigins
hash
*/

type Location interface {
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

type location struct {
	url *url.URL
}

func NewLocation(url *url.URL) Location {
	return location{url}
}

func (l location) GetHash() string {
	if l.url.Fragment == "" {
		return ""
	}
	return "#" + l.url.Fragment
}

func (l location) GetHost() string { return l.url.Host }

func (l location) GetHostname() string {
	return l.url.Hostname()
}

func (l location) GetHref() string { return l.url.String() }

func (l location) GetOrigin() string { return l.url.Scheme + "://" + l.url.Host }

func (l location) GetPathname() string { return l.url.Path }

func (l location) GetProtocol() string { return l.url.Scheme + ":" }

func (l location) GetPort() string { return l.url.Port() }

func (l location) GetSearch() string {
	if l.url.RawQuery != "" {
		return "?" + l.url.RawQuery
	} else {
		return ""
	}
}
