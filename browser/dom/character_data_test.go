package dom_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/gost-dom/browser/browser/dom"
)

var _ = Describe("CommentNode", func() {
	It("Should have the right node type", func() {
		node := dom.NewComment("dummy")
		Expect(node.NodeType()).To(Equal(dom.NodeType(8)))
	})

	It("Should return text and length", func() {
		node := dom.NewComment("A sequence of 27 characters")
		Expect(node.Data()).To(Equal("A sequence of 27 characters"))
		Expect(node.Length()).To(Equal(27))
	})

	It("Should return the right length for weird characters", func() {
		// This character counts for 1 character, but takes up multiple bytes
		node := dom.NewComment("𐀀")
		Expect(node.Length()).To(Equal(1))
	})
})
