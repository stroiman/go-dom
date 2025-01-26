package htmlelements_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	g "github.com/stroiman/go-dom/code-gen/generators"
	. "github.com/stroiman/go-dom/code-gen/html-elements"
)

func GenerateHtmlAnchor() (g.Generator, error) {
	return GenerateHTMLElement("HTMLAnchorElement")
}

var _ = Describe("ElementGenerator", func() {
	It("Should generate a getter and setter", func() {
		Expect(GenerateHtmlAnchor()).To(HaveRendered(ContainSubstring(
			`func (e *htmlAnchorElement) Target() string {`)))
	})

	It("Should generate a struct with embedded htmlElement", func() {
		Expect(GenerateHtmlAnchor()).To(HaveRendered(MatchRegexp(
			`type htmlAnchorElement struct {\n\tHTMLElement`)))
	})

	It("Should generate a constructor", func() {
		Expect(GenerateHtmlAnchor()).To(HaveRendered(ContainSubstring(
			`func NewHTMLAnchorElement(ownerDoc HTMLDocument) HTMLAnchorElement {
	result := &htmlAnchorElement{NewHTMLElement("a", ownerDoc)}
	result.SetSelf(result)
	return result
}`)))
	})
})
