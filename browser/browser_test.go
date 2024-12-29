package browser_test

import (
	"net/http"

	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Browser", func() {
	It("Should be able to read from an http.Handler instance", func() {
		handler := (http.HandlerFunc)(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html") // For good measure, not used yet"
			w.Write([]byte("<html></html>"))
		})
		browser := NewBrowserFromHandler(handler)
		result, err := browser.Open("/")
		Expect(err).ToNot(HaveOccurred())
		element := result.Document().DocumentElement()

		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
	})

	It("It returns an error on non-200", func() {
		// This is not necessarily desired behaviour right now.
		server := http.NewServeMux()
		server.Handle(
			"GET /index.html",
			http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.Write([]byte("<body>Hello, World!</body>"))
			}),
		)
		browser := NewBrowserFromHandler(server)
		Expect(browser.Open("/not-found.html")).Error().To(HaveOccurred())
	})
})
