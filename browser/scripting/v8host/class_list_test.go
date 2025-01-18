package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	matchers "github.com/stroiman/go-dom/browser/testing/gomega-matchers"
)

var _ = Describe("V8 Element.classList", func() {
	It("Should support .add", func() {
		ctx := NewTestContext(LoadHTML("<div id='target' class='c1'></div>"))
		ctx.MustRunTestScript(`
			const list = document.getElementById('target').classList;
			list.add('c2')
		`)
		element := ctx.Window().Document().GetElementById("target")
		Expect(element).To(matchers.HaveAttribute("class", "c1 c2"))
	})

	It("Should be an iterable", func() {
		ctx := NewTestContext(LoadHTML("<div id='target' class='a b c'></div>"))
		Expect(ctx.Eval(`
			list = document.getElementById("target").classList
			typeof list[Symbol.iterator]`)).To(Equal("function"))
	})

	It("Should be an iterable of class names", func() {
		ctx := NewTestContext(LoadHTML("<div id='target' class='a b c'></div>"))
		Expect(ctx.Eval(`
			const list = document.getElementById('target').classList;
			Array.from(list).join(",")
		`)).To(Equal("a,b,c"))
	})

	Describe(".toggle", func() {
		scriptContext := InitializeContext(LoadHTML(`<div id="target" class="a b c"></div>`))

		It("Removes an existing item and returns false", func() {
			Expect(scriptContext.Eval(`
				document.getElementById("target").classList.toggle("b")
			`)).To(BeFalse())
			div := scriptContext.Window().Document().GetElementById("target")
			Expect(div).To(matchers.HaveAttribute("class", "a c"))
		})

		It("Removes an adds a non-existing item and returns true", func() {
			Expect(scriptContext.Eval(`
				document.getElementById("target").classList.toggle("x")
			`)).To(BeTrue())
			div := scriptContext.Window().Document().GetElementById("target")
			Expect(div).To(matchers.HaveAttribute("class", "a b c x"))
		})

		Describe("force as true", func() {
			It("Should leave an existing item and return true", func() {
				Expect(scriptContext.Eval(`
				document.getElementById("target").classList.toggle("b", true)
			`)).To(BeTrue())
				div := scriptContext.Window().Document().GetElementById("target")
				Expect(div).To(matchers.HaveAttribute("class", "a b c"))
			})

			It("Should add a non-existing item and return true", func() {
				Expect(scriptContext.Eval(`
				document.getElementById("target").classList.toggle("x", true)
			`)).To(BeTrue())
				div := scriptContext.Window().Document().GetElementById("target")
				Expect(div).To(matchers.HaveAttribute("class", "a b c x"))
			})
		})

		Describe("force as false", func() {
			It("Should remove an existing item and return false", func() {
				Expect(scriptContext.Eval(`
				document.getElementById("target").classList.toggle("b", false)
			`)).To(BeFalse())
				div := scriptContext.Window().Document().GetElementById("target")
				Expect(div).To(matchers.HaveAttribute("class", "a c"))
			})

			It("Should ignore a non-existing item and return false", func() {
				Expect(scriptContext.Eval(`
				document.getElementById("target").classList.toggle("x", false)
			`)).To(BeFalse())
				div := scriptContext.Window().Document().GetElementById("target")
				Expect(div).To(matchers.HaveAttribute("class", "a b c"))
			})
		})
	})
})
