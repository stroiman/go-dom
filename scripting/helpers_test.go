package scripting_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/stroiman/go-dom/scripting"
	"github.com/tommie/v8go"
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
	result, err := c.RunScript(script)
	if err == nil {
		return v8ValueToGoValue(result), nil
	} else {
		return nil, err
	}
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

func v8ValueToGoValue(result *v8go.Value) interface{} {
	if result == nil {
		return nil
	}
	if result.IsBoolean() {
		return result.Boolean()
	}
	if result.IsInt32() {
		return result.Int32()
	}
	if result.IsString() {
		return result.String()
	}
	if result.IsNull() {
		return nil
	}
	if result.IsUndefined() {
		return nil
	}
	panic(fmt.Sprintf("Value not yet supported: %v", *result))
}

type CreateHook func(ctx *ScriptContext)

func InitializeContext(hooks ...CreateHook) *TestScriptContext {
	ctx := TestScriptContext{}

	BeforeEach(func() {
		ctx.ScriptContext = host.NewContext()
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
