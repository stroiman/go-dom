package scripting

import (
	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

func CreateLocationPrototype(host *ScriptHost) *v8.FunctionTemplate {
	builder := NewIllegalConstructorBuilder[Location](host)
	builder.instanceLookup = func(ctx *ScriptContext, this *v8.Object) (Location, error) {
		location := ctx.Window().Location()
		return location, nil
	}
	helper := builder.NewPrototypeBuilder()
	helper.CreateReadonlyProp("hash", Location.GetHash)
	helper.CreateReadonlyProp("host", Location.GetHost)
	helper.CreateReadonlyProp("hostname", Location.GetHostname)
	helper.CreateReadonlyProp("href", Location.GetHref)
	helper.CreateReadonlyProp("origin", Location.GetOrigin)
	helper.CreateReadonlyProp("pathname", Location.GetPathname)
	helper.CreateReadonlyProp("port", Location.GetPort)
	helper.CreateReadonlyProp("protocol", Location.GetProtocol)
	helper.CreateReadonlyProp("search", Location.GetSearch)
	return builder.constructor
}
