package v8host

import (
	"errors"
	"io"
	"strings"

	. "github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/internal/html"

	v8 "github.com/tommie/v8go"
)

type xmlHttpRequestV8Wrapper struct {
	handleReffedObject[XmlHttpRequest]
}

func (xhr xmlHttpRequestV8Wrapper) decodeDocument(
	ctx *V8ScriptContext,
	val *v8.Value,
) (io.Reader, error) {
	if val.IsNull() {
		return nil, nil
	}
	return nil, errors.New("Not supported yet")
}

func (xhr xmlHttpRequestV8Wrapper) decodeXMLHttpRequestBodyInit(
	ctx *V8ScriptContext,
	val *v8.Value,
) (io.Reader, error) {
	if val.IsString() {
		return strings.NewReader(val.String()), nil
	}
	if !val.IsObject() {
		return nil, errors.New("Not supported yet")
	}
	obj := val.Object()
	if res, err := getWrappedInstance[*html.FormData](obj); err == nil {
		return res.GetReader(), nil
	} else {
		return nil, err
	}
}

func newXmlHttpRequestV8Wrapper(host *V8ScriptHost) xmlHttpRequestV8Wrapper {
	return xmlHttpRequestV8Wrapper{newHandleReffedObject[XmlHttpRequest](host)}
}

func (xhr xmlHttpRequestV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	result := NewXmlHttpRequest(ctx.window.HTTPClient(), ctx.window.Location().Href())
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
	xhr.store(result, ctx, this)
	return nil, nil
}

func (xhr xmlHttpRequestV8Wrapper) open(
	info *v8.FunctionCallbackInfo,
) (result *v8.Value, err error) {
	args := newArgumentHelper(xhr.scriptHost, info)
	method, err0 := tryParseArg(args, 0, xhr.decodeUSVString)
	url, err1 := tryParseArg(args, 1, xhr.decodeUSVString)
	async, err2 := tryParseArg(args, 1, xhr.decodeBoolean)
	// username, err3 := TryParseArg(args, 1, u.DecodeUSVString)
	// password, err4 := TryParseArg(args, 1, u.DecodeUSVString)
	instance, errInstance := xhr.getInstance(info)
	if args.noOfReadArguments > 2 {
		if err = errors.Join(err0, err1, err2, errInstance); err != nil {
			return
		}
		instance.Open(method, url, RequestOptionAsync(async))
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

func (xhr xmlHttpRequestV8Wrapper) upload(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return info.This().Value, nil
}
