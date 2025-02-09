package htmlelements_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/gost-dom/code-gen/html-elements"
	g "github.com/gost-dom/generators"
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
