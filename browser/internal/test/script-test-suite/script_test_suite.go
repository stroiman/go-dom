package suite

import (
	"github.com/onsi/ginkgo/v2"
	"github.com/stroiman/go-dom/browser/html"
)

type ScriptTestSuite struct {
	Engine html.ScriptHost
	Prefix string
}

type ScriptTestContext struct {
	Window html.Window
}

func (ctx *ScriptTestContext) Eval(script string) (any, error) {
	return ctx.Window.Eval(script)
}

func (ctx *ScriptTestContext) Run(script string) error {
	return ctx.Window.Run(script)
}

func (ctx *ScriptTestContext) Dispose() {
	ctx.Window.Dispose()
}

func NewScriptTestSuite(
	engine html.ScriptHost,
	prefix string) *ScriptTestSuite {
	return &ScriptTestSuite{engine, prefix + ": "}
}

func (suite *ScriptTestSuite) NewContext() *ScriptTestContext {
	options := html.WindowOptions{
		ScriptHost: suite.Engine,
	}
	result := &ScriptTestContext{
		Window: html.NewWindow(options),
	}
	ginkgo.DeferCleanup(func() { result.Dispose() })
	return result
}

func (suite *ScriptTestSuite) NewWindow() html.Window {
	return suite.NewContext().Window
}

func (suite *ScriptTestSuite) CreateAllGinkgoTests() {
	suite.CreateWindowTests()
	suite.CreateDocumentTests()
	suite.CreateEventTargetTests()
}
