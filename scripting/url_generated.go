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

	prototype.SetAccessorProperty("href",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetHref),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHref),
		v8.None)
	prototype.SetAccessorProperty("origin",
		v8.NewFunctionTemplateWithError(iso, wrapper.Origin),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("protocol",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetProtocol),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetProtocol),
		v8.None)
	prototype.SetAccessorProperty("username",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetUsername),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetUsername),
		v8.None)
	prototype.SetAccessorProperty("password",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetPassword),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetPassword),
		v8.None)
	prototype.SetAccessorProperty("host",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetHost),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHost),
		v8.None)
	prototype.SetAccessorProperty("hostname",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetHostname),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHostname),
		v8.None)
	prototype.SetAccessorProperty("port",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetPort),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetPort),
		v8.None)
	prototype.SetAccessorProperty("pathname",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetPathname),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetPathname),
		v8.None)
	prototype.SetAccessorProperty("search",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetSearch),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetSearch),
		v8.None)
	prototype.SetAccessorProperty("searchParams",
		v8.NewFunctionTemplateWithError(iso, wrapper.SearchParams),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("hash",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetHash),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHash),
		v8.None)
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

func (u ESURL) GetHref(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetHref()
	return u.ToUSVString(ctx, result)
}

func (u ESURL) SetHref(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) Origin(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return u.ToUSVString(ctx, result)
}

func (u ESURL) GetProtocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetProtocol()
	return u.ToUSVString(ctx, result)
}

func (u ESURL) SetProtocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) GetUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) SetUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) GetPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) SetPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) GetHost(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetHost()
	return u.ToUSVString(ctx, result)
}

func (u ESURL) SetHost(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) GetHostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetHostname()
	return u.ToUSVString(ctx, result)
}

func (u ESURL) SetHostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) GetPort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetPort()
	return u.ToUSVString(ctx, result)
}

func (u ESURL) SetPort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) GetPathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetPathname()
	return u.ToUSVString(ctx, result)
}

func (u ESURL) SetPathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) GetSearch(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetSearch()
	return u.ToUSVString(ctx, result)
}

func (u ESURL) SetSearch(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) SearchParams(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}

func (u ESURL) GetHash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetHash()
	return u.ToUSVString(ctx, result)
}

func (u ESURL) SetHash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented")
}
