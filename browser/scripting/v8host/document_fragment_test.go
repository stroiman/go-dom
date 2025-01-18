package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("V8 DocumentFragment", func() {
	ctx := InitializeContextWithEmptyHtml()

	It("Should be a direct descendant of Node", func() {
		Expect(
			ctx.Eval(
				`Object.getPrototypeOf(DocumentFragment.prototype) === Node.prototype`,
			),
		).To(BeTrue())
	})

	It("Should have query functions", func() {
		ctx.MustRunTestScript(`const fragment = document.createDocumentFragment()`)
		Expect(
			ctx.Eval(
				`typeof fragment.querySelector`,
			),
		).To(Equal("function"))
		Expect(
			ctx.Eval(
				`typeof document.createDocumentFragment().querySelectorAll`,
			),
		).To(Equal("function"))
	})
})
