package scripting

// This file is generated. Do not edit.

import (
	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

func CreateXmlHttpRequestPrototype(host *ScriptHost) *v8.FunctionTemplate {
	// iso := host.iso
	builder := NewConstructorBuilder[browser.XmlHttpRequest](
		host,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			scriptContext := host.MustGetContext(info.Context())
			instance := scriptContext.Window().NewXmlHttpRequest()
			return scriptContext.CacheNode(info.This(), instance)
		},
	)
	return builder.constructor
}