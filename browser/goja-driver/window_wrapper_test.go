package goja_driver_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/browser/goja-driver"
	"github.com/stroiman/go-dom/browser/html"
)

type ScriptTestSuite struct {
	Engine html.ScriptEngineFactory
	Prefix string
}

func (suite *ScriptTestSuite) NewWindow() html.Window {
	options := html.WindowOptions{
		ScriptEngineFactory: suite.Engine,
	}
	return html.NewWindow(options)
}

func (suite *ScriptTestSuite) CreateWindowTests() {
	prefix := suite.Prefix
	Describe(prefix+" - Window wrapper", func() {
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

func init() {
	suite := &ScriptTestSuite{NewGojaScriptEngine(), "goja"}
	suite.CreateWindowTests()
}
