package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("V8 NodeList", func() {
	It("Should be a direct descendant of Node", func() {
		ctx := NewTestContext()
		Expect(
			ctx.RunTestScript(
				`Object.getPrototypeOf(NodeList.prototype) === Object.prototype`,
			),
		).To(BeTrue())
	})

	Describe("Node list has 3 elements", func() {
		ctx := InitializeContext(LoadHTML(
			`<div id="1"></div><div id="2"></div><div id="3"></div>`,
		))

		It("Should have a length property", func() {
			Expect(ctx.RunTestScript("document.body.childNodes.length")).To(BeEquivalentTo(3))
		})

		It("Should allow get by calling `item`", func() {
			Expect(ctx.RunTestScript(
				`document.body.childNodes.item(1).getAttribute("id")`,
			)).To(Equal("2"))
		})

		It("Should allow get by calling []", func() {
			Expect(ctx.RunTestScript(
				`document.body.childNodes[1].getAttribute("id")`,
			)).To(Equal("2"))
		})

		It("Should return null when getting item out of range", func() {
			Expect(ctx.RunTestScript(
				`document.body.childNodes.item(5) === null`,
			)).To(BeTrue())
		})

		It(
			"Should return undefined when getting item out of range by indexed getter",
			func() {
				Expect(ctx.RunTestScript(
					`document.body.childNodes[5] === undefined`,
				)).To(BeTrue())
			},
		)
	})
})
