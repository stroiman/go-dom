package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DOM Parser", func() {
	It("Should return a Document", func() {
		ctx := NewTestContext()
		Expect(ctx.RunTestScript(`
			const parser = new DOMParser()
			const doc = parser.parseFromString("<div id='target'></div>", "text/html")
			console.log("Constructor", Object.getPrototypeOf(doc).constructor.name)
		`)).Error().ToNot(HaveOccurred())
		Expect(
			ctx.RunTestScript("Object.getPrototypeOf(doc) === Document.prototype"),
		).To(BeTrue(), "result is a Document")
		Expect(
			ctx.RunTestScript("doc === window.document"),
		).To(BeFalse(), "Window.document isn't replaced")
		Expect(
			ctx.RunTestScript("doc.getElementById('target') instanceof HTMLDivElement"),
		).To(BeTrue(), "Element is a div")
	})

	It("Should return an HTMLDocument", func() {
		ctx := NewTestContext()
		Expect(ctx.RunTestScript(`
			const parser = new DOMParser()
			const doc = parser.parseFromString("<div id='target'></div>", "text/html")
			console.log("Constructor", Object.getPrototypeOf(doc).constructor.name)
		`)).Error().ToNot(HaveOccurred())
		Skip("HTMLDocument not properly implemented")
		Expect(
			ctx.RunTestScript("Object.getPrototypeOf(doc) === HTMLDocument.prototype"),
		).To(BeTrue(), "result is an HTMLDocument")
	})
})
