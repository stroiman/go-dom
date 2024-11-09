package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("ScriptHost", Ordered, func() {
	// var (
	// 	ctx TestScriptContext
	// )
	// BeforeEach(func() {
	// 	ctx = TestScriptContext{host.NewContext()}
	// })
	// AfterEach(func() {
	// 	ctx.Dispose()
	// })
	ctx := InitializeContext()

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

			It("It should have the prototype 'Window'", func() {
				Skip(
					"This is desired behaviour, but I haven't yet grokked the prototype on ObjectTemplates.",
				)
				Expect(
					ctx.MustRunTestScript(
						`Object.getPrototypeOf(window).constructor.name === "Window"`,
					),
				).To(BeTrue())
			})
		})
	})
})
