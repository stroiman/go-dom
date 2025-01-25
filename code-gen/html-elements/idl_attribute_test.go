package htmlelements_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	g "github.com/stroiman/go-dom/code-gen/generators"
	. "github.com/stroiman/go-dom/code-gen/html-elements"
)

var _ = Describe("IDLAttribute", func() {
	It("Should generate a getter and setter", func() {
		Expect(IDLAttribute{
			AttributeName: "target",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
		}).To(HaveRendered(
			`func (e *htmlAnchorElement) Target() string {
	return e.target
}

func (e *htmlAnchorElement) SetTarget(val string) {
	e.target = val
}`))
	})

	It("Should generate a getter when readOnly", func() {
		Expect(IDLAttribute{
			AttributeName: "target",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
			ReadOnly: true,
		}).To(HaveRendered(
			`func (e *htmlAnchorElement) Target() string {
	return e.target
}`))
	})
})
