package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("ScriptHost", Ordered, func() {
	var (
		host *ScriptHost
		ctx  TestScriptContext
	)

	BeforeAll(func() {
		host = NewScriptHost()
	})
	BeforeEach(func() {
		ctx = TestScriptContext{host.NewContext()}
	})
	AfterEach(func() {
		ctx.Dispose()
	})
	AfterAll(func() {
		host.Dispose()
	})

	Describe("Script host", func() {
		It("Has window as the global object", func() {
			result := ctx.MustRunTestScript("globalThis === window && window === window.window")
			Expect(result).To(BeTrue())
			result2 := ctx.MustRunTestScript("1 === 2")
			Expect(result2).To(BeFalse())
		})

		It("Should panic on error", func() {
			Expect(
				func() { ctx.MustRunTestScript("throw new Error()") },
			).To(Panic())
		})

		It("Should handle bools", func() {
			Expect(ctx.MustRunTestScript("1 === 1")).To(BeTrue())
		})
	})
})
