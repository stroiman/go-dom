package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/dom"
)

var _ = Describe("window.location", func() {
	It("Should have the location of the document", func() {
		window := dom.NewWindow(dom.WindowOptionLocation("http://example.com/foo"))
		ctx := host.NewContext(window)
		DeferCleanup(func() {
			ctx.Dispose()
		})
		Expect(ctx.Eval("location.href")).To(Equal("http://example.com/foo"))
	})
})
