package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("V8 Append (element, document, document fragment)", func() {
	It("Should accept multiple values", func() {
		ctx := scriptTestSuite.NewContext()
		Expect(ctx.Eval(`
			const d = document.createElement("div")
			d.append(
				document.createElement("p"),
				document.createElement("p"),
			);
			d.outerHTML`)).To(Equal("<div><p></p><p></p></div>"))
	})

	It("Should accept text values", func() {
		ctx := scriptTestSuite.NewContext()
		Expect(ctx.Eval(`
			const d = document.createElement("div")
			d.append(
				document.createElement("p"),
				"Foo",
				"bar",
				document.createElement("p"),
			);
			d.outerHTML`)).To(Equal("<div><p></p>Foobar<p></p></div>"))
	})
})
