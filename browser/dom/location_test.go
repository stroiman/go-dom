package dom_test

import (
	"net/http"

	. "github.com/stroiman/go-dom/browser/dom"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Window.Location", func() {
	It("Should parse a location without a query", func() {
		handler := http.HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) { res.Write([]byte("<html></html>")) },
		)
		browser := NewBrowserFromHandler(handler)
		win, err := browser.OpenWindow("http://localhost:9999/foo/bar")
		Expect(err).ToNot(HaveOccurred())
		location := win.Location()
		Expect(location.GetHost()).To(Equal("localhost:9999"), "host")
		Expect(location.GetHash()).To(Equal(""), "hash")
		Expect(location.GetHostname()).To(Equal("localhost"), "hostname")
		Expect(location.GetHref()).To(Equal("http://localhost:9999/foo/bar"), "href")
		Expect(location.Origin()).To(Equal("http://localhost:9999"), "origin")
		Expect(location.GetPathname()).To(Equal("/foo/bar"), "Pathname")
		Expect(location.GetPort()).To(Equal("9999"), "port")
		Expect(location.GetProtocol()).To(Equal("http:"), "protocol")
		Expect(location.GetSearch()).To(Equal(""), "query")
	})

	It("Should parse a location without a query", func() {
		handler := http.HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) { res.Write([]byte("<html></html>")) },
		)
		browser := NewBrowserFromHandler(handler)
		win, err := browser.OpenWindow("http://localhost:9999/foo#heading-1")
		Expect(err).ToNot(HaveOccurred())
		location := win.Location()
		Expect(location.GetHost()).To(Equal("localhost:9999"), "host")
		Expect(location.GetHash()).To(Equal("#heading-1"), "hash")
		Expect(location.GetHostname()).To(Equal("localhost"), "hostname")
		Expect(location.GetHref()).To(Equal("http://localhost:9999/foo#heading-1"), "href")
	})

	It("Should parse a location with a query", func() {
		handler := http.HandlerFunc(
			func(res http.ResponseWriter, req *http.Request) { res.Write([]byte("<html></html>")) },
		)
		browser := NewBrowserFromHandler(handler)
		win, err := browser.OpenWindow("http://localhost:9999/foo/bar?q=baz")
		Expect(err).ToNot(HaveOccurred())
		location := win.Location()
		Expect(location.GetHost()).To(Equal("localhost:9999"), "host")
		Expect(location.GetHostname()).To(Equal("localhost"), "hostname")
		Expect(location.GetHref()).To(Equal("http://localhost:9999/foo/bar?q=baz"), "href")
		Expect(location.Origin()).To(Equal("http://localhost:9999"), "origin")
		Expect(location.GetPathname()).To(Equal("/foo/bar"), "Pathname")
		Expect(location.GetPort()).To(Equal("9999"), "port")
		Expect(location.GetProtocol()).To(Equal("http:"), "protocol")
		Expect(location.GetSearch()).To(Equal("?q=baz"), "query")
	})
})
