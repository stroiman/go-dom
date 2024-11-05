package go_dom_test

import (
	"fmt"
	"strings"

	. "github.com/stroiman/go-dom"
	"github.com/stroiman/go-dom/interfaces"
	"github.com/stroiman/go-dom/lexer"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func parseString(s string) interfaces.Node {
	tokens := lexer.Tokenize(strings.NewReader(s))
	fmt.Println("Tokens", tokens)
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
