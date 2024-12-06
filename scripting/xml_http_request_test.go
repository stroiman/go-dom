package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 XmlHttpRequest", func() {
	It("Should be an EventTarget", func() {
		ctx := NewTestContext()
		Expect(ctx.RunTestScript("new XMLHttpRequest() instanceof EventTarget")).To(BeTrue())
	})
})
