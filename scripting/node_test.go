package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 Node", func() {
	Describe("contains", func() {
		It("Returns true, when passed a child", func() {
			ctx := NewTestContext(LoadHTML(`<div id="parent-1"><div id="child">child</div></div>`))
			Expect(
				ctx.RunTestScript(
					`document.getElementById("parent-1").contains(document.getElementById("child"))`,
				),
			).To(BeTrue())
		})
		It("Should return false, when passed a sibling", func() {
			ctx := NewTestContext(
				LoadHTML(`<div id="parent-1"></div><div id="parent-2"></div></div>`),
			)
			Expect(
				ctx.RunTestScript(
					`document.getElementById("parent-1").contains(document.getElementById("parent-2"))`,
				),
			).To(BeFalse())
		})
	})
})
