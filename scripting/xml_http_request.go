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
		if argsLen >= 2 {
			method, err0 := GetArgByteString(args, 0)
			url, err1 := GetArgUSVString(args, 1)
			err := errors.Join(err0, err1)
			if err != nil {
				return nil, err
			}
			instance.Open(method, url)
			return nil, nil
		}
		return nil, errors.New("Missing arguments")
	}))

	prototype.Set("setRequestHeader", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		args := info.Args()
		argsLen := len(args)
		if argsLen >= 2 {
			name, err0 := GetArgByteString(args, 0)
			value, err1 := GetArgByteString(args, 1)
			err := errors.Join(err0, err1)
			if err != nil {
				return nil, err
			}
			instance.SetRequestHeader(name, value)
			return nil, nil
		}
		return nil, errors.New("Missing arguments")
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
			err = instance.Send(body)
			return nil, err
		}
		return nil, errors.New("Missing arguments")
	}))

	prototype.Set("abort", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		err = instance.Abort()
		return nil, err
	}))

	prototype.Set("getResponseHeader", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		args := info.Args()
		argsLen := len(args)
		ctx := host.MustGetContext(info.Context())
		if argsLen >= 1 {
			name, err := GetArgByteString(args, 0)
			if err != nil {
				return nil, err
			}
			result, err := instance.GetResponseHeader(name)
			if err != nil {
				return nil, err
			} else {
				return ToByteString(ctx, result)
			}
		}
		return nil, errors.New("Missing arguments")
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
	}))

	prototype.Set("overrideMimeType", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		args := info.Args()
		argsLen := len(args)
		if argsLen >= 1 {
			mime, err := GetArgDOMString(args, 0)
			if err != nil {
				return nil, err
			}
			err = instance.OverrideMimeType(mime)
			return nil, err
		}
		return nil, errors.New("Missing arguments")
	}))

	builder.SetDefaultInstanceLookup()
	return builder.constructor
}
