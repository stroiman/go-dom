// This file is generated. Do not edit.

package v8host

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func createXmlHttpRequestPrototype(scriptHost *V8ScriptHost) *v8.FunctionTemplate {
	iso := scriptHost.iso
	wrapper := newXmlHttpRequestV8Wrapper(scriptHost)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("open", v8.NewFunctionTemplateWithError(iso, wrapper.open))
	prototypeTmpl.Set("setRequestHeader", v8.NewFunctionTemplateWithError(iso, wrapper.setRequestHeader))
	prototypeTmpl.Set("send", v8.NewFunctionTemplateWithError(iso, wrapper.send))
	prototypeTmpl.Set("abort", v8.NewFunctionTemplateWithError(iso, wrapper.abort))
	prototypeTmpl.Set("getResponseHeader", v8.NewFunctionTemplateWithError(iso, wrapper.getResponseHeader))
	prototypeTmpl.Set("getAllResponseHeaders", v8.NewFunctionTemplateWithError(iso, wrapper.getAllResponseHeaders))
	prototypeTmpl.Set("overrideMimeType", v8.NewFunctionTemplateWithError(iso, wrapper.overrideMimeType))

	prototypeTmpl.SetAccessorProperty("readyState",
		v8.NewFunctionTemplateWithError(iso, wrapper.readyState),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("timeout",
		v8.NewFunctionTemplateWithError(iso, wrapper.timeout),
		v8.NewFunctionTemplateWithError(iso, wrapper.setTimeout),
		v8.None)
	prototypeTmpl.SetAccessorProperty("withCredentials",
		v8.NewFunctionTemplateWithError(iso, wrapper.withCredentials),
		v8.NewFunctionTemplateWithError(iso, wrapper.setWithCredentials),
		v8.None)
	prototypeTmpl.SetAccessorProperty("upload",
		v8.NewFunctionTemplateWithError(iso, wrapper.upload),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseURL",
		v8.NewFunctionTemplateWithError(iso, wrapper.responseURL),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("status",
		v8.NewFunctionTemplateWithError(iso, wrapper.status),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("statusText",
		v8.NewFunctionTemplateWithError(iso, wrapper.statusText),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseType",
		v8.NewFunctionTemplateWithError(iso, wrapper.responseType),
		v8.NewFunctionTemplateWithError(iso, wrapper.setResponseType),
		v8.None)
	prototypeTmpl.SetAccessorProperty("response",
		v8.NewFunctionTemplateWithError(iso, wrapper.response),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseText",
		v8.NewFunctionTemplateWithError(iso, wrapper.responseText),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("responseXML",
		v8.NewFunctionTemplateWithError(iso, wrapper.responseXML),
		nil,
		v8.None)

	return constructor
}

func (xhr xmlHttpRequestV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	return xhr.CreateInstance(ctx, info.This())
}

func (xhr xmlHttpRequestV8Wrapper) setRequestHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.scriptHost, info)
	instance, err0 := xhr.getInstance(info)
	name, err1 := tryParseArg(args, 0, xhr.decodeByteString)
	value, err2 := tryParseArg(args, 1, xhr.decodeByteString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		instance.SetRequestHeader(name, value)
		return nil, nil
	}
	return nil, errors.New("XMLHttpRequest.setRequestHeader: Missing arguments")
}

func (xhr xmlHttpRequestV8Wrapper) send(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.scriptHost, info)
	instance, err0 := xhr.getInstance(info)
	body, err1 := tryParseArg(args, 0, xhr.decodeDocument, xhr.decodeXMLHttpRequestBodyInit)
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

func (xhr xmlHttpRequestV8Wrapper) abort(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	instance, err := xhr.getInstance(info)
	if err != nil {
		return nil, err
	}
	callErr := instance.Abort()
	return nil, callErr
}

func (xhr xmlHttpRequestV8Wrapper) getResponseHeader(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	args := newArgumentHelper(xhr.scriptHost, info)
	instance, err0 := xhr.getInstance(info)
	name, err1 := tryParseArg(args, 0, xhr.decodeByteString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetResponseHeader(name)
		return xhr.toNullableByteString(ctx, result)
	}
	return nil, errors.New("XMLHttpRequest.getResponseHeader: Missing arguments")
}

func (xhr xmlHttpRequestV8Wrapper) getAllResponseHeaders(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	instance, err := xhr.getInstance(info)
	if err != nil {
		return nil, err
	}
	result, callErr := instance.GetAllResponseHeaders()
	if callErr != nil {
		return nil, callErr
	} else {
		return xhr.toByteString(ctx, result)
	}
}

func (xhr xmlHttpRequestV8Wrapper) overrideMimeType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.scriptHost, info)
	instance, err0 := xhr.getInstance(info)
	mime, err1 := tryParseArg(args, 0, xhr.decodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.OverrideMimeType(mime)
		return nil, callErr
	}
	return nil, errors.New("XMLHttpRequest.overrideMimeType: Missing arguments")
}

func (xhr xmlHttpRequestV8Wrapper) readyState(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("XMLHttpRequest.readyState: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (xhr xmlHttpRequestV8Wrapper) timeout(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	instance, err := xhr.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Timeout()
	return xhr.toUnsignedLong(ctx, result)
}

func (xhr xmlHttpRequestV8Wrapper) setTimeout(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.scriptHost, info)
	instance, err0 := xhr.getInstance(info)
	val, err1 := tryParseArg(args, 0, xhr.decodeUnsignedLong)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetTimeout(val)
		return nil, nil
	}
	return nil, errors.New("XMLHttpRequest.setTimeout: Missing arguments")
}

func (xhr xmlHttpRequestV8Wrapper) withCredentials(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	instance, err := xhr.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.WithCredentials()
	return xhr.toBoolean(ctx, result)
}

func (xhr xmlHttpRequestV8Wrapper) setWithCredentials(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(xhr.scriptHost, info)
	instance, err0 := xhr.getInstance(info)
	val, err1 := tryParseArg(args, 0, xhr.decodeBoolean)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetWithCredentials(val)
		return nil, nil
	}
	return nil, errors.New("XMLHttpRequest.setWithCredentials: Missing arguments")
}

func (xhr xmlHttpRequestV8Wrapper) responseURL(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	instance, err := xhr.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.ResponseURL()
	return xhr.toUSVString(ctx, result)
}

func (xhr xmlHttpRequestV8Wrapper) status(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	instance, err := xhr.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Status()
	return xhr.toUnsignedShort(ctx, result)
}

func (xhr xmlHttpRequestV8Wrapper) statusText(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	instance, err := xhr.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.StatusText()
	return xhr.toByteString(ctx, result)
}

func (xhr xmlHttpRequestV8Wrapper) responseType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("XMLHttpRequest.responseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (xhr xmlHttpRequestV8Wrapper) setResponseType(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("XMLHttpRequest.setResponseType: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}

func (xhr xmlHttpRequestV8Wrapper) response(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	instance, err := xhr.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Response()
	return xhr.toAny(ctx, result)
}

func (xhr xmlHttpRequestV8Wrapper) responseText(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := xhr.mustGetContext(info)
	instance, err := xhr.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.ResponseText()
	return xhr.toUSVString(ctx, result)
}

func (xhr xmlHttpRequestV8Wrapper) responseXML(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("XMLHttpRequest.responseXML: Not implemented. Create an issue: https://github.com/gost-dom/browser/issues")
}
