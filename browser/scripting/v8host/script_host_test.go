package v8host_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/gost-dom/browser/browser/html"
)

var _ = Describe("ScriptHost", func() {
	ctx := InitializeContext()

	Describe("Script host", func() {
		Describe("Global Object", func() {
			It("Should be accessible as `window`", func() {
				Expect(
					ctx.Eval("globalThis === window && window === window.window"),
				).To(BeTrue())
			})

			It("It should have the prototype 'Window'", func() {
				Skip(
					"This is desired behaviour, but I haven't yet grokked the prototype on ObjectTemplates.",
				)
				Expect(
					ctx.Eval(
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
				options := html.WindowOptions{ScriptHost: host}
				win, err := html.NewWindowReader(reader, options)
				defer win.Close()
				Expect(err).ToNot(HaveOccurred())
				ctx := win.ScriptContext()
				Expect(
					ctx.Eval("window.sut"),
				).To(Equal(`<html><head></head><body>
    <script>window.sut = document.documentElement.outerHTML</script></body></html>`))
			})
		})
	})
})
