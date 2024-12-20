// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func CreateURLPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewESURL(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.NewInstance)
	constructor.GetInstanceTemplate().SetInternalFieldCount(1)
	prototype := constructor.PrototypeTemplate()

	prototype.Set("toJSON", v8.NewFunctionTemplateWithError(iso, wrapper.ToJSON))
	return constructor
}

func (u ESURL) NewInstance(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	url, err0 := TryParseArg(args, 0, u.DecodeUSVString)
	base, err1 := TryParseArg(args, 1, u.DecodeUSVString)
	ctx := u.host.MustGetContext(info.Context())
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		return u.CreateInstanceBase(ctx, info.This(), url, base)
	}
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0)
		if err != nil {
			return nil, err
		}
		return u.CreateInstance(ctx, info.This(), url)
	}
	return nil, errors.New("Missing arguments")
}

func (u ESURL) ToJSON(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}

	result, err := instance.ToJSON()
	if err != nil {
		return nil, err
	} else {
		return u.ToUSVString(ctx, result)
	}
}
