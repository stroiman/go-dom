// This file is generated. Do not edit.

package v8host

import (
	"errors"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("History", "", createHistoryPrototype)
}

func createHistoryPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newHistoryV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("go", v8.NewFunctionTemplateWithError(iso, wrapper.go_))
	prototypeTmpl.Set("back", v8.NewFunctionTemplateWithError(iso, wrapper.back))
	prototypeTmpl.Set("forward", v8.NewFunctionTemplateWithError(iso, wrapper.forward))
	prototypeTmpl.Set("pushState", v8.NewFunctionTemplateWithError(iso, wrapper.pushState))
	prototypeTmpl.Set("replaceState", v8.NewFunctionTemplateWithError(iso, wrapper.replaceState))

	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, wrapper.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("state",
		v8.NewFunctionTemplateWithError(iso, wrapper.state),
		nil,
		v8.None)

	return constructor
}

func (h historyV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(h.scriptHost.iso, "Illegal Constructor")
}

func (h historyV8Wrapper) go_(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.go")
	args := newArgumentHelper(h.scriptHost, info)
	instance, err0 := h.getInstance(info)
	delta, err1 := tryParseArgWithDefault(args, 0, h.defaultDelta, h.decodeLong)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.Go(delta)
		return nil, callErr
	}
	return nil, errors.New("History.go: Missing arguments")
}

func (h historyV8Wrapper) back(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.back")
	instance, err := h.getInstance(info)
	if err != nil {
		return nil, err
	}
	callErr := instance.Back()
	return nil, callErr
}

func (h historyV8Wrapper) forward(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.forward")
	instance, err := h.getInstance(info)
	if err != nil {
		return nil, err
	}
	callErr := instance.Forward()
	return nil, callErr
}

func (h historyV8Wrapper) pushState(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.pushState")
	args := newArgumentHelper(h.scriptHost, info)
	instance, err0 := h.getInstance(info)
	data, err1 := tryParseArg(args, 0, h.decodeAny)
	url, err3 := tryParseArgWithDefault(args, 2, h.defaultUrl, h.decodeUSVString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err3)
		if err != nil {
			return nil, err
		}
		callErr := instance.PushState(data, url)
		return nil, callErr
	}
	return nil, errors.New("History.pushState: Missing arguments")
}

func (h historyV8Wrapper) replaceState(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.replaceState")
	args := newArgumentHelper(h.scriptHost, info)
	instance, err0 := h.getInstance(info)
	data, err1 := tryParseArg(args, 0, h.decodeAny)
	url, err3 := tryParseArgWithDefault(args, 2, h.defaultUrl, h.decodeUSVString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err3)
		if err != nil {
			return nil, err
		}
		callErr := instance.ReplaceState(data, url)
		return nil, callErr
	}
	return nil, errors.New("History.replaceState: Missing arguments")
}

func (h historyV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := h.mustGetContext(info)
	log.Debug("V8 Function call: History.length")
	instance, err := h.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return h.toUnsignedLong(ctx, result)
}

func (h historyV8Wrapper) state(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := h.mustGetContext(info)
	log.Debug("V8 Function call: History.state")
	instance, err := h.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.State()
	return h.toJSON(ctx, result)
}
