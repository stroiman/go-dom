package dom_test

import (
	. "github.com/stroiman/go-dom/browser/dom"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("ClassList", func() {
	var (
		el        Element
		classList DOMTokenList
	)
	BeforeEach(func() {
		doc := CreateHTMLDocument()
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

		It("Should generate a invalidCharacterError error on string with space", func() {
			err := classList.Add("a b")
			Expect(IsInvalidCharacterError(err)).To(BeTrue())
		})

		It("Should generate an invalidCharacterError on empty string", func() {
			err := classList.Add("a b")
			Expect(IsInvalidCharacterError(err)).To(BeTrue())
		})
	})

	Describe("All", func() {
		It("Should return a list of all classes", func() {
			el.SetAttribute("class", "a b c")
			classes := make([]string, 0, 3)
			for class := range classList.All() {
				classes = append(classes, class)
			}
			Expect(classes).To(ConsistOf("a", "b", "c"))
		})
	})

	Describe("Contains", func() {
		It("Should return true for an existing class", func() {
			el.SetAttribute("class", "a b c")
			Expect(classList.Contains("a")).To(BeTrue())
		})

		It("Should return false for an non-existing class", func() {
			el.SetAttribute("class", "a b c")
			Expect(classList.Contains("x")).To(BeFalse())
		})
	})

	Describe("Length", func() {
		It("Should return the number of classes", func() {
			el.SetAttribute("class", "a b c")
			Expect(classList.Length()).To(Equal(3))
		})
	})

	Describe("Get/Set Value", func() {
		It("Should read the class attribute", func() {
			el.SetAttribute("class", "a b c")
			Expect(classList.GetValue()).To(Equal("a b c"))
		})

		It("Should write the class attribute", func() {
			classList.SetValue("x y  z")
			Expect(el.GetAttribute("class")).To(Equal("x y  z"))
		})
	})

	Describe("Item", func() {
		It("Should return the item with the specified value", func() {
			el.SetAttribute("class", "a b c")
			Expect(classList.Item(1)).To(HaveValue(Equal("b")))
		})

		It("Should return nil when the index is too large", func() {
			el.SetAttribute("class", "a b c")
			Expect(classList.Item(3)).To(BeNil())
		})
	})

	Describe("Remove", func() {
		It("Should remove an existing class", func() {
			el.SetAttribute("class", "a b c")
			classList.Remove("b")
			Expect(el.GetAttribute("class")).To(Equal("a c"))
		})

		It("Should remove the last class", func() {
			el.SetAttribute("class", "a b c")
			classList.Remove("c")
			Expect(el.GetAttribute("class")).To(Equal("a b"))
		})

		It("Should leave the list intact for a non-existing class", func() {
			el.SetAttribute("class", "a b c")
			classList.Remove("x")
			Expect(el.GetAttribute("class")).To(Equal("a b c"))
		})
	})

	Describe(".Replace", func() {
		It("Should remove, insert, and remove true on existing item", func() {
			el.SetAttribute("class", "a b c")
			Expect(classList.Replace("b", "x")).To(BeTrue())
			Expect(el.GetAttribute("class")).To(Equal("a c x"))
		})

		It("Should leave the list and return false on non-existing item", func() {
			el.SetAttribute("class", "a b c")
			Expect(classList.Replace("y", "x")).To(BeFalse())
			Expect(el.GetAttribute("class")).To(Equal("a b c"))
		})
	})
})
