package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 Document", Ordered, func() {
	ctx := InitializeContextWithEmptyHtml()

	Describe("Constructor", func() {
		It("Should be instance of Document", func() {
			Expect(ctx.RunTestScript(`
        const doc = new Document();
        doc instanceof Document && doc != document;
      `)).To(BeTrue())
		})

		It("Should have `createElement` as a function", func() {
			Expect(
				ctx.RunTestScript(`typeof (new Document().createElement)`),
			).To(Equal("function"))
		})

		It("Should support Document functions", func() {
			Skip("createElement and HTMLElement are missing")
			Expect(
				ctx.RunTestScript(`document.createElement("div") instanceof HTMLElement`),
			).Error().
				ToNot(HaveOccurred())
		})
	})
})
