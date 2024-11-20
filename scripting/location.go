package scripting

import (
	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type LocationGetter func(Location) string

func CreateLocationPrototype(host *ScriptHost) *v8.FunctionTemplate {
	CreateGetter := func(fn func(Location) string) v8.FunctionCallbackWithError {
		return func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(arg.Context())
			location := ctx.Window().Location()
			value := fn(location)
			return v8.NewValue(host.iso, value)
		}
	}
	iso := host.iso
	templateFn := v8.NewFunctionTemplateWithError(
		iso,
		func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, v8.NewTypeError(iso, "Illegal Constructor")
		},
	)
	template := templateFn.GetInstanceTemplate()
	template.SetInternalFieldCount(1)
	proto := templateFn.PrototypeTemplate()
	proto.SetAccessorPropertyCallback("hash", CreateGetter(Location.GetHash), nil, 0)
	proto.SetAccessorPropertyCallback("host", CreateGetter(Location.GetHost), nil, 0)
	proto.SetAccessorPropertyCallback("hostname", CreateGetter(Location.GetHostname), nil, 0)
	proto.SetAccessorPropertyCallback("href", CreateGetter(Location.GetHref), nil, 0)
	proto.SetAccessorPropertyCallback("origin", CreateGetter(Location.GetOrigin), nil, 0)
	proto.SetAccessorPropertyCallback("pathname", CreateGetter(Location.GetPathname), nil, 0)
	proto.SetAccessorPropertyCallback("port", CreateGetter(Location.GetPort), nil, 0)
	proto.SetAccessorPropertyCallback("protocol", CreateGetter(Location.GetProtocol), nil, 0)
	proto.SetAccessorPropertyCallback("search", CreateGetter(Location.GetSearch), nil, 0)
	return templateFn
}
