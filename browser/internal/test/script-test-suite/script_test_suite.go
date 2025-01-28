package suite

import (
	"strings"

	"github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/gost-dom/browser/browser/html"
	matchers "github.com/gost-dom/browser/browser/testing/gomega-matchers"
)

type ScriptTestSuite struct {
	Engine  html.ScriptHost
	Prefix  string
	SkipDOM bool
}

type ScriptTestContext struct {
	Window html.Window
	matchers.ScriptMatchers
}

func (ctx *ScriptTestContext) Eval(script string) (any, error) {
	return ctx.Window.Eval(script)
}

func (ctx *ScriptTestContext) Run(script string) error {
	return ctx.Window.Run(script)
}

func (ctx *ScriptTestContext) Close() {
	ctx.Window.Close()
}

type ScriptTestSuiteOption func(*ScriptTestSuite)

var SkipDOM = func(s *ScriptTestSuite) { s.SkipDOM = true }

func NewScriptTestSuite(
	engine html.ScriptHost,
	prefix string, options ...ScriptTestSuiteOption) *ScriptTestSuite {
	result := &ScriptTestSuite{
		Engine: engine,
		Prefix: prefix + ": ",
	}
	for _, option := range options {
		option(result)
	}
	return result
}

func (suite *ScriptTestSuite) newContext(win html.Window) *ScriptTestContext {
	result := &ScriptTestContext{Window: win,
		ScriptMatchers: matchers.ScriptMatchers{
			Ctx: win.ScriptContext().(matchers.MatcherScriptContext),
		},
	}
	ginkgo.DeferCleanup(func() { result.Close() })
	return result
}

func (suite *ScriptTestSuite) NewContext() *ScriptTestContext {
	win := html.NewWindow(html.WindowOptions{
		ScriptHost: suite.Engine,
	})
	return suite.newContext(win)
}

func (suite *ScriptTestSuite) LoadHTML(h string) *ScriptTestContext {
	options := html.WindowOptions{
		ScriptHost: suite.Engine,
	}
	win, err := html.NewWindowReader(strings.NewReader(h), options)
	gomega.Expect(err).ToNot(gomega.HaveOccurred())
	return suite.newContext(win)
}

func (suite *ScriptTestSuite) NewWindow() html.Window {
	return suite.NewContext().Window
}

func (suite *ScriptTestSuite) CreateAllGinkgoTests() {
	suite.CreateWindowTests()
	suite.CreateDocumentTests()
	suite.CreateEventTargetTests()
}
