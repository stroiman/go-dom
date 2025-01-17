package gojahost_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("goja: Node", func() {
	Describe("Prototype", func() {
		It("Should have function installed", func() {
			ctx := newCtx()
			Expect(ctx.Run("const proto = Node.prototype")).To(Succeed())
			Expect(ctx.Eval("typeof proto.contains")).To(Equal("function"), "contains")
			Expect(ctx.Eval("typeof proto.insertBefore")).To(Equal("function"), "insertBefore")
			Expect(ctx.Eval("typeof proto.appendChild")).To(Equal("function"), "appendChild")
			Expect(ctx.Eval("typeof proto.removeChild")).To(Equal("function"), "removeChild")
		})
	})
})
