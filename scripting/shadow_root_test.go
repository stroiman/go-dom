package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 ShadowRoot", func() {
	ctx := InitializeContextWithEmptyHtml()

	It("Should be a direct descendant of DocumentFragment", func() {
		Expect(
			ctx.RunTestScript(
				`Object.getPrototypeOf(ShadowRoot.prototype).constructor.name`,
			),
		).To(Equal("DocumentFragment"))
	})
})
