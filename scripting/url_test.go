package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 URL", func() {
	It("Is retrievable after construction", func() {
		ctx := NewTestContext()
		Expect(ctx.RunTestScript(`
			console.log("BEFORE")
			const u = new URL("foo/bar", "http://example.com");
			console.log("AFTER")
			u.href
		`)).To(Equal("http://example.com/foo/bar"))
	})
})
