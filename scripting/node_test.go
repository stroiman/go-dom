package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 Node", func() {
	It("Should support insertBefore", func() {
		ctx := NewTestContext(
			LoadHTML(`<div id="parent-1"><div id="child-1"></div><div id="child-2"></div></div>`),
		)
		Expect(ctx.RunTestScript(`
			const f = document.createDocumentFragment()
			const d1 = document.createElement("div")
			const d2 = document.createElement("div")
			d1.setAttribute("id", "d1")
			d2.setAttribute("id", "d2")
			f.appendChild(d1)
			f.appendChild(d2)
			parent = document.getElementById("parent-1")
			ref = document.getElementById("child-2")
			parent.insertBefore(f, ref)
			Array.from(parent.childNodes).map(x => x.getAttribute("id")).join(", ")
		`)).To(Equal("child-1, d1, d2, child-2"))

	})
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
