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

	It("Should have a nodeType of 11 ", func() {
		Skip("ShadowRoot is defined as a type, but there isn't a way to construct one yet")
		// It's not specified here: https://developer.mozilla.org/en-US/docs/Web/API/Node
		// so I assume that it inherits from DocumentFragment as it inherits from it
		// https://developer.mozilla.org/en-US/docs/Web/API/ShadowRoot
		Expect(
			ctx.RunTestScript(
				`new ShadowRoot().nodeType`,
			),
		).To(BeEquivalentTo(11))
	})
})
