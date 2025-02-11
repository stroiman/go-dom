// This file is generated. Do not edit.

package v8host

import (
	"errors"
	log "github.com/gost-dom/browser/internal/log"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("DOMTokenList", "", createDOMTokenListPrototype)
}

func createDOMTokenListPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newDOMTokenListV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	wrapper.installPrototype(constructor.PrototypeTemplate())

	wrapper.CustomInitialiser(constructor)
	return constructor
}
func (w domTokenListV8Wrapper) installPrototype(prototypeTmpl *v8.ObjectTemplate) {
	iso := w.scriptHost.iso
	prototypeTmpl.Set("item", v8.NewFunctionTemplateWithError(iso, w.item))
	prototypeTmpl.Set("contains", v8.NewFunctionTemplateWithError(iso, w.contains))
	prototypeTmpl.Set("add", v8.NewFunctionTemplateWithError(iso, w.add))
	prototypeTmpl.Set("remove", v8.NewFunctionTemplateWithError(iso, w.remove))
	prototypeTmpl.Set("toggle", v8.NewFunctionTemplateWithError(iso, w.toggle))
	prototypeTmpl.Set("replace", v8.NewFunctionTemplateWithError(iso, w.replace))
	prototypeTmpl.Set("supports", v8.NewFunctionTemplateWithError(iso, w.supports))

	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, w.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("value",
		v8.NewFunctionTemplateWithError(iso, w.value),
		v8.NewFunctionTemplateWithError(iso, w.setValue),
		v8.None)
}

func (w domTokenListV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(w.scriptHost.iso, "Illegal Constructor")
}

func (w domTokenListV8Wrapper) item(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: DOMTokenList.item")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	index, err1 := tryParseArg(args, 0, w.decodeUnsignedLong)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Item(index)
		return w.toNullableDOMString(ctx, result)
	}
	return nil, errors.New("DOMTokenList.item: Missing arguments")
}

func (w domTokenListV8Wrapper) contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: DOMTokenList.contains")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	token, err1 := tryParseArg(args, 0, w.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Contains(token)
		return w.toBoolean(ctx, result)
	}
	return nil, errors.New("DOMTokenList.contains: Missing arguments")
}

func (w domTokenListV8Wrapper) add(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: DOMTokenList.add")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	tokens, err1 := tryParseArg(args, 0, w.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.Add(tokens)
		return nil, callErr
	}
	return nil, errors.New("DOMTokenList.add: Missing arguments")
}

func (w domTokenListV8Wrapper) remove(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: DOMTokenList.remove")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	tokens, err1 := tryParseArg(args, 0, w.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.Remove(tokens)
		return nil, nil
	}
	return nil, errors.New("DOMTokenList.remove: Missing arguments")
}

func (w domTokenListV8Wrapper) replace(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: DOMTokenList.replace")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	token, err1 := tryParseArg(args, 0, w.decodeDOMString)
	newToken, err2 := tryParseArg(args, 1, w.decodeDOMString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result := instance.Replace(token, newToken)
		return w.toBoolean(ctx, result)
	}
	return nil, errors.New("DOMTokenList.replace: Missing arguments")
}

func (w domTokenListV8Wrapper) supports(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: DOMTokenList.supports")
	return nil, errors.New("DOMTokenList.supports: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (w domTokenListV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: DOMTokenList.length")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return w.toUnsignedLong(ctx, result)
}

func (w domTokenListV8Wrapper) value(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.mustGetContext(info)
	log.Debug("V8 Function call: DOMTokenList.value")
	instance, err := w.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return w.toDOMString(ctx, result)
}

func (w domTokenListV8Wrapper) setValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	log.Debug("V8 Function call: DOMTokenList.setValue")
	args := newArgumentHelper(w.scriptHost, info)
	instance, err0 := w.getInstance(info)
	val, err1 := tryParseArg(args, 0, w.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetValue(val)
		return nil, nil
	}
	return nil, errors.New("DOMTokenList.setValue: Missing arguments")
}
