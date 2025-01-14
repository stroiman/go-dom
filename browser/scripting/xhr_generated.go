// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func CreateXmlHttpRequestPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewESXmlHttpRequest(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("open", v8.NewFunctionTemplateWithError(iso, wrapper.Open))
	prototypeTmpl.Set("setRequestHeader", v8.NewFunctionTemplateWithError(iso, wrapper.SetRequestHeader))
	prototypeTmpl.Set("send", v8.NewFunctionTemplateWithError(iso, wrapper.Send))
	prototypeTmpl.Set("abort", v8.NewFunctionTemplateWithError(iso, wrapper.Abort))
	prototypeTmpl.Set("getResponseHeader", v8.NewFunctionTemplateWithError(iso, wrapper.GetResponseHeader))
	prototypeTmpl.Set("getAllResponseHeaders", v8.NewFunctionTemplateWithError(iso, wrapper.GetAllResponseHeaders))
	prototypeTmpl.Set("overrideMimeType", v8.NewFunctionTemplateWithError(iso, wrapper.OverrideMimeType))

	prototypeTmpl.SetAccessorProperty("readyState",
		v8.NewFunctionTemplateWithError(iso, wrapper.ReadyState),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("timeout",
		v8.NewFunctionTemplateWithError(iso, wrapper.Timeout),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetTimeout),
		v8.None)
	prototypeTmpl.SetAccessorProperty("withCredentials",
		v8.NewFunctionTemplateWithError(iso, wrapper.WithCredentials),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetWithCredentials),
		v8.None)
	prototypeTmpl.SetAccessorProperty("upload",
		v8.NewFunctionTemplateWithError(iso, wrapper.Upload),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseURL",
		v8.NewFunctionTemplateWithError(iso, wrapper.ResponseURL),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("status",
		v8.NewFunctionTemplateWithError(iso, wrapper.Status),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("statusText",
		v8.NewFunctionTemplateWithError(iso, wrapper.StatusText),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseType",
		v8.NewFunctionTemplateWithError(iso, wrapper.ResponseType),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetResponseType),
		v8.None)
	prototypeTmpl.SetAccessorProperty("response",
		v8.NewFunctionTemplateWithError(iso, wrapper.Response),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseText",
		v8.NewFunctionTemplateWithError(iso, wrapper.ResponseText),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseXML",
		v8.NewFunctionTemplateWithError(iso, wrapper.ResponseXML),
		nil,
		v8.None)

	return constructor
}

func (xhr ESXmlHttpRequest) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	return xhr.CreateInstance(ctx, info.This())
}

func (xhr ESXmlHttpRequest) SetRequestHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.host, info)
	instance, err0 := xhr.GetInstance(info)
	name, err1 := TryParseArg(args, 0, xhr.DecodeByteString)
	value, err2 := TryParseArg(args, 1, xhr.DecodeByteString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		instance.SetRequestHeader(name, value)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
}

func (xhr ESXmlHttpRequest) Send(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.host, info)
	instance, err0 := xhr.GetInstance(info)
	body, err1 := TryParseArg(args, 0, xhr.DecodeDocument, xhr.DecodeXMLHttpRequestBodyInit)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.SendBody(body)
		return nil, callErr
	}
	if err0 != nil {
		return nil, err0
	}
	callErr := instance.Send()
	return nil, callErr
}

func (xhr ESXmlHttpRequest) Abort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	callErr := instance.Abort()
	return nil, callErr
}

func (xhr ESXmlHttpRequest) GetResponseHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	args := newArgumentHelper(xhr.host, info)
	instance, err0 := xhr.GetInstance(info)
	name, err1 := TryParseArg(args, 0, xhr.DecodeByteString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetResponseHeader(name)
		return xhr.ToNullableByteString(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (xhr ESXmlHttpRequest) GetAllResponseHeaders(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result, callErr := instance.GetAllResponseHeaders()
	if callErr != nil {
		return nil, callErr
	} else {
		return xhr.ToByteString(ctx, result)
	}
}

func (xhr ESXmlHttpRequest) OverrideMimeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.host, info)
	instance, err0 := xhr.GetInstance(info)
	mime, err1 := TryParseArg(args, 0, xhr.DecodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.OverrideMimeType(mime)
		return nil, callErr
	}
	return nil, errors.New("Missing arguments")
}

func (xhr ESXmlHttpRequest) ReadyState(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.ReadyState")
}

func (xhr ESXmlHttpRequest) Timeout(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Timeout()
	return xhr.ToUnsignedLong(ctx, result)
}

func (xhr ESXmlHttpRequest) SetTimeout(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.host, info)
	instance, err0 := xhr.GetInstance(info)
	val, err1 := TryParseArg(args, 0, xhr.DecodeUnsignedLong)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetTimeout(val)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
}

func (xhr ESXmlHttpRequest) WithCredentials(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.WithCredentials()
	return xhr.ToBoolean(ctx, result)
}

func (xhr ESXmlHttpRequest) SetWithCredentials(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.host, info)
	instance, err0 := xhr.GetInstance(info)
	val, err1 := TryParseArg(args, 0, xhr.DecodeBoolean)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetWithCredentials(val)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
}

func (xhr ESXmlHttpRequest) ResponseURL(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.ResponseURL()
	return xhr.ToUSVString(ctx, result)
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

func (xhr ESXmlHttpRequest) ResponseType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.ResponseType")
}

func (xhr ESXmlHttpRequest) SetResponseType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: XMLHttpRequest.SetResponseType")
}

func (xhr ESXmlHttpRequest) Response(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.host.MustGetContext(info.Context())
	instance, err := xhr.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Response()
	return xhr.ToAny(ctx, result)
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
