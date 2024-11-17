package scripting_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("Load from server", func() {
	ctx := InitializeContextWithEmptyHtml()

	It("Loads from an HTTP server", func() {
		server := http.NewServeMux()
		server.Handle(
			"GET /index.html",
			http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.Write([]byte("<body>Hello, World!</body>"))
			}),
		)
		browser := ctx.NewBrowserFromHandler(server)
		window, err := browser.OpenWindow("/index.html")
		Expect(err).ToNot(HaveOccurred())
		Expect(window.Document().Body().OuterHTML()).To(Equal("<body>Hello, World!</body>"))
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
		browser := ctx.NewBrowserFromHandler(server)
		Expect(browser.OpenWindow("/not-found.html")).Error().To(HaveOccurred())
	})

	It("Executes script from script tags", func() {
		server := http.NewServeMux()
		server.Handle(
			"GET /index.html",
			http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.Write(
					[]byte(
						`<html><head><script src="/js/script.js"></script></head><body>Hello, World!</body>`,
					),
				)
			}),
		)
		server.HandleFunc(
			"GET /js/script.js",
			func(res http.ResponseWriter, req *http.Request) {
				res.Header().Add("Content-Type", "text/javascript")
				res.Write([]byte(`var scriptLoaded = true`))
			},
		)
		browser := ctx.NewBrowserFromHandler(server)
		Expect(browser.OpenWindow("/index.html")).Error().ToNot(HaveOccurred())
		Expect(ctx.RunTestScript("window.scriptLoaded")).To(BeTrue())
	})
})
