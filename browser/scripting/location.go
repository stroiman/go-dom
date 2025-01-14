package scripting

import (
	. "github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

func createLocationPrototype(host *ScriptHost) *v8.FunctionTemplate {
	builder := NewIllegalConstructorBuilder[Location](host)
	builder.instanceLookup = func(ctx *ScriptContext, this *v8.Object) (Location, error) {
		location := ctx.Window().Location()
		return location, nil
	}
	helper := builder.NewPrototypeBuilder()
	helper.CreateReadonlyProp("hash", Location.Hash)
	helper.CreateReadonlyProp("host", Location.Host)
	helper.CreateReadonlyProp("hostname", Location.Hostname)
	helper.CreateReadonlyProp("href", Location.Href)
	helper.CreateReadonlyProp("origin", Location.Origin)
	helper.CreateReadonlyProp("pathname", Location.Pathname)
	helper.CreateReadonlyProp("port", Location.Port)
	helper.CreateReadonlyProp("protocol", Location.Protocol)
	helper.CreateReadonlyProp("search", Location.Search)
	return builder.constructor
}
