package dom_test

import (
	"fmt"
	"net/http"

	"github.com/stroiman/go-dom/browser/dom"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
)

var _ = Describe("Parser", func() {
	It("Should be able to parse an empty HTML document", func() {
		result := dom.ParseHtmlString("<html><head></head><body></body></html>")
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
		result := dom.ParseHtmlString("<head></head><body></body>")
		element := result.DocumentElement()
		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
		Expect(result).To(
			MatchStructure("HTML",
				MatchStructure("HEAD"),
				MatchStructure("BODY")))
	})

	It("Should create a HEAD if missing", func() {
		result := dom.ParseHtmlString("<html><body></body></html>")
		Expect(result).To(
			MatchStructure("HTML",
				MatchStructure("HEAD"),
				MatchStructure("BODY")))
	})

	It("Should create HTML and HEAD if missing", func() {
		result := dom.ParseHtmlString("<body></body>")
		Expect(result).To(
			MatchStructure("HTML",
				MatchStructure("HEAD"),
				MatchStructure("BODY")))
	})

	It("Should embed a root <div> in an HTML document structure", func() {
		result := dom.ParseHtmlString("<div></div>")
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
		browser := dom.NewBrowserFromHandler(handler)
		result := browser.Open("/")
		element := result.DocumentElement()

		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
	})

	It("Should parse text nodes in body", func() {
		result := dom.ParseHtmlString(`<html><body><div>
  <h1>Hello</h1>
  <p>Lorem Ipsum</p>
</div></body></html>`)
		Expect(result.Body()).To(MatchStructure("BODY", MatchStructure("DIV",
			BeTextNode("\n  "),
			MatchStructure("H1", BeTextNode("Hello")),
			BeTextNode("\n  "),
			MatchStructure("P", BeTextNode("Lorem Ipsum")),
			BeTextNode("\n"),
		)))
	})
})

func BeTextNode(content string) types.GomegaMatcher {
	return gcustom.MakeMatcher(func(actual interface{}) (bool, error) {
		_, ok := actual.(dom.TextNode)
		return ok, nil
	})
}

func MatchStructure(name string, children ...types.GomegaMatcher) types.GomegaMatcher {
	return WithTransform(func(node interface{}) (res struct {
		Name     string
		Children []dom.Node
	}) {
		var element dom.Element
		switch elm := node.(type) {
		case dom.Document:
			element = elm.DocumentElement()
		case dom.Element:
			element = elm
		default:
			panic(fmt.Sprintf("Unknown type %T for element", elm))
		}
		res.Name = element.NodeName()
		res.Children = element.ChildNodes().All()
		return
	}, And(
		HaveField("Name", Equal(name)),
		HaveField("Children", HaveExactElements(children)),
	))
}
