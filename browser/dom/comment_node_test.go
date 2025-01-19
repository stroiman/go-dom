package dom_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/dom"
)

var _ = Describe("CommentNode", func() {
	It("Should have the right node type", func() {
		node := dom.NewCommentNode("dummy")
		Expect(node.NodeType()).To(Equal(dom.NodeType(8)))
	})
})
