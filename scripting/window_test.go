package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
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

	Describe("Window Events", func() {
		Describe("DOMContentLoaded", func() {
			It("Should be fired _after_ script has executed", func() {
				ctx.Window().SetScriptRunner(ctx.ScriptContext)
				Expect(ctx.Window().LoadHTML(`<body><script>
  scripts = []
  function listener1() {
    scripts.push("DOMContentLoaded")
  }
  function listener2() {
    scripts.push("load")
  }
  window.document.addEventListener("DOMContentLoaded", listener1);
  window.document.addEventListener("load", listener2);
</script></body>`)).To(Succeed())
				Expect(ctx.RunTestScript("scripts.join(',')")).To(Equal("DOMContentLoaded,load"))
			})
		})
	})
})
