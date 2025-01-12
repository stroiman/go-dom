package html_test

import (
	. "github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/testing/gomega-matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTMLDocument", func() {
	Describe("Empty document created by `NewDocument`", func() {
		It("Should have an HTML document element", func() {
			doc := NewHTMLDocument(nil)
			Expect(doc.DocumentElement()).To(HaveTag("HTML"))
		})

		It("Should have an empty HEAD", func() {
			doc := NewHTMLDocument(nil)
			Expect(doc.Head()).To(HaveTag("HEAD"))
		})

		It("Should have an empty BODY", func() {
			doc := NewHTMLDocument(nil)
			Expect(doc.Body()).To(HaveTag("BODY"))
		})
	})
})
