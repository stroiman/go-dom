package htmlelements_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/code-gen/html-elements"
)

var _ = Describe("IdlInterface", func() {
	It("Should generate an interface", func() {
		actual := IdlInterface{
			Name:       "HTMLAnchorElement",
			Attributes: []IdlInterfaceAttribute{NewStringAttribute("target")},
		}
		Expect(actual).To(HaveRendered(
			`type HTMLAnchorElement interface {
	Target() string
	SetTarget(string)
}`))
	})

	It("Should add inherited type", func() {
		actual := IdlInterface{
			Name:       "HTMLAnchorElement",
			Inherits:   "HTMLElement",
			Attributes: []IdlInterfaceAttribute{NewStringAttribute("target")},
		}
		Expect(actual).To(HaveRendered(
			`type HTMLAnchorElement interface {
	HTMLElement
	Target() string
	SetTarget(string)
}`))
	})

	It("Should not sanitize Type", func() {
		actual := IdlInterface{
			Name:       "HTMLAnchorElement",
			Attributes: []IdlInterfaceAttribute{NewStringAttribute("type")},
		}
		Expect(actual).To(HaveRendered(
			`type HTMLAnchorElement interface {
	Type() string
	SetType(string)
}`))

	})
})
