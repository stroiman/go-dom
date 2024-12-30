package dom_test

import (
	"net/http"

	"github.com/stroiman/go-dom/browser/html"
	domHTTP "github.com/stroiman/go-dom/browser/internal/http"

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
				win.Dispose()
			}
		})
		return win
	}

	It("Should parse a location without a query", func() {
		win := OpenWindow("http://localhost:9999/foo/bar")
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
		win := OpenWindow("http://localhost:9999/foo#heading-1")
		location := win.Location()
		Expect(location.GetHost()).To(Equal("localhost:9999"), "host")
		Expect(location.GetHash()).To(Equal("#heading-1"), "hash")
		Expect(location.GetHostname()).To(Equal("localhost"), "hostname")
		Expect(location.GetHref()).To(Equal("http://localhost:9999/foo#heading-1"), "href")
	})

	It("Should parse a location with a query", func() {
		win := OpenWindow(
			"http://localhost:9999/foo/bar?q=baz",
		)
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
