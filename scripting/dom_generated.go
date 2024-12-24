// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func CreateDOMTokenListPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewESDOMTokenList(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.NewInstance)
	constructor.GetInstanceTemplate().SetInternalFieldCount(1)
	prototype := constructor.PrototypeTemplate()

	prototype.Set("item", v8.NewFunctionTemplateWithError(iso, wrapper.Item))
	prototype.Set("contains", v8.NewFunctionTemplateWithError(iso, wrapper.Contains))
	prototype.Set("add", v8.NewFunctionTemplateWithError(iso, wrapper.Add))
	prototype.Set("remove", v8.NewFunctionTemplateWithError(iso, wrapper.Remove))
	prototype.Set("toggle", v8.NewFunctionTemplateWithError(iso, wrapper.Toggle))
	prototype.Set("replace", v8.NewFunctionTemplateWithError(iso, wrapper.Replace))
	prototype.Set("supports", v8.NewFunctionTemplateWithError(iso, wrapper.Supports))

	prototype.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, wrapper.Length),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("value",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetValue),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetValue),
		v8.None)
	wrapper.CustomInitialiser(constructor)
	return constructor
}

func (u ESDOMTokenList) NewInstance(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(u.host.iso, "Illegal Constructor")
}

func (u ESDOMTokenList) Item(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := info.Args()
	argsLen := len(args)
	if argsLen < 1 {
		return nil, errors.New("Too few arguments")
	}
	index, err := u.GetArgUnsignedLong(args, 0)
	if err != nil {
		return nil, err
	}
	result := instance.Item(index)
	return u.ToNullableDOMString(ctx, result)
}

func (u ESDOMTokenList) Contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := info.Args()
	argsLen := len(args)
	if argsLen < 1 {
		return nil, errors.New("Too few arguments")
	}
	token, err := u.GetArgDOMString(args, 0)
	if err != nil {
		return nil, err
	}
	result := instance.Contains(token)
	return u.ToBoolean(ctx, result)
}

func (u ESDOMTokenList) Add(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := info.Args()
	argsLen := len(args)
	if argsLen < 1 {
		return nil, errors.New("Too few arguments")
	}
	tokens, err := u.GetArgDOMString(args, 0)
	if err != nil {
		return nil, err
	}
	err = instance.Add(tokens)
	return nil, err
}

func (u ESDOMTokenList) Remove(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := info.Args()
	argsLen := len(args)
	if argsLen < 1 {
		return nil, errors.New("Too few arguments")
	}
	tokens, err := u.GetArgDOMString(args, 0)
	if err != nil {
		return nil, err
	}
	instance.Remove(tokens)
	return nil, nil
}

func (u ESDOMTokenList) Replace(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := info.Args()
	argsLen := len(args)
	if argsLen < 2 {
		return nil, errors.New("Too few arguments")
	}
	token, err0 := u.GetArgDOMString(args, 0)
	newToken, err1 := u.GetArgDOMString(args, 1)
	err = errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	result := instance.Replace(token, newToken)
	return u.ToBoolean(ctx, result)
}

func (u ESDOMTokenList) Supports(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESDOMTokenList) Length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return u.ToUnsignedLong(ctx, result)
}

func (u ESDOMTokenList) GetValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetValue()
	return u.ToDOMString(ctx, result)
}

func (u ESDOMTokenList) SetValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := info.Args()
	argsLen := len(args)
	if argsLen < 1 {
		return nil, errors.New("Too few arguments")
	}
	val, err := u.GetArgDOMString(args, 0)
	if err != nil {
		return nil, err
	}
	instance.SetValue(val)
	return nil, nil
}
