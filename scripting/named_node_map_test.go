package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 NamedNodeMap", func() {
	ctx := InitializeContextWithEmptyHtml()

	It("Should inherit directly from Object", func() {
		Expect(
			ctx.RunTestScript("Object.getPrototypeOf(NamedNodeMap.prototype) === Object.prototype"),
		).To(BeTrue())
	})
})
