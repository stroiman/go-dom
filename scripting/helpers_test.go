package scripting_test

import (
	"fmt"
	"net/url"

	. "github.com/onsi/ginkgo/v2"
	"github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser"
	. "github.com/stroiman/go-dom/scripting"
)

type TestScriptContext struct {
	*ScriptContext
}

// RunTestScript wraps the [v8.RunScript] function but converts the return value
// to a Go object.
//
// If JavaScript code throws an error, an error is returned.
//
// Note: This conversion is incomplete.
func (c TestScriptContext) RunTestScript(script string) (any, error) {
	return c.Eval(script)
}

func (c TestScriptContext) MustRunTestScript(script string) any {
	result, err := c.RunTestScript(script)
	if err != nil {
		panic(
			fmt.Sprintf(
				"Script error. Script src:\n-----\n%s\n-----\nError:\n%s",
				script,
				err.Error(),
			),
		)
	}
	return result
}

type CreateHook func(ctx *ScriptContext)

// NewTextContext loads HTML into a browser for a single Ginkgo test. It
// installs the proper Ginkgo cleanup handler.
func NewTestContext(hooks ...CreateHook) TestScriptContext {
	var unhandledErrors []any
	ctx := TestScriptContext{}
	window := browser.NewWindow(new(url.URL))
	// window.LoadHTML(html)
	ctx.ScriptContext = host.NewContext(window)
	DeferCleanup(ctx.Dispose)
	for _, hook := range hooks {
		hook(ctx.ScriptContext)
	}
	DeferCleanup(func() {
		gomega.Expect(unhandledErrors).To(gomega.BeEmpty())
	})
	ctx.Window().
		AddEventListener("error",
			browser.NewEventHandlerFunc(func(e browser.Event) error {
				unhandledErrors = append(unhandledErrors, e)
				return nil
			}))
	return ctx
}

// InitializeContextWithEmptyHtml is useful when multiple tests need has the
// same initial HTML. The html will be parsed by a normal HTML parser, which
// automatically wraps content in <html> and <body> if those are missing. So you
// So passing `<div>foo</div>` will be the same as
// `<html><body><div>foo</div></body></html>`.
//
// Example:
//
//	Describe("Tests with shared setup", func () {
//		ctx := InitializeContextWithEmptyHtml(
//			"<body><div>Hello, world!</div></body>");
//
//		 It("Should should find Hello, world! in first div", func () { /*...*/ }
//		It("Should should have one child of body", func () { /*...*/ }
//	})
func InitializeContext(hooks ...CreateHook) *TestScriptContext {
	ctx := TestScriptContext{}

	BeforeEach(func() {
		window := browser.NewWindow(new(url.URL))
		ctx.ScriptContext = host.NewContext(window)
		for _, hook := range hooks {
			hook(ctx.ScriptContext)
		}
	})

	AfterEach(func() {
		ctx.Dispose()
	})

	return &ctx
}

func LoadHTML(html string) CreateHook {
	return func(ctx *ScriptContext) {
		ctx.Window().LoadHTML(html)
	}
}

func InitializeContextWithEmptyHtml() *TestScriptContext {
	return InitializeContext(LoadHTML("<html></html>"))
}

type TestBrowserWrapper struct {
	*browser.Browser
	*ScriptContext
	url string
}

func CreateTestBrowser() *TestBrowserWrapper {
	result := new(TestBrowserWrapper)

	BeforeEach(func() {

	})

	AfterEach(func() {
		result.ScriptContext.Dispose()
		// Ready for GC
		result.ScriptContext = nil
		result.Browser = nil
	})
	return result
}
