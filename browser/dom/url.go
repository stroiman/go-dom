package dom

import (
	netURL "net/url"
	"strings"
)

type URL interface {
	GetHash() string
	GetHost() string
	GetHostname() string
	GetHref() string
	// TODO
	// SetHref(href string)
	Origin() string
	GetPathname() string
	GetPort() string
	GetProtocol() string
	GetSearch() string
	ToJSON() (string, error)
}

type url struct {
	// Cannot be named `url` conflicts with `net/url` import in other files in
	// same package.
	url *netURL.URL
}

func NewUrl(rawUrl string) (URL, error) {
	if res, err := netURL.Parse(rawUrl); err == nil {
		return &url{res}, nil
	} else {
		return nil, err
	}
}

func NewUrlBase(relativeUrl string, base string) (result URL, err error) {
	var u *netURL.URL
	if u, err = netURL.Parse(base); err != nil {
		return
	}
	u.RawQuery = ""
	u.Fragment = ""
	if strings.HasPrefix(relativeUrl, "/") {
		u.Path = relativeUrl
	} else {
		// A DOM Url treats the relative path from the last slash in the base URL.
		// Go's URL doesn't. To get the right behaviour, use the parent dir if the
		// base path doesn't end in a slash
		if u.Path != "" && !strings.HasSuffix(u.Path, "/") {
			u = u.JoinPath("..")
		}
		u = u.JoinPath(relativeUrl)
	}
	result = NewURLFromNetURL(u)
	return
}

func ParseURL(rawUrl string) URL {
	res, err := NewUrl(rawUrl)
	if err != nil {
		res = nil
	}
	return res
}

func ParseURLBase(relativeUrl string, base string) URL {
	res, err := NewUrlBase(relativeUrl, base)
	if err != nil {
		res = nil
	}
	return res
}

func CanParseURL(rawUrl string) bool {
	_, err := NewUrl(rawUrl)
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

func (l url) Origin() string { return l.url.Scheme + "://" + l.url.Host }

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

func (l url) ToJSON() (string, error) { return l.GetHref(), nil }
