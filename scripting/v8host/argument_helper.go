package v8host

import (
	"errors"
	"fmt"

	v8 "github.com/tommie/v8go"
)

type argumentHelper struct {
	*v8.FunctionCallbackInfo
	ctx               *V8ScriptContext
	noOfReadArguments int
}

func newArgumentHelper(host *V8ScriptHost, info *v8.FunctionCallbackInfo) *argumentHelper {
	ctx := host.mustGetContext(info.Context())
	return &argumentHelper{info, ctx, 0}
}

func (h argumentHelper) getFunctionArg(index int) (*v8.Function, error) {
	args := h.Args()
	if index >= len(args) {
		return nil, ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsFunction() {
		return arg.AsFunction()
	}
	return nil, h.newTypeError("Expected function", arg)
}

func (h argumentHelper) getInt32Arg(index int) (int32, error) {
	args := h.Args()
	if index >= len(args) {
		return 0, ErrWrongNoOfArguments
	}
	arg := args[index]
	if arg.IsNumber() {
		return arg.Int32(), nil
	}
	return 0, h.newTypeError("Expected int32", arg)
}

func (h argumentHelper) getStringArg(index int) (string, error) {
	args := h.Args()
	if index >= len(args) {
		return "", ErrWrongNoOfArguments
	}
	arg := args[index]
	return arg.String(), nil
	// if arg.IsString() {
	// 	return arg.String(), nil
	// }
	// return "", h.newTypeError("Expected string", arg)
}

func (h *argumentHelper) getArg(index int) *v8.Value {
	args := h.FunctionCallbackInfo.Args()
	if len(args) <= index {
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

func (h *argumentHelper) newTypeError(msg string, v *v8.Value) error {
	json, _ := v8.JSONStringify(h.ctx.v8ctx, v)
	return v8.NewTypeError(
		h.ctx.host.iso,
		fmt.Sprintf("TypeError: %s\n%s", msg, json),
	)
}

var (
	ErrIncompatibleType   = errors.New("Incompatible type")
	ErrWrongNoOfArguments = errors.New("Not enough arguments passed")
)
