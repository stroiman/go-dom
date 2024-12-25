package scripting

import (
	"errors"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type ESXmlHttpRequest struct{ ESWrapper[XmlHttpRequest] }

func NewESXmlHttpRequest(host *ScriptHost) ESXmlHttpRequest {
	return ESXmlHttpRequest{NewESWrapper[XmlHttpRequest](host)}
}

func (w ESXmlHttpRequest) CreateInstance(ctx *ScriptContext, this *v8.Object) (*v8.Value, error) {
	result := ctx.Window().NewXmlHttpRequest()
	ctx.CacheNode(this, result)
	return nil, nil
}

func (xhr ESXmlHttpRequest) Open(info *v8.FunctionCallbackInfo) (result *v8.Value, err error) {
	args := newArgumentHelper(xhr.host, info)
	method, err0 := TryParseArg(args, 0, xhr.DecodeUSVString)
	url, err1 := TryParseArg(args, 1, xhr.DecodeUSVString)
	async, err2 := TryParseArg(args, 1, xhr.DecodeBoolean)
	// username, err3 := TryParseArg(args, 1, u.DecodeUSVString)
	// password, err4 := TryParseArg(args, 1, u.DecodeUSVString)
	instance, errInstance := xhr.GetInstance(info)
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
