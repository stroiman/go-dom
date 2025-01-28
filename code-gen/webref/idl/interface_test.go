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

	Describe("Includes", func() {
		It("Should have right operations on url.URL", func() {
			idl, err := LoadIdlParsed("url")
			Expect(err).ToNot(HaveOccurred())
			ops := idl.Interfaces["URL"].Operations
			Expect(ops).To(ContainElement(HaveField("Name", "toJSON")))
			Expect(ops).To(ContainElement(SatisfyAll(
				HaveField("Name", "parse"),
				HaveField("Static", true)),
			))
		})
	})
})
