package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 DocumentFragment", func() {
	ctx := InitializeContextWithEmptyHtml()

	It("Should be a direct descendant of Node", func() {
		Expect(
			ctx.RunTestScript(
				`Object.getPrototypeOf(DocumentFragment.prototype) === Node.prototype`,
			),
		).To(BeTrue())
	})

	It("Should have query functions", func() {
		ctx.MustRunTestScript(`const fragment = document.createDocumentFragment()`)
		Expect(
			ctx.RunTestScript(
				`typeof fragment.querySelector`,
			),
		).To(Equal("function"))
		Expect(
			ctx.RunTestScript(
				`typeof document.createDocumentFragment().querySelectorAll`,
			),
		).To(Equal("function"))
	})
})
