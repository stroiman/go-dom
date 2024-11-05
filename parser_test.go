package go_dom_test

import (
	"strings"

	. "github.com/stroiman/go-dom"
	//. "github.com/stroiman/go-dom/dom-types"
	"github.com/stroiman/go-dom/interfaces"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func parseString(s string) interfaces.Node {
	return Parse(strings.NewReader(s))
}

var _ = Describe("Parser", func() {
	It("Should start with a test", func() {
		result := parseString("<html></html>")
		element := result.(interfaces.Element)
		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
	})
})
