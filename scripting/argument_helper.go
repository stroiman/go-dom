package scripting

import (
	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

type argumentHelper struct {
	*v8.FunctionCallbackInfo
	ctx               *ScriptContext
	noOfReadArguments int
}

func newArgumentHelper(host *ScriptHost, info *v8.FunctionCallbackInfo) argumentHelper {
	ctx := host.MustGetContext(info.Context())
	return argumentHelper{info, ctx, 0}
}

func (h argumentHelper) GetFunctionArg(index int) (*v8.Function, error) {
	args := h.Args()
	if index >= len(args) {
		return nil, ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsFunction() {
		return arg.AsFunction()
	}
	return nil, ErrIncompatibleType
}

func (h argumentHelper) GetInt32Arg(index int) (int32, error) {
	args := h.Args()
	if index >= len(args) {
		return 0, ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsNumber() {
		return arg.Int32(), nil
	}
	return 0, ErrIncompatibleType
}

func (h argumentHelper) GetStringArg(index int) (string, error) {
	args := h.Args()
	if index >= len(args) {
		return "", ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsString() {
		return arg.String(), nil
	}
	return "", ErrIncompatibleType
}

func (h argumentHelper) GetNodeArg(index int) (browser.Node, error) {
	args := h.Args()
	if index >= len(args) {
		return nil, ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsObject() {
		o := arg.Object()
		cached, ok_1 := h.ctx.GetCachedNode(o)
		if node, ok_2 := cached.(browser.Node); ok_1 && ok_2 {
			return node, nil
		}
	}
	return nil, v8.NewTypeError(h.ctx.host.iso, "Must be a node")
}

func (h *argumentHelper) GetArg(index int) *v8.Value {
	args := h.FunctionCallbackInfo.Args()
	if len(args) < index {
		return nil
	}
	arg := args[index]
	if arg.IsUndefined() {
		return nil
	}
	if h.noOfReadArguments <= index {
		h.noOfReadArguments = index + 1
	}
	return arg
}
