package html_test

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/html"
	"github.com/stroiman/go-dom/browser/internal/domslices"
	. "github.com/stroiman/go-dom/browser/internal/http"
	. "github.com/stroiman/go-dom/browser/testing/gomega-matchers"
	"github.com/stroiman/go-dom/browser/testing/testservers"
)

var _ = Describe("Window", func() {
	It("Should have a document.documentElement instance of HTMLElement", func() {
		win, err := NewWindowReader(strings.NewReader("<html><body></body></html>"))
		Expect(err).ToNot(HaveOccurred())
		Expect(win.Document().DocumentElement()).To(BeHTMLElement())
	})

	It("Should respect the <!DOCTYPE>", func() {
		Skip("<!DOCTYPE> should be respected")
		// win, err := NewWindowReader(strings.NewReader("<!DOCTYPE HTML><html><body></body></html>"))
	})

	Describe("History()", func() {
		var win html.Window
		BeforeEach(func() {
			win = html.NewWindow(
				WindowOptions{
					HttpClient: NewHttpClientFromHandler(
						http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
							w.Write([]byte(fmt.Sprintf("<body><h1>%s</h1></body>", r.URL.Path)))
						}),
					),
				},
			)
		})

		It("Should have a length of one when starting", func() {
			Expect(win.History().Length()).To(Equal(1))
		})

		It("Should have a length of two when navigating", func() {
			Expect(win.Navigate("/page-2")).To(Succeed())
			Expect(win.History().Length()).To(Equal(2))
		})

		It("Should have a length of two when navigating", func() {
			Expect(win.Navigate("/page-2")).To(Succeed())
			Expect(win.History().Length()).To(Equal(2))
			Expect(win.Document().QuerySelector("h1")).To(HaveTextContent("/page-2"))
		})

		It("Should go back, but keep the length", func() {
			Expect(win.Navigate("/page-2")).To(Succeed())
			Expect(win.History().Go(-1)).To(Succeed())
			Expect(win.Document().QuerySelector("h1")).To(HaveTextContent("Go-DOM"))
			Expect(win.Location().Href()).To(Equal("about:blank"))
		})

		It("Should truncate history when going forward", func() {
			Expect(win.Navigate("/page-2")).To(Succeed())
			Expect(win.Navigate("/page-3")).To(Succeed())
			Expect(win.Navigate("/page-4")).To(Succeed())
			Expect(win.Navigate("/page-5")).To(Succeed())
			Expect(win.History().Length()).To(Equal(5))
			Expect(win.History().Go(-3)).To(Succeed())
			Expect(win.History().Length()).To(Equal(5))
			Expect(win.Navigate("/page-6")).To(Succeed())
			Expect(win.History().Length()).To(Equal(3))
			Expect(win.Location().Pathname()).To(Equal("/page-6"))

		})
	})

	Describe("Location()", func() {
		var window Window

		BeforeEach(func() {
			server := testservers.NewAnchorTagNavigationServer()
			DeferCleanup(func() { server = nil })
			window = NewWindowFromHandler(server)
		})

		It("Should be about:blank", func() {
			Expect(window.Location().Href()).To(Equal("about:blank"))
		})

		It("Should return the path loaded from", func() {
			Expect(window.Navigate("/index")).To(Succeed())
			Expect(window.Location().Pathname()).To(Equal("/index"))
		})

		Describe("Navigate", func() {
			It("Should load a blank page when loading about:blank", func() {
				Expect(window.Navigate("about:blank")).To(Succeed())
				Expect(window.Document().QuerySelector("h1")).To(HaveTextContent("Go-DOM"))
			})
		})

		Describe("User navigation (clicking links)", func() {
			var links []dom.Node

			BeforeEach(func() {
				Expect(window.Navigate("/index")).To(Succeed())
				nodes, err := window.Document().QuerySelectorAll("a")
				Expect(err).ToNot(HaveOccurred())
				links = nodes.All()
			})

			It("Should update when using a link with absolute url", func() {
				link, ok := domslices.SliceFindFunc(links, func(n dom.Node) bool {
					return n.TextContent() == "Products from absolute url"
				})
				Expect(ok).To(BeTrue())
				link.(dom.Element).Click()
				Expect(window.Location().Pathname()).To(Equal("/products"))
			})

			It("Should update when using a link with relative url", func() {
				link, ok := domslices.SliceFindFunc(links, func(n dom.Node) bool {
					return n.TextContent() == "Products from relative url"
				})
				Expect(ok).To(BeTrue())
				link.(dom.Element).Click()
				Expect(window.Location().Pathname()).To(Equal("/products"))
			})
		})
	})
})
