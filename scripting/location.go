package scripting

import (
	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

func CreateLocationPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	templateFn := v8.NewFunctionTemplateWithError(
		iso,
		func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, v8.NewTypeError(iso, "Illegal Constructor")
		},
	)
	template := templateFn.GetInstanceTemplate()
	template.SetInternalFieldCount(1)
	helper := PrototypeBuilder[Location]{
		host,
		templateFn.PrototypeTemplate(),
		func(ctx *ScriptContext) (Location, error) {
			location := ctx.Window().Location()
			return location, nil
		},
	}
	helper.CreateReadonlyProp("hash", Location.GetHash)
	helper.CreateReadonlyProp("host", Location.GetHost)
	helper.CreateReadonlyProp("hostname", Location.GetHostname)
	helper.CreateReadonlyProp("href", Location.GetHref)
	helper.CreateReadonlyProp("origin", Location.GetOrigin)
	helper.CreateReadonlyProp("pathname", Location.GetPathname)
	helper.CreateReadonlyProp("port", Location.GetPort)
	helper.CreateReadonlyProp("protocol", Location.GetProtocol)
	helper.CreateReadonlyProp("search", Location.GetSearch)
	return templateFn
}
