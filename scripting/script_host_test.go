package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("ScriptHost", Ordered, func() {
	var (
		host *ScriptHost
		ctx  *ScriptContext
	)

	BeforeAll(func() {
		host = NewScriptHost()
	})
	BeforeEach(func() {
		ctx = host.NewContext()
	})
	AfterEach(func() {
		ctx.Dispose()
	})
	AfterAll(func() {
		host.Dispose()
	})

	Describe("Script host", func() {
		It("Has window as the global object", func() {
			result, err := ctx.RunScript("globalThis === window && window === window.window")
			Expect(err).ToNot(HaveOccurred())
			Expect(result.Boolean()).To(BeTrue())
			result2, err := ctx.RunScript("1 === 2")
			Expect(err).ToNot(HaveOccurred())
			Expect(result2.Boolean()).To(BeFalse())

		})
	})
})
