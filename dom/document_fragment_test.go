package dom_test

import (
	"github.com/gost-dom/browser/dom"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("DocumentFragment", func() {
	Describe("CloneNode", func() {
		It("Should create a clone of the fragment", func() {

			doc := ParseHtmlString("")
			f := doc.CreateDocumentFragment()
			dNode, _ := f.AppendChild(doc.CreateElement("div"))
			div := dNode.(dom.Element)
			div.SetAttribute("class", "foo")
			div.SetTextContent("Text")
			clone := f.CloneNode(true).(dom.DocumentFragment)

			Expect(
				clone.FirstChild().(dom.Element).OuterHTML(),
			).To(Equal(`<div class="foo">Text</div>`))
		})
	})
})
