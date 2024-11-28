package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 Element", func() {
	It("It should be a direct descendant of Node", func() {
		ctx := NewTestContext()
		Expect(
			ctx.RunTestScript("Object.getPrototypeOf(Element.prototype) === Node.prototype"),
		).To(BeTrue())
	})

	It("Should have nodeType == 1", func() {
		ctx := NewTestContext(LoadHTML(`<div id="1" class="foo"></div>`))
		Expect(ctx.RunTestScript("document.body.nodeType")).To(BeEquivalentTo(1))
	})

	Describe("Attributes", func() {
		It("Should support getAtribute", func() {
			ctx := NewTestContext(LoadHTML(`<div id="1" class="foo"></div>`))
			Expect(ctx.RunTestScript(
				`document.getElementById("1").getAttribute("class")`,
			)).To(Equal("foo"))
		})

		It("Should support hasAtribute", func() {
			ctx := NewTestContext(LoadHTML(`<div id="1" class="foo"></div>`))
			Expect(ctx.RunTestScript(
				`document.getElementById("1").hasAttribute("class")`,
			)).To(BeTrue())
			Expect(ctx.RunTestScript(
				`document.getElementById("1").hasAttribute("foo")`,
			)).To(BeFalse())
		})
	})

	It("Should support insertAdjacentHTML", func() {
		ctx := NewTestContext(LoadHTML(`<div id="1" class="foo"></div>`))
		Expect(ctx.RunTestScript(
			`document.getElementById("1").insertAdjacentHTML("beforebegin", "<p>foo</p>")`,
		)).Error().ToNot(HaveOccurred())
		Expect(
			ctx.Window().Document().Body().OuterHTML(),
		).To(Equal(`<body><p>foo</p><div id="1" class="foo"></div></body>`))
	})
})
