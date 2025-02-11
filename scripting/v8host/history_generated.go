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

	wrapper.installPrototype(constructor.PrototypeTemplate())

	return constructor
}
func (w historyV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("go", v8.NewFunctionTemplateWithError(iso, w.go_))
	prototypeTmpl.Set("back", v8.NewFunctionTemplateWithError(iso, w.back))
	prototypeTmpl.Set("forward", v8.NewFunctionTemplateWithError(iso, w.forward))
	prototypeTmpl.Set("pushState", v8.NewFunctionTemplateWithError(iso, w.pushState))
	prototypeTmpl.Set("replaceState", v8.NewFunctionTemplateWithError(iso, w.replaceState))

	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, w.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("state",
		v8.NewFunctionTemplateWithError(iso, w.state),
		nil,
		v8.None)
}

func (w historyV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w historyV8Wrapper) go_(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.go")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	delta, err1 := tryParseArgWithDefault(args, 0, w.defaultDelta, w.decodeLong)
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

func (w historyV8Wrapper) back(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.back")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	callErr := instance.Back()
	return nil, callErr
}

func (w historyV8Wrapper) forward(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.forward")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	callErr := instance.Forward()
	return nil, callErr
}

func (w historyV8Wrapper) pushState(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.pushState")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	data, err1 := tryParseArg(args, 0, w.decodeAny)
	url, err3 := tryParseArgWithDefault(args, 2, w.defaultUrl, w.decodeUSVString)
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

func (w historyV8Wrapper) replaceState(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: History.replaceState")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	data, err1 := tryParseArg(args, 0, w.decodeAny)
	url, err3 := tryParseArgWithDefault(args, 2, w.defaultUrl, w.decodeUSVString)
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

func (w historyV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: History.length")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return w.toUnsignedLong(ctx, result)
}

func (w historyV8Wrapper) state(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: History.state")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.State()
	return w.toJSON(ctx, result)
}
