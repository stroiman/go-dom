package html_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"
	matchers "github.com/stroiman/go-dom/browser/testing/gomega-matchers"
)

var _ = Describe("htmlAnchorElement", func() {
	var (
		win html.Window
		doc dom.Document
		a   html.HTMLAnchorElement
	)

	BeforeEach(func() {
		win = html.NewWindow(html.WindowOptions{
			BaseLocation: "http://example.com/",
		})
		doc = win.Document()
		a = doc.CreateElement("a").(html.HTMLAnchorElement)
	})

	It("Should provide an absolute href", func() {
		a.SetAttribute("href", "/local")
		Expect(a.Href()).To(Equal("http://example.com/local"))
	})

	It("Should update the href data attribute when stetting an idl attribute", func() {
		a.SetPathname("/local")
		Expect(a).To(matchers.HaveAttribute("href", "http://example.com/local"))
	})
})
