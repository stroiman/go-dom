package scripting

import (
	"errors"

	v8 "github.com/tommie/v8go"
)

func CreateWindowTemplate(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	windowTemplateFn := v8.NewFunctionTemplateWithError(
		iso,
		func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, v8.NewTypeError(iso, "Illegal Constructor")
		},
	)
	windowTemplate := windowTemplateFn.GetInstanceTemplate()
	windowTemplate.SetInternalFieldCount(1)
	windowTemplate.SetAccessorPropertyCallback(
		"window",
		func(i *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return i.This().Value, nil
		},
		nil, v8.ReadOnly,
	)
	windowTemplate.SetAccessorPropertyCallback(
		"document",
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			if ctx, ok := host.GetContext(info.Context()); ok {
				return ctx.GetInstanceForNode(host.document, ctx.window.Document())
			}
			return nil, errors.New("Must have a context")
		},
		nil, v8.ReadOnly)
	windowTemplate.Set("Document", host.document)
	windowTemplate.Set("Node", host.node)
	windowTemplate.Set("CustomEvent", host.customEvent)
	windowTemplate.Set("EventTarget", host.eventTarget)
	windowTemplate.Set("Window", windowTemplateFn)
	windowTemplate.Set("HTMLElement", host.htmlElement)
	windowTemplate.Set("Location", host.location)
	windowTemplate.Set("location", host.location.GetInstanceTemplate())
	return windowTemplateFn
}
