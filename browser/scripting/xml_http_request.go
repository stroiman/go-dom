package scripting

import (
	"errors"

	. "github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"

	v8 "github.com/tommie/v8go"
)

type xmlHttpRequestV8Wrapper struct {
	nodeV8WrapperBase[html.XmlHttpRequest]
}

func (xhr xmlHttpRequestV8Wrapper) decodeDocument(
	ctx *ScriptContext,
	val *v8.Value,
) (*html.XHRRequestBody, error) {
	if val.IsNull() {
		return nil, nil
	}
	return nil, errors.New("Not supported yet")
}

func (xhr xmlHttpRequestV8Wrapper) decodeXMLHttpRequestBodyInit(
	ctx *ScriptContext,
	val *v8.Value,
) (*html.XHRRequestBody, error) {
	if val.IsString() {
		return html.NewXHRRequestBodyOfString(val.String()), nil
	}
	if !val.IsObject() {
		return nil, errors.New("Not supported yet")
	}
	obj := val.Object()
	formData := getWrappedInstance[*html.FormData](obj)
	// if ok {
	return html.NewXHRRequestBodyOfFormData(formData), nil
	// }
	// return nil, errors.New("Not a node")
}

func newXmlHttpRequestV8Wrapper(host *ScriptHost) xmlHttpRequestV8Wrapper {
	return xmlHttpRequestV8Wrapper{newNodeV8WrapperBase[html.XmlHttpRequest](host)}
}

func (xhr xmlHttpRequestV8Wrapper) CreateInstance(
	ctx *ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	result := ctx.Window().NewXmlHttpRequest()
	result.SetCatchAllHandler(NewEventHandlerFunc(func(event Event) error {
		prop := "on" + event.Type()
		handler, err := this.Get(prop)
		if err == nil && handler.IsFunction() {
			v8Event, err := ctx.getInstanceForNode(event)
			if err == nil {
				f, _ := handler.AsFunction()
				f.Call(this, v8Event)
			}
		}
		return nil
	}))
	ctx.cacheNode(this, result)
	return nil, nil
}

func (xhr xmlHttpRequestV8Wrapper) Open(
	info *v8.FunctionCallbackInfo,
) (result *v8.Value, err error) {
	args := newArgumentHelper(xhr.host, info)
	method, err0 := TryParseArg(args, 0, xhr.decodeUSVString)
	url, err1 := TryParseArg(args, 1, xhr.decodeUSVString)
	async, err2 := TryParseArg(args, 1, xhr.decodeBoolean)
	// username, err3 := TryParseArg(args, 1, u.DecodeUSVString)
	// password, err4 := TryParseArg(args, 1, u.DecodeUSVString)
	instance, errInstance := xhr.getInstance(info)
	if args.noOfReadArguments > 2 {
		if err = errors.Join(err0, err1, err2, errInstance); err != nil {
			return
		}
		instance.Open(method, url, html.RequestOptionAsync(async))
		return
	}
	if args.noOfReadArguments < 2 {
		return nil, errors.New("Not enough arguments")
	}
	if err = errors.Join(err0, err1, errInstance); err == nil {
		instance.Open(method, url)
	}
	return
}

func (xhr xmlHttpRequestV8Wrapper) Upload(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return info.This().Value, nil
}
