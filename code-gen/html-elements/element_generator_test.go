package htmlelements_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/code-gen/html-elements"
)

var _ = Describe("ElementGenerator", func() {
	It("Should generate a getter and setter", func() {
		Expect(GenerateHtmlAnchor()).To(Render(
			`func (e *htmlAnchorElement) Target() string {
	return e.target
}

func (e *htmlAnchorElement) SetTarget(val string) {
	e.target = val
}`))
	})
})
