package scripting_test

import (
	"fmt"
	"net/url"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("window.location", func() {
	It("Should have the location of the document", func() {
		url, err := url.Parse("http://example.com/foo")
		Expect(err).ToNot(HaveOccurred())
		window := browser.NewWindow(url)
		ctx := host.NewContext(window)
		DeferCleanup(func() {
			fmt.Println("Disposing context")
			ctx.Dispose()
		})
		Expect(ctx.Run("location.href")).To(Equal("http://example.com/foo"))
	})
})
