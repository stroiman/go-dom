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
	})
})
