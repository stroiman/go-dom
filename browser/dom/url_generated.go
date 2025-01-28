// This file is generated. Do not edit.

package dom

type URL interface {
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
	SearchParams() string
	Hash() string
	SetHash(string)
	ToJSON() (string, error)
}
