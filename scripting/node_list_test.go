package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("V8 NodeList", func() {
	ctx := InitializeContextWithEmptyHtml()

	It("Should be a direct descendant of Node", func() {
		Expect(
			ctx.RunTestScript(
				`Object.getPrototypeOf(NodeList.prototype) === Object.prototype`,
			),
		).To(BeTrue())
	})
})
