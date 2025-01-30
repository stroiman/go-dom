package dom_test

import (
	"net/http"

	"github.com/gost-dom/browser/html"
	domHTTP "github.com/gost-dom/browser/internal/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Window.Location", func() {
	OpenWindow := func(location string) html.Window {
		handler := http.HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) { res.Write([]byte("<html></html>")) },
		)
		windowOptions := html.WindowOptions{
			HttpClient: domHTTP.NewHttpClientFromHandler(handler),
		}
		win, err := html.OpenWindowFromLocation(location, windowOptions)
		Expect(err).ToNot(HaveOccurred())
		DeferCleanup(func() {
			if win != nil {
				win.Close()
			}
		})
		return win
	}

	It("Should parse a location without a query", func() {
		win := OpenWindow("http://localhost:9999/foo/bar")
		location := win.Location()
		Expect(location.Host()).To(Equal("localhost:9999"), "host")
		Expect(location.Hash()).To(Equal(""), "hash")
		Expect(location.Hostname()).To(Equal("localhost"), "hostname")
		Expect(location.Href()).To(Equal("http://localhost:9999/foo/bar"), "href")
		Expect(location.Origin()).To(Equal("http://localhost:9999"), "origin")
		Expect(location.Pathname()).To(Equal("/foo/bar"), "Pathname")
		Expect(location.Port()).To(Equal("9999"), "port")
		Expect(location.Protocol()).To(Equal("http:"), "protocol")
		Expect(location.Search()).To(Equal(""), "query")
	})

	It("Should parse a location without a query", func() {
		win := OpenWindow("http://localhost:9999/foo#heading-1")
		location := win.Location()
		Expect(location.Host()).To(Equal("localhost:9999"), "host")
		Expect(location.Hash()).To(Equal("#heading-1"), "hash")
		Expect(location.Hostname()).To(Equal("localhost"), "hostname")
		Expect(location.Href()).To(Equal("http://localhost:9999/foo#heading-1"), "href")
	})

	It("Should parse a location with a query", func() {
		win := OpenWindow(
			"http://localhost:9999/foo/bar?q=baz",
		)
		location := win.Location()
		Expect(location.Host()).To(Equal("localhost:9999"), "host")
		Expect(location.Hostname()).To(Equal("localhost"), "hostname")
		Expect(location.Href()).To(Equal("http://localhost:9999/foo/bar?q=baz"), "href")
		Expect(location.Origin()).To(Equal("http://localhost:9999"), "origin")
		Expect(location.Pathname()).To(Equal("/foo/bar"), "Pathname")
		Expect(location.Port()).To(Equal("9999"), "port")
		Expect(location.Protocol()).To(Equal("http:"), "protocol")
		Expect(location.Search()).To(Equal("?q=baz"), "query")
	})
})
