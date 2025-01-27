package dom_test

import (
	"errors"

	. "github.com/stroiman/go-dom/browser/dom"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
)

var _ = Describe("URL", func() {
	// Most of this is implicitly tested through Location tests.
	Describe("Parsing valid url", func() {
		It("ParseURL return no error on a valid URL", func() {
			result := ParseURL("http://example.com")
			Expect(result).To(HaveHRef(Equal("http://example.com")))
		})

		It("CanParseURL returns true", func() {
			Expect(CanParseURL("http://example.com")).To(BeTrue())
		})
	})

	Describe("Parsing invalid url", func() {
		It("NewUrlBase returns error on an invalid URL", func() {
			Expect(NewUrl(":foo")).Error().To(HaveOccurred())
		})

		It("ParseURL returns nil", func() {
			Expect(ParseURL(":foo")).To(BeNil())
		})

		It("CanParseURL returns false", func() {
			Expect(CanParseURL(":foo")).To(BeFalse())
		})
	})

	Describe("Constructing with a base", func() {
		// Following examples are taken from MDN documentation
		// https://developer.mozilla.org/en-US/docs/Web/API/URL_API/Resolving_relative_references
		It("Should handle example 1", func() {
			url := "articles"
			base := "https://developer.mozilla.org/some/path"
			Expect(
				NewUrlBase(url, base),
			).To(HaveHref("https://developer.mozilla.org/some/articles"), "new URL")
			Expect(
				ParseURLBase(url, base),
			).To(HaveHref("https://developer.mozilla.org/some/articles"), "URL.parse")
		})

		It("Should handle example 2", func() {
			Expect(NewUrlBase(
				"./article",
				"https://test.example.org/api/",
			)).To(HaveHref("https://test.example.org/api/article"))
			Expect(
				NewUrlBase("article", "https://test.example.org/api/v1")).To(HaveHref(
				"https://test.example.org/api/article"))
		})

		It("Should handle example 3", func() {
			Expect(
				NewUrlBase("./story/", "https://test.example.org/api/v2/"),
			).To(HaveHref("https://test.example.org/api/v2/story/"))
			Expect(
				NewUrlBase("./story", "https://test.example.org/api/v2/v3"),
			).To(HaveHref("https://test.example.org/api/v2/story"))

		})
		It("Should handle example 4 - parent directory relative", func() {
			Expect(
				NewUrlBase("../path", "https://test.example.org/api/v1/v2/"),
			).To(HaveHref("https://test.example.org/api/v1/path"))
			Expect(
				NewUrlBase("../../path", "https://test.example.org/api/v1/v2/v3"),
			).To(HaveHref("https://test.example.org/api/path"))
			Expect(
				NewUrlBase("../../../../path", "https://test.example.org/api/v1/v2/"),
			).To(HaveHref("https://test.example.org/path"))
		})

		It("Should handle example 5 - root relative", func() {
			Expect(
				NewUrlBase("/some/path", "https://test.example.org/api/"),
			).To(HaveHref("https://test.example.org/some/path"))
			Expect(
				NewUrlBase("/", "https://test.example.org/api/v1/"),
			).To(HaveHref("https://test.example.org/"))
			Expect(
				NewUrlBase("/article", "https://example.com/api/v1/"),
			).To(HaveHref("https://example.com/article"))
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
	Describe("Search", func() {
		It("Should accept both with and without a ?", func() {
			u := ParseURL("https://example.com:1234/path/name?foo=bar")
			u.SetSearch("q=help")
			Expect(u.Search()).To(Equal("?q=help"))
			u.SetSearch("?q=help")
			Expect(u.Search()).To(Equal("?q=help"))
		})
	})

	Describe("Port", func() {
		It("Should remove the host part when set to empty string", func() {
			u := ParseURL("https://example.com:1234/path/name?foo=bar")
			u.SetPort("")
			Expect(u.Host()).To(Equal("example.com"))
			Expect(u.Hostname()).To(Equal("example.com"))

			u.SetPort("1234")
			Expect(u.Host()).To(Equal("example.com:1234"))
			Expect(u.Hostname()).To(Equal("example.com"))
		})
	})

	Describe("Hostname", func() {
		It("Should keep the same port", func() {
			u := ParseURL("https://example.com:1234/path/name?foo=bar")
			u.SetHostname("m.example.com")
			Expect(u.Host()).To(Equal("m.example.com:1234"))
			Expect(u.Hostname()).To(Equal("m.example.com"))
		})
	})
})

func HaveHRef(expected interface{}) types.GomegaMatcher {
	if m, ok := expected.(types.GomegaMatcher); ok {
		return WithTransform(func(u URL) string { return u.Href() }, m)
	} else {
		return HaveHRef(Equal(expected))
	}
}

func HaveHref(expected string) types.GomegaMatcher {
	return gcustom.MakeMatcher(func(u URL) (bool, error) {
		if u == nil {
			return false, errors.New("URL is nil")
		}
		return u.Href() == expected, nil
	})
}
