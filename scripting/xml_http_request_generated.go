// This file is generated. Do not edit.

package scripting

import (
	"errors"
	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

func CreateXmlHttpRequestPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	instance := NewJSXmlHttpRequest(host)
	builder := NewConstructorBuilder[browser.XmlHttpRequest](host, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		scriptContext := host.MustGetContext(info.Context())
		instance := scriptContext.Window().NewXmlHttpRequest()
		return scriptContext.CacheNode(info.This(), instance)
	})
	builder.SetDefaultInstanceLookup()
	protoBuilder := builder.NewPrototypeBuilder()
	prototype := protoBuilder.proto

	prototype.Set("open", v8.NewFunctionTemplateWithError(iso, instance.Open))
	prototype.Set("setRequestHeader", v8.NewFunctionTemplateWithError(iso, instance.SetRequestHeader))
	prototype.Set("send", v8.NewFunctionTemplateWithError(iso, instance.Send))
	prototype.Set("abort", v8.NewFunctionTemplateWithError(iso, instance.Abort))
	prototype.Set("getResponseHeader", v8.NewFunctionTemplateWithError(iso, instance.GetResponseHeader))
	prototype.Set("getAllResponseHeaders", v8.NewFunctionTemplateWithError(iso, instance.GetAllResponseHeaders))
	prototype.Set("overrideMimeType", v8.NewFunctionTemplateWithError(iso, instance.OverrideMimeType))
	return builder.constructor
}

func (xhr JSXmlHttpRequest) Open(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	{
		instance, err := xhr.GetInstance(info)
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
	}
}

func (xhr JSXmlHttpRequest) SetRequestHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	{
		instance, err := xhr.GetInstance(info)
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
	}
}

func (xhr JSXmlHttpRequest) Send(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	{
		ctx := xhr.host.MustGetContext(info.Context())
		instance, err := xhr.GetInstance(info)
		if err != nil {
			return nil, err
		}
		args := info.Args()
		argsLen := len(args)
		if argsLen >= 1 {
			body, err := TryParseArgs(ctx, args, 0, GetBodyFromDocument, GetBodyFromXMLHttpRequestBodyInit)
			if err != nil {
				return nil, err
			}
			err = instance.SendBody(body)
			return nil, err
		}

		err = instance.Send()
		return nil, err
	}
}

func (xhr JSXmlHttpRequest) Abort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	{
		instance, err := xhr.GetInstance(info)
		if err != nil {
			return nil, err
		}

		err = instance.Abort()
		return nil, err
	}
}

func (xhr JSXmlHttpRequest) GetResponseHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	{
		ctx := xhr.host.MustGetContext(info.Context())
		instance, err := xhr.GetInstance(info)
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
		result := instance.GetResponseHeader(name)
		return ToNullableByteString(ctx, result)
	}
}

func (xhr JSXmlHttpRequest) GetAllResponseHeaders(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	{
		ctx := xhr.host.MustGetContext(info.Context())
		instance, err := xhr.GetInstance(info)
		if err != nil {
			return nil, err
		}

		result, err := instance.GetAllResponseHeaders()
		if err != nil {
			return nil, err
		} else {
			return ToByteString(ctx, result)
		}
	}
}

func (xhr JSXmlHttpRequest) OverrideMimeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	{
		instance, err := xhr.GetInstance(info)
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
	}
}
