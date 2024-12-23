package browser_test

import (
	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClassList", func() {
	var (
		el        Element
		classList DOMTokenList
	)
	BeforeEach(func() {
		doc := NewDocument(NewWindow(nil))
		el = doc.CreateElement("div")
		classList = el.ClassList()
	})

	Describe("Add", func() {
		It("Should add a new class", func() {
			Expect(classList.Add("c1", "c2")).To(Succeed())
			Expect(el.GetAttribute("class")).To(Equal("c1 c2"))
		})

		It("Should ignore existing classes", func() {
			el.SetAttribute("class", "c1 c2")
			Expect(classList.Add("c2", "c3")).To(Succeed())
			Expect(el.GetAttribute("class")).To(Equal("c1 c2 c3"))

		})

		It("Should generate a syntax error on empty string", func() {
			err := classList.Add("")
			Expect(IsSyntaxError(err)).To(BeTrue())
		})

		It("Should generate a syntax error on empty string", func() {
			err := classList.Add("")
			Expect(IsSyntaxError(err)).To(BeTrue())
		})

		It("Should generate an invalidCharacterError on empty string", func() {
			err := classList.Add("a b")
			Expect(IsInvalidCharacterError(err)).To(BeTrue())
		})
	})
})
