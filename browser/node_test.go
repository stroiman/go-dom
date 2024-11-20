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
			Expect(doc.Body().InsertBefore(newElm, div)).To(Succeed())
			Expect(
				doc.Body(),
			).To(HaveOuterHTML(`<body><div>First</div><p></p><div id="1">1</div></body>`))
			Expect(newElm.Parent()).To(Equal(div.Parent()))
		})
	})
})
