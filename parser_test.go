package go_dom_test

import (
	"net/http"
	"strings"

	. "github.com/stroiman/go-dom"
	"github.com/stroiman/go-dom/interfaces"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func parseString(s string) interfaces.Document {
	return Parse(strings.NewReader(s))
}

var _ = Describe("Parser", func() {
	It("Should be able to parse an empty HTML document", func() {
		result := parseString("<html></html>")
		element := result.DocumentElement()
		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
	})

	It("Should be able to read from an http.Handler instance", func() {
		handler := (http.HandlerFunc)(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html") // For good measure, not used yet"
			w.Write([]byte("<html></html>"))
		})
		browser := NewBrowserFromHandler(handler)
		result := browser.Open("/")
		element := result.DocumentElement()
		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
	})
})
