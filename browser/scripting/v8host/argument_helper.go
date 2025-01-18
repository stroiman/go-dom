package v8host

import (
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
	return nil, ErrIncompatibleType
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
	return 0, ErrIncompatibleType
}

func (h argumentHelper) getStringArg(index int) (string, error) {
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
