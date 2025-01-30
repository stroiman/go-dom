package suite

import (
	"strings"

	"github.com/gost-dom/browser/html"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func (suite *ScriptTestSuite) CreateWindowTests() {
	prefix := suite.Prefix
	Describe(prefix+"Window object", func() {
		var ctx html.ScriptContext

		BeforeEach(func() {
			ctx = suite.NewContext()
		})

		Describe("location property", func() {
			It("Should be a Location", func() {
				Expect(ctx.Eval("window.location instanceof Location")).To(BeTrue())
			})
		})

		Describe("Inheritance", func() {
			It("Should be an EventTarget", func() {
				Expect(ctx.Eval("window instanceof EventTarget")).To(BeTrue())
			})
		})

		Describe("Constructor", func() {
			It("Should be defined", func() {
				Expect(
					ctx.Eval("Window && typeof Window === 'function'"),
				).To(BeTrue())
			})

			It("Should not be callable", func() {
				Expect(ctx.Eval("Window()")).Error().To(HaveOccurred())
			})

			It("Should throw a TypeError when constructed", func() {
				Expect(suite.NewWindow().Eval(
					`let error;
				try { new Window() } catch(err) { 
					error = err;
				}
				error && Object.getPrototypeOf(error).constructor.name
				`)).To(Equal("TypeError"))
			})
		})

		Describe("window.document", func() {
			It("Should have a document property", func() {
				Expect(suite.NewWindow().Eval("document instanceof Document")).To(BeTrue())
			})

			It("Is an enumerable property", func() {
				Expect(
					suite.NewWindow().
						Eval(`
							const keys = []
							for (let key in window) {
								keys.push(key);
							}
							keys.join(", ")
						`)).To(ContainSubstring("document"))
			})

			It("Should return the same instance on multiple calls", func() {
				Expect(suite.NewWindow().Eval(
					`const a = window.document;
					const b = window.document;
					a === b`,
				)).To(BeTrue())
			})
		})

		Describe("Window Events", func() {
			Describe("DOMContentLoaded", func() {
				It("Should be fired _after_ script has executed", func() {
					win, err := html.NewWindowReader(strings.NewReader(
						`<body><script>
  scripts = []
  function listener1() {
    scripts.push("DOMContentLoaded")
  }
  function listener2() {
    scripts.push("load")
  }
  window.document.addEventListener("DOMContentLoaded", listener1);
  window.document.addEventListener("load", listener2);
</script></body>`), html.WindowOptions{ScriptHost: suite.Engine},
					)
					DeferCleanup(func() { win.Close() })
					Expect(err).ToNot(HaveOccurred())
					ctx := win.ScriptContext()
					Expect(ctx.Eval("scripts.join(',')")).To(Equal("DOMContentLoaded,load"))
				})
			})
		})
	})
}
