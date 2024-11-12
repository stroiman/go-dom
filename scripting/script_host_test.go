package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("ScriptHost", Ordered, func() {
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

		Describe("When a document is loaded", func() {
			It("Should have document instanceof Document", func() {
				Expect(
					ctx.MustRunTestScript("document instanceof Document"),
				).To(BeTrue())
				Expect(
					ctx.MustRunTestScript("Object.getPrototypeOf(document).constructor.name"),
				).To(Equal("Document"))
			})

			It("Should have document instanceof Document", func() {
				Expect(
					ctx.MustRunTestScript("document instanceof Document"),
				).To(BeTrue())
				Expect(
					ctx.MustRunTestScript("Object.getPrototypeOf(document).constructor.name"),
				).To(Equal("Document"))
				Expect(
					ctx.MustRunTestScript("Object.getPrototypeOf(document) === Document.prototype"),
				).To(BeTrue())
			})
		})

		Describe("Load document with script", func() {
			It("Runs the script when connected to DOM", func() {
				window := ctx.Window()
				window.SetScriptRunner(ctx)
				window.LoadHTML(`
<html>
  <body>
    <script>window.sut = document.outerHTML</script>
    <div>I should not be in the output</div>
  </body>
</html>
`,
				)
				Expect(
					ctx.MustRunTestScript("window.sut"),
				).To(Equal(`<html><head></head><body>
    <script>window.sut = document.outerHTML</script></body></html>`))
			})
		})
	})
})
