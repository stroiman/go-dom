package http

import (
	netHTTP "net/http"
	"net/http/cookiejar"
)

func NewHttpClientFromHandler(handler netHTTP.Handler) netHTTP.Client {
	cookiejar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	return netHTTP.Client{
		Transport: TestRoundTripper{Handler: handler},
		Jar:       cookiejar,
	}
}
