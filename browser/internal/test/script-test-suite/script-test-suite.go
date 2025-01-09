package suite

import "github.com/stroiman/go-dom/browser/html"

type ScriptTestSuite struct {
	Engine html.ScriptEngineFactory
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

func NewScriptTestSuite(
	engine html.ScriptEngineFactory,
	prefix string) *ScriptTestSuite {
	return &ScriptTestSuite{engine, prefix + ": "}
}

func (suite *ScriptTestSuite) NewWindow() html.Window {
	options := html.WindowOptions{
		ScriptEngineFactory: suite.Engine,
	}
	return html.NewWindow(options)
}

func (suite *ScriptTestSuite) CreateAllGinkgoTests() {
	suite.CreateWindowTests()
	suite.CreateDocumentTests()
}
