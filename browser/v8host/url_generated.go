// This file is generated. Do not edit.

package v8host

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("URL", "", createUrlPrototype)
}

func createUrlPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newUrlV8Wrapper(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("toJSON", v8.NewFunctionTemplateWithError(iso, wrapper.ToJSON))

	prototypeTmpl.SetAccessorProperty("href",
		v8.NewFunctionTemplateWithError(iso, wrapper.Href),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHref),
		v8.None)
	prototypeTmpl.SetAccessorProperty("origin",
		v8.NewFunctionTemplateWithError(iso, wrapper.Origin),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("protocol",
		v8.NewFunctionTemplateWithError(iso, wrapper.Protocol),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetProtocol),
		v8.None)
	prototypeTmpl.SetAccessorProperty("username",
		v8.NewFunctionTemplateWithError(iso, wrapper.Username),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetUsername),
		v8.None)
	prototypeTmpl.SetAccessorProperty("password",
		v8.NewFunctionTemplateWithError(iso, wrapper.Password),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetPassword),
		v8.None)
	prototypeTmpl.SetAccessorProperty("host",
		v8.NewFunctionTemplateWithError(iso, wrapper.Host),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHost),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hostname",
		v8.NewFunctionTemplateWithError(iso, wrapper.Hostname),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHostname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("port",
		v8.NewFunctionTemplateWithError(iso, wrapper.Port),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetPort),
		v8.None)
	prototypeTmpl.SetAccessorProperty("pathname",
		v8.NewFunctionTemplateWithError(iso, wrapper.Pathname),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetPathname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("search",
		v8.NewFunctionTemplateWithError(iso, wrapper.Search),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetSearch),
		v8.None)
	prototypeTmpl.SetAccessorProperty("searchParams",
		v8.NewFunctionTemplateWithError(iso, wrapper.SearchParams),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("hash",
		v8.NewFunctionTemplateWithError(iso, wrapper.Hash),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetHash),
		v8.None)

	return constructor
}

func (u urlV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	url, err1 := TryParseArg(args, 0, u.decodeUSVString)
	base, err2 := TryParseArg(args, 1, u.decodeUSVString)
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

func (u urlV8Wrapper) ToJSON(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result, callErr := instance.ToJSON()
	if callErr != nil {
		return nil, callErr
	} else {
		return u.toUSVString(ctx, result)
	}
}

func (u urlV8Wrapper) Href(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) SetHref(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetHref: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) Origin(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) Protocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) SetProtocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetProtocol: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) Username(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.Username: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) SetUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetUsername: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) Password(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.Password: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) SetPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetPassword: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) Host(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) SetHost(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetHost: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) Hostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) SetHostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetHostname: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) Port(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) SetPort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetPort: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) Pathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) SetPathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetPathname: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) Search(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) SetSearch(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetSearch: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) SearchParams(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SearchParams: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}

func (u urlV8Wrapper) Hash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) SetHash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.SetHash: Not implemented. Create an issue: https://github.com/stroiman/go-dom/issues")
}
