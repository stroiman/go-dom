package suite

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func (suite *ScriptTestSuite) CreateDocumentTests() {
	prefix := suite.Prefix

	Describe(prefix+"Document", func() {
		Describe("prototype", func() {
			It("Has a `createElement`", func() {
				window := suite.NewWindow()
				Expect(window.Eval(`
				Object.getOwnPropertyNames(Document.prototype).includes("createElement")
			`)).To(BeTrue())
				Expect(window.Eval(`
				const e = Document.prototype.createElement.call(document, "div")
				typeof e
			`)).To(Equal("object"))
			})
		})

		Describe("Instance properties", func() {
			It("Has a `location` property", func() {
				ctx := suite.NewContext()
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
	})
}
