package browser_test

import (
	. "github.com/stroiman/go-dom/browser"

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
			Skip("TODO")
			doc := ParseHtmlString(`<body><div>First</div><div id="1">1</div></body>`)
			newElm := doc.CreateElement("p")
			Expect(doc.Body().InsertBefore(newElm, nil)).To(Succeed())
			Expect(
				doc.Body(),
			).To(HaveOuterHTML(`<body><div>First</div><div id="1">1</div><p></p></body>`))
			Expect(newElm.Parent()).To(Equal(doc.Body()))
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
})
