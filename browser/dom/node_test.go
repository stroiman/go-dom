package dom_test

import (
	. "github.com/stroiman/go-dom/browser/dom"
	. "github.com/stroiman/go-dom/browser/testing/gomega-matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Node", func() {
	Describe("InsertBefore", func() {
		It("Should insert the element if it's a new element", func() {
			doc := ParseHtmlString(`<body><div>First</div><div id="1">1</div></body>`)
			div := doc.GetElementById("1")
			Expect(div).ToNot(BeNil())
			newElm := doc.CreateElement("p")
			Expect(doc.Body().InsertBefore(newElm, div)).Error().To(Succeed())
			Expect(
				doc.Body(),
			).To(HaveOuterHTML(`<body><div>First</div><p></p><div id="1">1</div></body>`))
			Expect(newElm.Parent()).To(Equal(doc.Body()))
		})

		It("Should append the element if the reference is nil", func() {
			doc := ParseHtmlString(`<body><div>First</div><div id="1">1</div></body>`)
			newElm := doc.CreateElement("p")
			Expect(doc.Body().InsertBefore(newElm, nil)).Error().ToNot(HaveOccurred())
			Expect(
				doc.Body(),
			).To(HaveOuterHTML(`<body><div>First</div><div id="1">1</div><p></p></body>`))
			Expect(newElm.Parent()).To(Equal(doc.Body()))
		})

		Describe("Inserting a documentFragment", func() {
			It("Inserts the nodes in the right order", func() {
				doc := ParseHtmlString(`<body><div>First</div><div id="1">1</div></body>`)
				fragment := doc.CreateDocumentFragment()
				d1 := doc.CreateElement("div")
				d2 := doc.CreateElement("div")
				d1.SetAttribute("id", "c-1")
				d2.SetAttribute("id", "c-2")
				fragment.Append(d1)
				fragment.Append(d2)
				ref := doc.GetElementById("1")

				result, err := doc.Body().InsertBefore(fragment, ref)
				Expect(err).ToNot(HaveOccurred())
				Expect(
					doc.Body(),
				).To(HaveOuterHTML(`<body><div>First</div><div id="c-1"></div><div id="c-2"></div><div id="1">1</div></body>`))
				Expect(result.ChildNodes().All()).To(BeEmpty())
			})
		})

		Describe("Moving an existing node", func() {
			var doc Document

			BeforeEach(func() {
				doc = ParseHtmlString(
					`<body>
  <div id="parent-1"><div id="1">1</div><div id="2">2</div><div id="3">3</div></div>
  <div id="parent-2"><div id="ref"></div></div>
</body>`,
				)
			})

			It("Should be removed from parent when using `InsertBefore`", func() {
				elm := doc.GetElementById("2")
				ref := doc.GetElementById("ref")
				oldParent := doc.GetElementById("parent-1")
				newParent := doc.GetElementById("parent-2")
				Expect(
					oldParent,
				).To(HaveOuterHTML(`<div id="parent-1"><div id="1">1</div><div id="2">2</div><div id="3">3</div></div>`))
				Expect(newParent.InsertBefore(elm, ref)).Error().ToNot(HaveOccurred())
				Expect(
					oldParent,
				).To(HaveOuterHTML(`<div id="parent-1"><div id="1">1</div><div id="3">3</div></div>`))
				Expect(
					newParent,
				).To(HaveOuterHTML(`<div id="parent-2"><div id="2">2</div><div id="ref"></div></div>`))
			})

			It("Should be removed from parent when using `AppendChild`", func() {
				elm := doc.GetElementById("2")
				oldParent := doc.GetElementById("parent-1")
				newParent := doc.GetElementById("parent-2")
				Expect(
					oldParent,
				).To(HaveOuterHTML(`<div id="parent-1"><div id="1">1</div><div id="2">2</div><div id="3">3</div></div>`))
				newParent.AppendChild(elm)
				Expect(
					oldParent,
				).To(HaveOuterHTML(`<div id="parent-1"><div id="1">1</div><div id="3">3</div></div>`))
				Expect(
					newParent,
				).To(HaveOuterHTML(`<div id="parent-2"><div id="ref"></div><div id="2">2</div></div>`))
			})
		})
	})

	Describe("Contains", func() {
		It("Should return true for an immediate child", func() {
			doc := ParseHtmlString(
				`<body><div id="parent-1"><div id="1">1</div></div></body>`,
			)
			child := doc.GetElementById("1")
			parent := doc.GetElementById("parent-1")
			Expect(parent.Contains(child)).To(BeTrue())
		})

		It("Should return true for a grandchild", func() {
			doc := ParseHtmlString(
				`<body><div id="parent-1"><div id="1"><div id="2"><div id="3"></div></div></div></div></body>`,
			)
			child := doc.GetElementById("3")
			parent := doc.GetElementById("parent-1")
			Expect(parent.Contains(child)).To(BeTrue())
		})

		It("Should return false for a sibling", func() {
			doc := ParseHtmlString(
				`<body><div id="parent-1"><div id="1"></div></div><div id="parent-2"></div></body>`,
			)
			child := doc.GetElementById("parent-2")
			parent := doc.GetElementById("parent-1")
			Expect(parent.Contains(child)).To(BeFalse())
		})
	})

	Describe("TextContents", func() {
		It("Should return all text", func() {
			doc := ParseHtmlString(
				`<body><div style="display: none">Hidden text</div><div>Visible text</div></body>`,
			)
			text := doc.Body().GetTextContent()
			Expect(text).To(Equal("Hidden textVisible text"))
		})
	})
})
