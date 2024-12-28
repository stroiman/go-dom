package dom_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"

	. "github.com/stroiman/go-dom/browser/dom"
)

func HaveOuterHTML(expected string) types.GomegaMatcher {
	return WithTransform(func(node Node) (string, error) {
		if elm, ok := node.(Element); ok {
			return elm.OuterHTML(), nil
		}
		return "", fmt.Errorf("Not an element: %v", node)
	}, Equal(expected))
}

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
