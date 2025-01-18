package suite

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/html"
)

func (suite *ScriptTestSuite) CreateDocumentTests() {
	prefix := suite.Prefix

	Describe(prefix+"Document", func() {
		var ctx html.ScriptContext

		BeforeEach(func() {
			ctx = suite.NewContext()
		})

		Describe("`createElement`", func() {
			It("Should return an HTMLElement", func() {
				Expect(
					ctx.Eval(`const base = document.createElement("base");
						base instanceof HTMLElement`),
				).To(BeTrue())
			})

			It("Should exist on the prototype", func() {
				Expect(ctx.Eval(`
				Object.getOwnPropertyNames(Document.prototype).includes("createElement")
			`)).To(BeTrue())
				Expect(ctx.Eval(`
				Object.getOwnPropertyNames(Document).includes("createElement")
			`)).To(BeFalse())
			})
		})

		Describe("Instance properties", func() {
			It("Has a `location` property", func() {
				Expect(
					ctx.Eval("Object.getOwnPropertyNames(document)"),
				).To(ContainElements("location"))
			})
		})

		Describe("getElementById", func() {
			It("Should return the right element", func() {
				ctx := suite.LoadHTML(
					`<body><div id='elm-1'>Elm: 1</div><div id='elm-2'>Elm: 2</div></body>`,
				)
				Expect(ctx.Eval(`
          const e = document.getElementById("elm-2")
          e.outerHTML
        `)).To(Equal(`<div id="elm-2">Elm: 2</div>`))
				Expect(
					ctx.Eval(`Object.getPrototypeOf(e).constructor.name`),
				).To(Equal("HTMLDivElement"))
			})
		})

		itShouldBeADocument := func() {
			It("Should be an instance of Document", func() {
				Expect(ctx.Eval("actual instanceof Document")).To(BeTrue())
			})
			It("Should have nodeType 9", func() {
				Expect(ctx.Eval("actual.nodeType")).To(BeEquivalentTo(9))
			})
			It("Should be an instance of Node", func() {
				Expect(ctx.Eval("actual instanceof Node")).To(BeTrue())
			})
			It("Should be an instance of EventTarget", func() {
				Expect(ctx.Eval("actual instanceof EventTarget")).To(BeTrue())
			})
			It("Should be an instance of Object", func() {
				Expect(ctx.Eval("actual instanceof Object")).To(BeTrue())
			})
		}

		itShouldBeAnHTMLDocument := func(expectHTMLDocument bool) {
			It("Should be an instance of HTMLDocument", func() {
				Expect(
					ctx.Eval("actual instanceof HTMLDocument"),
				).To(Equal(expectHTMLDocument))
			})

			itShouldBeADocument()

			It("Should have a class hierarchy of 5 classes", func() {
				expectedBaseClassCount := 4
				if expectHTMLDocument {
					expectedBaseClassCount++
				}
				Expect(ctx.Eval(`
        let baseClassCount = 0
        let current = actual
        while(current = Object.getPrototypeOf(current))
          baseClassCount++
        baseClassCount;
      `)).To(BeEquivalentTo(expectedBaseClassCount))
			})
		}

		Describe("new Document()", func() {
			BeforeEach(func() {
				Expect(ctx.Run("const actual = new Document()")).To(Succeed())
			})

			It("Should now be the same as window.document", func() {
				Expect(ctx.Eval("actual === window.document")).To(BeFalse())
			})

			Describe("Class hierarchy", func() {
				itShouldBeAnHTMLDocument(false)
			})
		})

		Describe("Class Hierarchy of `window.document`", func() {
			BeforeEach(func() {
				ctx = suite.LoadHTML("<html></html>")
				ctx.Run("const actual = window.document")
			})
			itShouldBeAnHTMLDocument(true)
		})
	})
}
