package scripting_test

import (
	"fmt"

	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/scripting"
	"github.com/tommie/v8go"
)

type TestScriptContext struct {
	*ScriptContext
}

func (c TestScriptContext) RunTestScript(script string) any {
	result, err := c.RunScript(script)
	Expect(
		err,
	).ToNot(HaveOccurred(),
		fmt.Sprintf("Script error. Script src:\n-----\n%s\n-----\n", script),
	)
	return v8ValueToGoValue(result)
}

func v8ValueToGoValue(result *v8go.Value) interface{} {
	if result.IsBoolean() {
		return result.Boolean()
	}
	if result.IsInt32() {
		return result.Int32()
	}
	panic(fmt.Sprintf("Value not yet supported: %v", *result))
}
