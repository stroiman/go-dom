package htmlelements_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	g "github.com/stroiman/go-dom/code-gen/generators"
	. "github.com/stroiman/go-dom/code-gen/html-elements"
)

var _ = Describe("IDLAttribute", func() {
	It("Should generate a getter and setter", func() {
		Skip("Temporary different solution")
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

	It("Should generate a getter and setter", func() {
		Expect(IDLAttribute{
			AttributeName: "target",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
		}).To(HaveRendered(
			`func (e *htmlAnchorElement) Target() string {
	result, _ := e.GetAttribute("target")
	return result
}

func (e *htmlAnchorElement) SetTarget(val string) {
	e.SetAttribute("target", val)
}`))
	})

	It("Should generate a getter when readOnly", func() {
		actual := IDLAttribute{
			AttributeName: "target",
			Receiver: Receiver{
				Name: g.Id("e"),
				Type: g.NewType("htmlAnchorElement").Pointer(),
			},
			ReadOnly: true,
		}
		Expect(actual).To(HaveRendered(ContainSubstring(`Target() string`)))
		Expect(actual).ToNot(HaveRendered(ContainSubstring(`SetTarget()`)))
	})
})
