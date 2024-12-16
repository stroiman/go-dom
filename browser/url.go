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

func ParseURL(rawUrl string) (URL, error) {
	if res, err := netURL.Parse(rawUrl); err == nil {
		return &url{res}, nil
	} else {
		return nil, err
	}
}

func CanParseURL(rawUrl string) bool {
	_, err := ParseURL(rawUrl)
	return err == nil
}

func CreateObjectURL(object any) (URL, error) {
	return nil, newNotImplementedError("URL.CreateObjectURL not implemented yet")
}

func RevokeObjectURL(object any) (URL, error) {
	return nil, newNotImplementedError("URL.RevokeObjectURL not implemented yet")
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
