// This file is generated. Do not edit.

package v8host

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("URL", "", createUrlPrototype)
}

func createUrlPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newUrlV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("toJSON", v8.NewFunctionTemplateWithError(iso, wrapper.toJSON))

	prototypeTmpl.SetAccessorProperty("href",
		v8.NewFunctionTemplateWithError(iso, wrapper.href),
		v8.NewFunctionTemplateWithError(iso, wrapper.setHref),
		v8.None)
	prototypeTmpl.SetAccessorProperty("origin",
		v8.NewFunctionTemplateWithError(iso, wrapper.origin),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("protocol",
		v8.NewFunctionTemplateWithError(iso, wrapper.protocol),
		v8.NewFunctionTemplateWithError(iso, wrapper.setProtocol),
		v8.None)
	prototypeTmpl.SetAccessorProperty("username",
		v8.NewFunctionTemplateWithError(iso, wrapper.username),
		v8.NewFunctionTemplateWithError(iso, wrapper.setUsername),
		v8.None)
	prototypeTmpl.SetAccessorProperty("password",
		v8.NewFunctionTemplateWithError(iso, wrapper.password),
		v8.NewFunctionTemplateWithError(iso, wrapper.setPassword),
		v8.None)
	prototypeTmpl.SetAccessorProperty("host",
		v8.NewFunctionTemplateWithError(iso, wrapper.host),
		v8.NewFunctionTemplateWithError(iso, wrapper.setHost),
		v8.None)
	prototypeTmpl.SetAccessorProperty("hostname",
		v8.NewFunctionTemplateWithError(iso, wrapper.hostname),
		v8.NewFunctionTemplateWithError(iso, wrapper.setHostname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("port",
		v8.NewFunctionTemplateWithError(iso, wrapper.port),
		v8.NewFunctionTemplateWithError(iso, wrapper.setPort),
		v8.None)
	prototypeTmpl.SetAccessorProperty("pathname",
		v8.NewFunctionTemplateWithError(iso, wrapper.pathname),
		v8.NewFunctionTemplateWithError(iso, wrapper.setPathname),
		v8.None)
	prototypeTmpl.SetAccessorProperty("search",
		v8.NewFunctionTemplateWithError(iso, wrapper.search),
		v8.NewFunctionTemplateWithError(iso, wrapper.setSearch),
		v8.None)
	prototypeTmpl.SetAccessorProperty("searchParams",
		v8.NewFunctionTemplateWithError(iso, wrapper.searchParams),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("hash",
		v8.NewFunctionTemplateWithError(iso, wrapper.hash),
		v8.NewFunctionTemplateWithError(iso, wrapper.setHash),
		v8.None)

	return constructor
}

func (u urlV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.scriptHost, info)
	url, err1 := tryParseArg(args, 0, u.decodeUSVString)
	base, err2 := tryParseArg(args, 1, u.decodeUSVString)
	ctx := u.mustGetContext(info)
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
	return nil, errors.New("URL.constructor: Missing arguments")
}

func (u urlV8Wrapper) toJSON(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
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

func (u urlV8Wrapper) href(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Href()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) setHref(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setHref: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) origin(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Origin()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) protocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Protocol()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) setProtocol(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setProtocol: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) username(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.username: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) setUsername(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setUsername: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) password(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.password: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) setPassword(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setPassword: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) host(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Host()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) setHost(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setHost: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) hostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hostname()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) setHostname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setHostname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) port(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Port()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) setPort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setPort: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) pathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Pathname()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) setPathname(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setPathname: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) search(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Search()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) setSearch(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setSearch: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) searchParams(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.searchParams: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (u urlV8Wrapper) hash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.mustGetContext(info)
	instance, err := u.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Hash()
	return u.toUSVString(ctx, result)
}

func (u urlV8Wrapper) setHash(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("URL.setHash: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
