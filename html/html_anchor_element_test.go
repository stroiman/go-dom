package html_test

import (
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	matchers "github.com/gost-dom/browser/testing/gomega-matchers"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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

	It("Should not have an href", func() {
		Expect(a.Href()).To(Equal(""))
	})

	It("Should not have a hostname when no href is set", func() {
		Expect(a.Hostname()).To(Equal(""))
	})

	It("Should not have a hostname when href is set", func() {
		a.SetAttribute("href", "/foo/bar")
		Expect(a.Hostname()).To(Equal("example.com"))
		Expect(a.Hostname()).To(Equal(win.Location().Hostname()))
	})

	It("Should provide an absolute href", func() {
		a.SetAttribute("href", "/local")
		Expect(a.Href()).To(Equal("http://example.com/local"))
	})

	It("Should ignore the most data properties when there's no href", func() {
		a.SetPathname("/local")
		Expect(a.Pathname()).To(Equal(""))
	})

	It("Should update the href data attribute when stetting an idl attribute", func() {
		a.SetHref("/")
		a.SetPathname("/local")
		Expect(a).To(matchers.HaveAttribute("href", "http://example.com/local"))
	})
})
