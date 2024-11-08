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
		It("Should panic on error", func() {
			Expect(
				func() { ctx.MustRunTestScript("throw new Error()") },
			).To(Panic())
		})

		Describe("Global Object", func() {
			It("Should be accessible as `window`", func() {
				Expect(
					ctx.MustRunTestScript("globalThis === window && window === window.window"),
				).To(BeTrue())
			})

			It("It should have the right prototype", func() {
				Skip(
					"Prototype is not set as I thought it would be, need to understand.",
				)
				Expect(
					ctx.MustRunTestScript(
						"window.__proto__ === Object.prototype",
					),
				).To(BeTrue())
			})
		})
	})
})
