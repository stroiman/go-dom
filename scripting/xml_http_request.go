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
		args := info.Args()
		method, err0 := GetArgByteString(args, 0)
		url, err1 := GetArgUSVString(args, 1)
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		instance.Open(method, url)
		return nil, nil
	}))

	prototype.Set("setRequestHeader", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		args := info.Args()
		name, err0 := GetArgByteString(args, 0)
		value, err1 := GetArgByteString(args, 1)
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		instance.SetRequestHeader(name, value)
		return nil, nil
	}))

	prototype.Set("send", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		args := info.Args()
		body, err := GetArg(args, 0)
		if err != nil {
			return nil, err
		}
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		instance.Send(body)
		return nil, nil
	}))

	prototype.Set("abort", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		args := info.Args()
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		instance.Abort()
		return nil, nil
	}))

	prototype.Set("getResponseHeader", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		args := info.Args()
		name, err := GetArgByteString(args, 0)
		if err != nil {
			return nil, err
		}
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		instance.GetResponseHeader(name)
		return nil, nil
	}))

	prototype.Set("getAllResponseHeaders", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		args := info.Args()
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		instance.GetAllResponseHeaders()
		return nil, nil
	}))

	prototype.Set("overrideMimeType", v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		args := info.Args()
		mime, err := GetArgDOMString(args, 0)
		if err != nil {
			return nil, err
		}
		instance, err := builder.GetInstance(info)
		if err != nil {
			return nil, err
		}
		instance.OverrideMimeType(mime)
		return nil, nil
	}))

	builder.SetDefaultInstanceLookup()
	return builder.constructor
}
