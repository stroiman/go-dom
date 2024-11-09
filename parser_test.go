package go_dom_test

import (
	"fmt"
	"net/http"
	"strings"

	. "github.com/stroiman/go-dom"
	dom "github.com/stroiman/go-dom/dom-types"
	"github.com/stroiman/go-dom/interfaces"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
)

func parseString(s string) interfaces.Document {
	return Parse(strings.NewReader(s))
}

var _ = Describe("Parser", func() {
	It("Should be able to parse an empty HTML document", func() {
		result, ok := (parseString("<html><head></head><body></body></html>")).(*dom.Document)
		Expect(ok).To(BeTrue())
		element := result.DocumentElement()
		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
		Expect(result).To(
			MatchStructure("HTML",
				MatchStructure("HEAD"),
				MatchStructure("BODY"),
			))
	})

	It("Should wrap contents in an HTML element if missing", func() {
		result, ok := (parseString("<head></head><body></body>")).(*dom.Document)
		Expect(ok).To(BeTrue())
		element := result.DocumentElement()
		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
		Expect(result).To(
			MatchStructure("HTML",
				MatchStructure("HEAD"),
				MatchStructure("BODY")))
	})

	It("Should create a HEAD if missing", func() {
		result, ok := (parseString("<html><body></body></html>")).(*dom.Document)
		Expect(ok).To(BeTrue())
		Expect(result).To(
			MatchStructure("HTML",
				MatchStructure("HEAD"),
				MatchStructure("BODY")))
	})

	It("Should create HTML and HEAD if missing", func() {
		result, ok := (parseString("<body></body>")).(*dom.Document)
		Expect(ok).To(BeTrue())
		Expect(result).To(
			MatchStructure("HTML",
				MatchStructure("HEAD"),
				MatchStructure("BODY")))
	})

	It("Should embed a root <div> in an HTML document structure", func() {
		result, ok := (parseString("<div></div>")).(*dom.Document)
		Expect(ok).To(BeTrue())
		Expect(result).To(
			MatchStructure("HTML",
				MatchStructure("HEAD"),
				MatchStructure("BODY",
					MatchStructure("DIV"),
				)))
	})

	It("Should be able to read from an http.Handler instance", func() {
		handler := (http.HandlerFunc)(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html") // For good measure, not used yet"
			w.Write([]byte("<html></html>"))
		})
		browser := NewBrowserFromHandler(handler)
		result, ok := browser.Open("/").(*dom.Document)
		Expect(ok).To(BeTrue())
		element := result.DocumentElement()
		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
	})
})

func MatchStructure(name string, children ...types.GomegaMatcher) types.GomegaMatcher {
	return WithTransform(func(node interface{}) (res struct {
		Name     string
		Children []*dom.Element
	}) {
		var element *dom.Element
		switch elm := node.(type) {
		case *dom.Document:
			element = elm.DocumentElement()
		case *dom.Element:
			element = elm
		default:
			panic(fmt.Sprintf("Unknown type %T for element", elm))
		}
		res.Name = element.TagName()
		res.Children = element.ChildNodes()
		return
	}, And(
		HaveField("Name", Equal(name)),
		HaveField("Children", HaveExactElements(children)),
	))
}
