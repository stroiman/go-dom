package browser_test

import (
	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Document", func() {
	It("Has an outerHTML", func() {
		// Parsing an empty HTML doc generates both head and body
		doc := ParseHtmlString("")
		Expect(
			doc.DocumentElement().OuterHTML(),
		).To(Equal("<html><head></head><body></body></html>"))
	})
})
