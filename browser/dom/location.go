package dom

import netURL "net/url"

/*
TODO:
ancestorOrigins
hash
*/

type Location interface {
	URL
}

type location struct {
	URL
}

func NewLocation(url URL) Location {
	return location{url}
}

func NewLocationFromNetURL(url *netURL.URL) Location {
	return location{NewURLFromNetURL(url)}
}
