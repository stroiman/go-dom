// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func CreateXmlHttpRequestPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewESXmlHttpRequest(host)
	constructor := v8.NewFunctionTemplateWithError(iso, func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
		ctx := host.MustGetContext(info.Context())
		instance := wrapper.CreateInstance(ctx)
		_, err := ctx.CacheNode(info.This(), instance)
		return nil, err
	})
	constructor.GetInstanceTemplate().SetInternalFieldCount(1)
	prototype := constructor.PrototypeTemplate()

	prototype.Set("open", v8.NewFunctionTemplateWithError(iso, wrapper.Open))
	prototype.Set("setRequestHeader", v8.NewFunctionTemplateWithError(iso, wrapper.SetRequestHeader))
	prototype.Set("send", v8.NewFunctionTemplateWithError(iso, wrapper.Send))
	prototype.Set("abort", v8.NewFunctionTemplateWithError(iso, wrapper.Abort))
	prototype.Set("getResponseHeader", v8.NewFunctionTemplateWithError(iso, wrapper.GetResponseHeader))
	prototype.Set("getAllResponseHeaders", v8.NewFunctionTemplateWithError(iso, wrapper.GetAllResponseHeaders))
	prototype.Set("overrideMimeType", v8.NewFunctionTemplateWithError(iso, wrapper.OverrideMimeType))
	return constructor
}

func (xhr ESXmlHttpRequest) Open(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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
		method, err0 := xhr.GetArgByteString(args, 0)
		url, err1 := xhr.GetArgUSVString(args, 1)
		err = errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.Open(method, url)
		return nil, nil
	}
}

func (xhr ESXmlHttpRequest) SetRequestHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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
		name, err0 := xhr.GetArgByteString(args, 0)
		value, err1 := xhr.GetArgByteString(args, 1)
		err = errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetRequestHeader(name, value)
		return nil, nil
	}
}

func (xhr ESXmlHttpRequest) Send(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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

func (xhr ESXmlHttpRequest) Abort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	{
		instance, err := xhr.GetInstance(info)
		if err != nil {
			return nil, err
		}

		err = instance.Abort()
		return nil, err
	}
}

func (xhr ESXmlHttpRequest) GetResponseHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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
		name, err := xhr.GetArgByteString(args, 0)
		if err != nil {
			return nil, err
		}
		result := instance.GetResponseHeader(name)
		return ToNullableByteString(ctx, result)
	}
}

func (xhr ESXmlHttpRequest) GetAllResponseHeaders(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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

func (xhr ESXmlHttpRequest) OverrideMimeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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
		mime, err := xhr.GetArgDOMString(args, 0)
		if err != nil {
			return nil, err
		}
		err = instance.OverrideMimeType(mime)
		return nil, err
	}
}
