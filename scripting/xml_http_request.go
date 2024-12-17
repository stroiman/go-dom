package scripting

import (
	. "github.com/stroiman/go-dom/browser"
)

type ESXmlHttpRequest struct{ ESWrapper[XmlHttpRequest] }

func NewESXmlHttpRequest(host *ScriptHost) ESXmlHttpRequest {
	return ESXmlHttpRequest{NewESWrapper[XmlHttpRequest](host)}
}

func (w ESXmlHttpRequest) CreateInstance(ctx *ScriptContext) XmlHttpRequest {
	return ctx.Window().NewXmlHttpRequest()
}
