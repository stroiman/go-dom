package scripting_test

import (
	netURL "net/url"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/dom"
	// . "github.com/stroiman/go-dom/browser/scripting"
)

var _ = Describe("window.location", func() {
	It("Should have the location of the document", func() {
		url, err := netURL.Parse("http://example.com/foo")
		Expect(err).ToNot(HaveOccurred())
		window := dom.NewWindow(url)
		ctx := host.NewContext(window)
		DeferCleanup(func() {
			ctx.Dispose()
		})
		Expect(ctx.Eval("location.href")).To(Equal("http://example.com/foo"))
	})
})
