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

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("toJSON", v8.NewFunctionTemplateWithError(iso, wrapper.ToJSON))

	prototypeTmpl.SetAccessorProperty("href",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetHref),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHref),
		v8.None)
	prototypeTmpl.SetAccessorProperty("origin",
		v8.NewFunctionTemplateWithError(iso, wrapper.Origin),
		nil,
		v8.ReadOnly)
	prototypeTmpl.SetAccessorProperty("protocol",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetProtocol),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetProtocol),
		v8.None)
	prototypeTmpl.SetAccessorProperty("username",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetUsername),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetUsername),
		v8.None)
	prototypeTmpl.SetAccessorProperty("password",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetPassword),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetPassword),
		v8.None)
	prototypeTmpl.SetAccessorProperty("host",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetHost),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHost),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hostname",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetHostname),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHostname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("port",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetPort),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetPort),
		v8.None)
	prototypeTmpl.SetAccessorProperty("pathname",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetPathname),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetPathname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("search",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetSearch),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetSearch),
		v8.None)
	prototypeTmpl.SetAccessorProperty("searchParams",
		v8.NewFunctionTemplateWithError(iso, wrapper.SearchParams),
		nil,
		v8.ReadOnly)
	prototypeTmpl.SetAccessorProperty("hash",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetHash),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHash),
		v8.None)

	return constructor
}

func (u ESURL) NewInstance(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	url, err1 := TryParseArg(args, 0, u.DecodeUSVString)
	base, err2 := TryParseArg(args, 1, u.DecodeUSVString)
	ctx := u.host.MustGetContext(info.Context())
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err1, err2)
		if err != nil {
			return nil, err
		}
		return u.CreateInstanceBase(ctx, info.This(), url, base)
	}
	if args.noOfReadArguments >= 1 {
		if err1 != nil {
			return nil, err1
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
	result, callErr := instance.ToJSON()
	if callErr != nil {
		return nil, callErr
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
	return nil, errors.New("Not implemented: URL.SetHref")
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
	return nil, errors.New("Not implemented: URL.SetProtocol")
}

func (u ESURL) GetUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: URL.GetUsername")
}

func (u ESURL) SetUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: URL.SetUsername")
}

func (u ESURL) GetPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: URL.GetPassword")
}

func (u ESURL) SetPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: URL.SetPassword")
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
	return nil, errors.New("Not implemented: URL.SetHost")
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
	return nil, errors.New("Not implemented: URL.SetHostname")
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
	return nil, errors.New("Not implemented: URL.SetPort")
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
	return nil, errors.New("Not implemented: URL.SetPathname")
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
	return nil, errors.New("Not implemented: URL.SetSearch")
}

func (u ESURL) SearchParams(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: URL.SearchParams")
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
	return nil, errors.New("Not implemented: URL.SetHash")
}
