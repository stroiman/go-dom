package idl_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/code-gen/webref/idl"
)

var _ = Describe("Idl/Interface", func() {
	Describe("Includes", func() {
		It("Should have HTMLAnchorElement include HTMLHyperlinkElementUtils", func() {
			idl, err := LoadIdlParsed("html")
			Expect(err).ToNot(HaveOccurred())
			Expect(
				idl.Interfaces["HTMLAnchorElement"].Includes,
			).To(
				ContainElement(HaveField("Name", "HTMLHyperlinkElementUtils")),
			)
			// Equal([]string{"HTMLHyperlinkElementUtils"}))
		})
	})
})
