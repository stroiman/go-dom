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
				return ctx.GetInstanceForNodeByName("Document", ctx.window.Document())
			}
			return nil, errors.New("Must have a context")
		},
		nil, v8.ReadOnly)
	return windowTemplateFn
}

func installGlobals(
	windowFnTemplate *v8.FunctionTemplate,
	host *ScriptHost,
	globalInstalls []globalInstall,
) {
	windowTemplate := windowFnTemplate.GetInstanceTemplate()
	for _, globalInstall := range globalInstalls {
		windowTemplate.Set(globalInstall.name, globalInstall.constructor)
	}
	location := host.globals.namedGlobals["Location"]
	windowTemplate.Set("location", location.GetInstanceTemplate())
}
