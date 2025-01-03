package scripting_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/html"
	"github.com/stroiman/go-dom/browser/scripting"
)

var _ = Describe("Window", func() {
	ctx := InitializeContext()

	Describe("Constructor", func() {
		It("Should be defined", func() {
			Expect(ctx.RunScript("Window")).ToNot(BeNil())
		})

		It("Should not be callable", func() {
			Expect(ctx.RunTestScript("Window()")).Error().To(HaveOccurred())
		})

		It("Should not be newable", func() {
			Expect(ctx.RunTestScript("new Window()")).Error().To(HaveOccurred())
		})
	})

	Describe("Inheritance", func() {
		It("Should be an EventTarget", func() {
			Expect(ctx.RunTestScript("window instanceof EventTarget")).To(BeTrue())
		})
	})

	Describe("location property", func() {
		It("Should be a Location", func() {
			Expect(ctx.RunTestScript("window.location instanceof Location")).To(BeTrue())
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
</script></body>`), html.WindowOptions{ScriptEngineFactory: (*scripting.Wrapper)(host)},
				)
				Expect(err).ToNot(HaveOccurred())
				ctx := win.GetScriptEngine()
				Expect(ctx.Eval("scripts.join(',')")).To(Equal("DOMContentLoaded,load"))
			})
		})
	})
})
