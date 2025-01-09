package suite

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func (suite *ScriptTestSuite) CreateWindowTests() {
	prefix := suite.Prefix
	Describe(prefix+"Window object", func() {
		It("Should throw a TypeError when constructed", func() {
			Expect(suite.NewWindow().Eval(
				`let error;
				try { new Window() } catch(err) { 
					error = err;
				}
				error && Object.getPrototypeOf(error).constructor.name
				`)).To(Equal("TypeError"))
		})

		Describe("window.document", func() {
			It("Should have a document property", func() {
				Expect(suite.NewWindow().Eval("document")).ToNot(BeNil())
				Expect(suite.NewWindow().Eval("document instanceof Document")).To(BeTrue())
			})

			It("Is an enumerable property", func() {
				Expect(
					suite.NewWindow().
						Eval(`
							const keys = []
							for (let key in window) {
								keys.push(key);
							}
							keys.includes('document')
						`)).To(BeTrue())
			})

			It("Should return the same instance on multiple calls", func() {
				Expect(suite.NewWindow().Eval(
					`const a = window.document;
					const b = window.document;
					a === b`,
				)).To(BeTrue())
			})
		})
	})
}
