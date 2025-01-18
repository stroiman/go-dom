package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("V8 URL", func() {
	It("Is retrievable after construction", func() {
		ctx := NewTestContext()
		Expect(ctx.Eval(`
			const u = new URL("foo/bar", "http://example.com");
			u.href
		`)).To(Equal("http://example.com/foo/bar"))
	})
})
