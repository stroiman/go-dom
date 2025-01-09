package suite

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func (suite *ScriptTestSuite) CreateDocumentTests() {
	prefix := suite.Prefix
	Describe(prefix+" - Document wrapper", func() {
		It("Prototype", func() {
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
}
