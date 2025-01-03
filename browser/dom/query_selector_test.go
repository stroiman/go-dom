package dom_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/browser/testing/gomega-matchers"
)

var _ = Describe("QuerySelector", func() {
	It("Should support these cases", func() {
		doc := ParseHtmlString("<body><div>hello</div><p>world!</p><div>Selector</div></body>")
		Expect(
			(doc.QuerySelector("div")),
		).To(HaveOuterHTML("<div>hello</div>"))
	})

	It("Should find by attribute", func() {
		doc := ParseHtmlString(
			`<body><div>hello</div><p>world!</p><div data-foo="bar">Selector</div></body>`,
		)
		Expect(
			(doc.QuerySelector("div[data-foo='bar']")),
		).To(HaveOuterHTML(`<div data-foo="bar">Selector</div>`))
	})
})
