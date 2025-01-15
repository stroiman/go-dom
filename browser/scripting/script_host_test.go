package scripting_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/html"
	"github.com/stroiman/go-dom/browser/scripting"
)

var _ = Describe("ScriptHost", func() {
	ctx := InitializeContext()

	Describe("Script host", func() {
		Describe("Global Object", func() {
			It("Should be accessible as `window`", func() {
				Expect(
					ctx.RunTestScript("globalThis === window && window === window.window"),
				).To(BeTrue())
			})

			It("It should have the prototype 'Window'", func() {
				Skip(
					"This is desired behaviour, but I haven't yet grokked the prototype on ObjectTemplates.",
				)
				Expect(
					ctx.RunTestScript(
						`Object.getPrototypeOf(window).constructor.name === "Window"`,
					),
				).To(BeTrue())
			})
		})

		Describe("Load document with script", func() {
			It("Runs the script when connected to DOM", func() {
				reader := strings.NewReader(`<html><body>
    <script>window.sut = document.documentElement.outerHTML</script>
    <div>I should not be in the output</div>
  </body></html>
`)
				options := html.WindowOptions{ScriptHost: (*scripting.Wrapper)(host)}
				win, err := html.NewWindowReader(reader, options)
				defer win.Close()
				Expect(err).ToNot(HaveOccurred())
				ctx := win.GetScriptEngine()
				Expect(
					ctx.Eval("window.sut"),
				).To(Equal(`<html><head></head><body>
    <script>window.sut = document.documentElement.outerHTML</script></body></html>`))
			})
		})
	})
})
