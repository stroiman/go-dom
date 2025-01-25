package htmlelements_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	. "github.com/stroiman/go-dom/code-gen/generators"
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

func Render(expected string) types.GomegaMatcher {
	return WithTransform(
		// func(g Generator) string { return GeneratorStringer{Generator: g}.String() },
		func(g Generator) string { return fmt.Sprintf("%#v", g.Generate()) },
		Equal(expected),
	)
}
