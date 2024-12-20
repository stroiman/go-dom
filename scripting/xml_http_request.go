package scripting

import (
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
