package htmlelements_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	g "github.com/gost-dom/browser/code-gen/generators"
	. "github.com/gost-dom/browser/code-gen/html-elements"
)

func GenerateURL() (g.Generator, error) {
	g, err := CreateGenerator(URLSpec)
	return g.GenerateInterface(), err
}

var _ = Describe("ElementGenerator", func() {
	It("Should generate a getter and setter", func() {
		Expect(GenerateURL()).To(HaveRendered(ContainSubstring(
			`ToJSON() (string, error)`)))
	})
})
