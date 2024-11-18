package scripting_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
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
	return c.Run(script)
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

func InitializeContext(hooks ...CreateHook) *TestScriptContext {
	ctx := TestScriptContext{}

	BeforeEach(func() {
		window := browser.NewWindow()
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

func InitializeContextWithEmptyHtml() *TestScriptContext {
	return InitializeContext(
		func(ctx *ScriptContext) {
			ctx.Window().LoadHTML("<html></html>") // Still creates head and body element
		})
}
