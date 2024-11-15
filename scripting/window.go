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
	windowTemplate.SetAccessorProperty(
		"window",
		v8.AccessProp{
			Get: func(i *v8.FunctionCallbackInfo) *v8.Value {
				return i.This().Value
			},
			Attributes: v8.ReadOnly,
		},
	)
	windowTemplate.SetAccessorPropertyWithError(
		"document",
		v8.AccessPropWithError{
			Get: func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				if ctx, ok := host.GetContext(info.Context()); ok {
					return ctx.GetInstanceForNode(host.document, ctx.window.Document())
				}
				return nil, errors.New("Must have a context")
			},
		})
	windowTemplate.Set("Document", host.document)
	windowTemplate.Set("Node", host.node)
	windowTemplate.Set("EventTarget", host.eventTarget)
	windowTemplate.Set("Window", windowTemplateFn)
	windowTemplate.Set("HTMLElement", host.htmlElement)
	return windowTemplateFn
}
