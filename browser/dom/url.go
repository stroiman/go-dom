package dom

import (
	"fmt"
	netURL "net/url"
	"strings"
)

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
	return &url{u}
}

func (l url) Hash() string {
	if l.url.Fragment == "" {
		return ""
	}
	return "#" + l.url.Fragment
}

func (l url) Host() string { return l.url.Host }

func (l url) Hostname() string {
	return l.url.Hostname()
}

func (l url) Href() string { return l.url.String() }

func (l url) Origin() string { return l.url.Scheme + "://" + l.url.Host }

func (l url) Pathname() string { return l.url.Path }

func (l url) Protocol() string { return l.url.Scheme + ":" }

func (l url) Port() string { return l.url.Port() }

func (l url) Search() string {
	if l.url.RawQuery != "" {
		return "?" + l.url.RawQuery
	} else {
		return ""
	}
}

func (l url) ToJSON() (string, error) { return l.Href(), nil }

func (l url) Password() string { p, _ := l.url.User.Password(); return p }
func (l *url) SetPassword(val string) {
	l.url.User = netURL.UserPassword(l.Username(), val)
}

func (l *url) SetUsername(
	val string,
) {
	p, _ := l.url.User.Password()
	l.url.User = netURL.UserPassword(val, p)
}
func (l url) Username() string { return l.url.User.Username() }
func (l url) SearchParams() string {
	// TODO
	panic("URL.SearchParams Not implemented. Please file a feature request")
}

func (l *url) SetHash(val string) { l.url.Fragment = val }
func (l *url) SetHost(val string) { l.url.Host = val }
func (l *url) SetHostname(val string) {
	port := l.Port()
	if port == "" {
		l.url.Host = val
	} else {
		l.url.Host = fmt.Sprintf("%s:%s", val, port)
	}
}
func (l *url) SetHref(val string) {
	v, err := netURL.Parse(val)
	// TODO: Generate error
	if err != nil {
		panic(err)
	}
	l.url = v
}
func (l *url) SetPathname(val string) { l.url.Path = val }
func (l *url) SetPort(val string) {
	if val == "" {
		l.url.Host = l.url.Hostname()
	} else {
		l.url.Host = fmt.Sprintf("%s:%s", l.Hostname(), val)
	}
}
func (l *url) SetProtocol(val string) { l.url.Scheme = val }
func (l *url) SetSearch(val string) {
	if strings.HasPrefix(val, "?") {
		val = val[1:]
	}
	l.url.RawQuery = val
}
