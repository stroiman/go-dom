// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("DOMTokenList", "", createDomTokenListPrototype)
}

func createDomTokenListPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newDomTokenListV8Wrapper(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("item", v8.NewFunctionTemplateWithError(iso, wrapper.Item))
	prototypeTmpl.Set("contains", v8.NewFunctionTemplateWithError(iso, wrapper.Contains))
	prototypeTmpl.Set("add", v8.NewFunctionTemplateWithError(iso, wrapper.Add))
	prototypeTmpl.Set("remove", v8.NewFunctionTemplateWithError(iso, wrapper.Remove))
	prototypeTmpl.Set("toggle", v8.NewFunctionTemplateWithError(iso, wrapper.Toggle))
	prototypeTmpl.Set("replace", v8.NewFunctionTemplateWithError(iso, wrapper.Replace))
	prototypeTmpl.Set("supports", v8.NewFunctionTemplateWithError(iso, wrapper.Supports))

	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, wrapper.Length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("value",
		v8.NewFunctionTemplateWithError(iso, wrapper.Value),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetValue),
		v8.None)

	wrapper.CustomInitialiser(constructor)
	return constructor
}

func (u domTokenListV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(u.host.iso, "Illegal Constructor")
}

func (u domTokenListV8Wrapper) Item(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	index, err1 := TryParseArg(args, 0, u.decodeUnsignedLong)
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

func (u domTokenListV8Wrapper) Contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	token, err1 := TryParseArg(args, 0, u.decodeDOMString)
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

func (u domTokenListV8Wrapper) Add(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	tokens, err1 := TryParseArg(args, 0, u.decodeDOMString)
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

func (u domTokenListV8Wrapper) Remove(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	tokens, err1 := TryParseArg(args, 0, u.decodeDOMString)
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

func (u domTokenListV8Wrapper) Replace(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	token, err1 := TryParseArg(args, 0, u.decodeDOMString)
	newToken, err2 := TryParseArg(args, 1, u.decodeDOMString)
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

func (u domTokenListV8Wrapper) Supports(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("DOMTokenList.supports: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u domTokenListV8Wrapper) Length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return u.toUnsignedLong(ctx, result)
}

func (u domTokenListV8Wrapper) Value(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Value()
	return u.toDOMString(ctx, result)
}

func (u domTokenListV8Wrapper) SetValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.getInstance(info)
	val, err1 := TryParseArg(args, 0, u.decodeDOMString)
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
