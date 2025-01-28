package v8host

import (
	"github.com/gost-dom/browser/browser/dom"
	. "github.com/gost-dom/browser/browser/dom"

	v8 "github.com/tommie/v8go"
)

type documentFragmentV8Wrapper struct {
	esElementContainerWrapper[DocumentFragment]
}

func (w documentFragmentV8Wrapper) constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := w.scriptHost.mustGetContext(info.Context())
	result := dom.NewDocumentFragment(ctx.window.Document())
	_, err := w.store(result, ctx, info.This())
	return nil, err
}

func createDocumentFragmentPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	wrapper := documentFragmentV8Wrapper{newESContainerWrapper[DocumentFragment](host)}
	constructor := v8.NewFunctionTemplateWithError(host.iso, wrapper.constructor)
	constructor.InstanceTemplate().SetInternalFieldCount(1)
	wrapper.Install(constructor)
	return constructor
}
