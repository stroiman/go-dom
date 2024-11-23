package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 Element", func() {
	ctx := InitializeContextWithEmptyHtml()

	It("It should be a direct descendant of Node", func() {
		Expect(
			ctx.RunTestScript("Object.getPrototypeOf(Element.prototype) === Node.prototype"),
		).To(BeTrue())
	})

	Describe("GetAttribute", func() {
		It("Should return the right value", func() {
			ctx.Window().LoadHTML(`<div id="1" class="foo"></div>`)
			Expect(ctx.RunTestScript(
				`document.getElementById("1").getAttribute("class")`,
			)).To(Equal("foo"))
		})
	})
})
