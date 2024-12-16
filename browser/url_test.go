package browser_test

import (
	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

var _ = Describe("URL", func() {
	// Most of this is implicitly tested through Location tests.
	Describe("Parsing valid url", func() {
		It("ParseURL return no error on a valid URL", func() {
			result, err := ParseURL("http://example.com")
			Expect(err).ToNot(HaveOccurred())
			Expect(result).To(HaveHRef(Equal("http://example.com")))
		})

		It("CanParseURL returns true", func() {
			Expect(CanParseURL("http://example.com")).To(BeTrue())
		})
	})

	Describe("Parsing invalid url", func() {
		It("ParseURL returns error on an invalid URL", func() {
			Expect(ParseURL(":foo")).Error().To(HaveOccurred())
		})

		It("CanParseURL returns false", func() {
			Expect(CanParseURL(":foo")).To(BeFalse())
		})
	})

	Describe("Not implemented functions", func() {
		It("Returns correct error on CreateObjectURL", func() {
			_, err := CreateObjectURL("")
			Expect(err).To(HaveOccurred())
			Expect(IsNotImplementedError(err)).To(BeTrue())
		})
		It("Returns correct error on RevokeObjectURL", func() {
			_, err := CreateObjectURL("")
			Expect(err).To(HaveOccurred())
			Expect(IsNotImplementedError(err)).To(BeTrue())
		})
	})
})

func HaveHRef(matcher types.GomegaMatcher) types.GomegaMatcher {
	return WithTransform(func(u URL) string { return u.GetHref() }, matcher)
}
