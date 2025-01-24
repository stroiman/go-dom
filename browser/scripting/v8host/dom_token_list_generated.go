// This file is generated. Do not edit.

package v8host

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("DOMTokenList", "", createDomTokenListPrototype)
}

func createDomTokenListPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newDomTokenListV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("item", v8.NewFunctionTemplateWithError(iso, wrapper.item))
	prototypeTmpl.Set("contains", v8.NewFunctionTemplateWithError(iso, wrapper.contains))
	prototypeTmpl.Set("add", v8.NewFunctionTemplateWithError(iso, wrapper.add))
	prototypeTmpl.Set("remove", v8.NewFunctionTemplateWithError(iso, wrapper.remove))
	prototypeTmpl.Set("toggle", v8.NewFunctionTemplateWithError(iso, wrapper.toggle))
	prototypeTmpl.Set("replace", v8.NewFunctionTemplateWithError(iso, wrapper.replace))
	prototypeTmpl.Set("supports", v8.NewFunctionTemplateWithError(iso, wrapper.supports))

	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, wrapper.length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("value",
		v8.NewFunctionTemplateWithError(iso, wrapper.value),
		v8.NewFunctionTemplateWithError(iso, wrapper.setValue),
		v8.None)

	wrapper.CustomInitialiser(constructor)
	return constructor
}

func (u domTokenListV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(u.host.iso, "Illegal Constructor")
}

func (u domTokenListV8Wrapper) item(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	index, err1 := tryParseArg(args, 0, u.decodeUnsignedLong)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Item(index)
		return u.toNullableDOMString(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (u domTokenListV8Wrapper) contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	token, err1 := tryParseArg(args, 0, u.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Contains(token)
		return u.toBoolean(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (u domTokenListV8Wrapper) add(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	tokens, err1 := tryParseArg(args, 0, u.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.Add(tokens)
		return nil, callErr
	}
	return nil, errors.New("Missing arguments")
}

func (u domTokenListV8Wrapper) remove(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	tokens, err1 := tryParseArg(args, 0, u.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.Remove(tokens)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
}

func (u domTokenListV8Wrapper) replace(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	token, err1 := tryParseArg(args, 0, u.decodeDOMString)
	newToken, err2 := tryParseArg(args, 1, u.decodeDOMString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result := instance.Replace(token, newToken)
		return u.toBoolean(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (u domTokenListV8Wrapper) supports(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("DOMTokenList.supports: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u domTokenListV8Wrapper) length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return u.toUnsignedLong(ctx, result)
}

func (u domTokenListV8Wrapper) value(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return u.toDOMString(ctx, result)
}

func (u domTokenListV8Wrapper) setValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	val, err1 := tryParseArg(args, 0, u.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetValue(val)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
}
