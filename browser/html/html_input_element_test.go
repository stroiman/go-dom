package html_test

import (
	"github.com/stroiman/go-dom/browser/html"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTMLInputElement", func() {
	Describe("Type", func() {
		It("Has a default value of 'text'", func() {
			e := html.NewHTMLInputElement(nil)
			Expect(e.Type()).To(Equal("text"))
		})
	})
})
