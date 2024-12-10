// This file is generated. Do not edit.

package scripting

import (
	"errors"
	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

func CreateXmlHttpRequestPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	builder := NewConstructorBuilder[browser.XmlHttpRequest](host, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		scriptContext := host.MustGetContext(info.Context())
		instance := scriptContext.Window().NewXmlHttpRequest()
		return scriptContext.CacheNode(info.This(), instance)
	})
	protoBuilder := builder.NewPrototypeBuilder()
	prototype := protoBuilder.proto

	prototype.Set("open", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		args := info.Args()
		argsLen := len(args)
		if argsLen < 2 {
			return nil, errors.New("Too few arguments")
		}
		method, err0 := GetArgByteString(args, 0)
		url, err1 := GetArgUSVString(args, 1)
		err = errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.Open(method, url)
		return nil, nil
		// Opt: []
	}))

	prototype.Set("setRequestHeader", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		args := info.Args()
		argsLen := len(args)
		if argsLen < 2 {
			return nil, errors.New("Too few arguments")
		}
		name, err0 := GetArgByteString(args, 0)
		value, err1 := GetArgByteString(args, 1)
		err = errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetRequestHeader(name, value)
		return nil, nil
		// Opt: []
	}))

	prototype.Set("send", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		args := info.Args()
		argsLen := len(args)
		if argsLen >= 1 {
			body, err := GetArg(args, 0)
			if err != nil {
				return nil, err
			}
			err = instance.SendBody(body)
			return nil, err
		}

		err = instance.Send()
		return nil, err
		// Opt: [{body  true false}]
	}))

	prototype.Set("abort", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}

		err = instance.Abort()
		return nil, err
		// Opt: []
	}))

	prototype.Set("getResponseHeader", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		args := info.Args()
		argsLen := len(args)
		if argsLen < 1 {
			return nil, errors.New("Too few arguments")
		}
		name, err := GetArgByteString(args, 0)
		if err != nil {
			return nil, err
		}
		ctx := host.MustGetContext(info.Context())
		result := instance.GetResponseHeader(name)
		return ToNullableByteString(ctx, result)
		// Opt: []
	}))

	prototype.Set("getAllResponseHeaders", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}

		ctx := host.MustGetContext(info.Context())
		result, err := instance.GetAllResponseHeaders()
		if err != nil {
			return nil, err
		} else {
			return ToByteString(ctx, result)
		}
		// Opt: []
	}))

	prototype.Set("overrideMimeType", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		args := info.Args()
		argsLen := len(args)
		if argsLen < 1 {
			return nil, errors.New("Too few arguments")
		}
		mime, err := GetArgDOMString(args, 0)
		if err != nil {
			return nil, err
		}
		err = instance.OverrideMimeType(mime)
		return nil, err
		// Opt: []
	}))

	builder.SetDefaultInstanceLookup()
	return builder.constructor
}
