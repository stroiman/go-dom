package browser_test

import (
	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Document", func() {
	It("Is an HTMLElement", func() {
		doc := ParseHtmlString("")
		Expect(doc.DocumentElement()).To(BeHTMLElement())
	})

	It("Has an outerHTML", func() {
		// Parsing an empty HTML doc generates both head and body
		doc := ParseHtmlString("")
		Expect(
			doc.DocumentElement().OuterHTML(),
		).To(Equal("<html><head></head><body></body></html>"))
	})

	Describe("FindElementById", func() {
		It("Should return the right element", func() {
			doc := ParseHtmlString(`<body>
  <div id="uncle></div>
  <div id="parent">
    <div id="child">
      <div id="dummy"></div>
      <div id="grand-child"></div>
    </div>
  </div></body>`)
			elm := doc.GetElementById("grand-child")
			Expect(elm).To(HaveOuterHTML(`<div id="grand-child"></div>`))
		})
	})
})

func BeHTMLElement() OmegaMatcher { return HtmlElementMatcher{} }

type HtmlElementMatcher struct{}

func (m HtmlElementMatcher) Match(value any) (bool, error) {
	_, ok := value.(HTMLElement)
	return ok, nil
}

func (m HtmlElementMatcher) FailureMessage(actual interface{}) (message string) {
	return "Should be en HTMLElement"
}

func (m HtmlElementMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return "Should not be an HTMLElement"
}
