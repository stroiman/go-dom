// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func CreateXmlHttpRequestPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewESXmlHttpRequest(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.NewInstance)
	constructor.GetInstanceTemplate().SetInternalFieldCount(1)
	prototype := constructor.PrototypeTemplate()

	prototype.Set("open", v8.NewFunctionTemplateWithError(iso, wrapper.Open))
	prototype.Set("setRequestHeader", v8.NewFunctionTemplateWithError(iso, wrapper.SetRequestHeader))
	prototype.Set("send", v8.NewFunctionTemplateWithError(iso, wrapper.Send))
	prototype.Set("abort", v8.NewFunctionTemplateWithError(iso, wrapper.Abort))
	prototype.Set("getResponseHeader", v8.NewFunctionTemplateWithError(iso, wrapper.GetResponseHeader))
	prototype.Set("getAllResponseHeaders", v8.NewFunctionTemplateWithError(iso, wrapper.GetAllResponseHeaders))
	prototype.Set("overrideMimeType", v8.NewFunctionTemplateWithError(iso, wrapper.OverrideMimeType))

	prototype.SetAccessorProperty("readyState",
		v8.NewFunctionTemplateWithError(iso, wrapper.ReadyState),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("timeout",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetTimeout),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetTimeout),
		v8.None)
	prototype.SetAccessorProperty("withCredentials",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetWithCredentials),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetWithCredentials),
		v8.None)
	prototype.SetAccessorProperty("upload",
		v8.NewFunctionTemplateWithError(iso, wrapper.Upload),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("responseURL",
		v8.NewFunctionTemplateWithError(iso, wrapper.ResponseURL),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("status",
		v8.NewFunctionTemplateWithError(iso, wrapper.Status),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("statusText",
		v8.NewFunctionTemplateWithError(iso, wrapper.StatusText),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("responseType",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetResponseType),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetResponseType),
		v8.None)
	prototype.SetAccessorProperty("response",
		v8.NewFunctionTemplateWithError(iso, wrapper.Response),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("responseText",
		v8.NewFunctionTemplateWithError(iso, wrapper.ResponseText),
		nil,
		v8.ReadOnly)
	prototype.SetAccessorProperty("responseXML",
		v8.NewFunctionTemplateWithError(iso, wrapper.ResponseXML),
		nil,
		v8.ReadOnly)
	return constructor
}

func (xhr ESXmlHttpRequest) NewInstance(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	return xhr.CreateInstance(ctx, info.This())
}

func (xhr ESXmlHttpRequest) SetRequestHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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

func (xhr ESXmlHttpRequest) Send(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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

func (xhr ESXmlHttpRequest) Abort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	err = instance.Abort()
	return nil, err
}

func (xhr ESXmlHttpRequest) GetResponseHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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
	return xhr.ToNullableByteString(ctx, result)
}

func (xhr ESXmlHttpRequest) GetAllResponseHeaders(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result, err := instance.GetAllResponseHeaders()
	if err != nil {
		return nil, err
	} else {
		return xhr.ToByteString(ctx, result)
	}
}

func (xhr ESXmlHttpRequest) OverrideMimeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
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

func (xhr ESXmlHttpRequest) ReadyState(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.ReadyState")
}

func (xhr ESXmlHttpRequest) GetTimeout(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetTimeout()
	return xhr.ToUnsignedLong(ctx, result)
}

func (xhr ESXmlHttpRequest) SetTimeout(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := info.Args()
	argsLen := len(args)
	if argsLen < 1 {
		return nil, errors.New("Too few arguments")
	}
	val, err := xhr.GetArgUnsignedLong(args, 0)
	if err != nil {
		return nil, err
	}
	instance.SetTimeout(val)
	return nil, nil
}

func (xhr ESXmlHttpRequest) GetWithCredentials(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetWithCredentials()
	return xhr.ToBoolean(ctx, result)
}

func (xhr ESXmlHttpRequest) SetWithCredentials(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	args := info.Args()
	argsLen := len(args)
	if argsLen < 1 {
		return nil, errors.New("Too few arguments")
	}
	val, err := xhr.GetArgBoolean(args, 0)
	if err != nil {
		return nil, err
	}
	instance.SetWithCredentials(val)
	return nil, nil
}

func (xhr ESXmlHttpRequest) Upload(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.Upload")
}

func (xhr ESXmlHttpRequest) ResponseURL(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.ResponseURL")
}

func (xhr ESXmlHttpRequest) Status(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Status()
	return xhr.ToUnsignedShort(ctx, result)
}

func (xhr ESXmlHttpRequest) StatusText(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.StatusText()
	return xhr.ToByteString(ctx, result)
}

func (xhr ESXmlHttpRequest) GetResponseType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.GetResponseType")
}

func (xhr ESXmlHttpRequest) SetResponseType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.SetResponseType")
}

func (xhr ESXmlHttpRequest) Response(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.Response")
}

func (xhr ESXmlHttpRequest) ResponseText(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.ResponseText()
	return xhr.ToUSVString(ctx, result)
}

func (xhr ESXmlHttpRequest) ResponseXML(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.ResponseXML")
}
