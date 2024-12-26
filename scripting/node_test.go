package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 Node", func() {
	It("Should support removeChild", func() {
		ctx := NewTestContext(LoadHTML(`<div id="parent-1"><div id="child">child</div></div>`))
		Expect(ctx.RunTestScript(`
			const child = document.getElementById('child');
			const parent = document.getElementById('parent-1')
			parent.removeChild(child)
		`)).Error().ToNot(HaveOccurred())
		Expect(
			ctx.Window().Document().GetElementById("parent-1").ChildNodes().Length(),
		).To(Equal(0))
	})
	Describe("firstChild", func() {
		It("Returns the correct node", func() {
			ctx := NewTestContext(LoadHTML(`<div id="parent-1"><div id="child">child</div></div>`))
			Expect(
				ctx.RunTestScript(
					`document.getElementById("parent-1").firstChild.getAttribute("id")`,
				),
			).To(Equal("child"))
		})
	})

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
