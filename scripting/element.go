package scripting

import (
	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

func CreateElement(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	res := v8.NewFunctionTemplateWithError(
		iso,
		func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, v8.NewTypeError(iso, "Illegal Constructor")
		},
	)
	proto := res.PrototypeTemplate()
	proto.SetAccessorPropertyCallback("outerHTML",
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			this, ok := ctx.GetCachedNode(info.This())
			if e, e_ok := this.(Element); e_ok && ok {
				return v8.NewValue(iso, e.OuterHTML())
			} else {
				return nil, v8.NewTypeError(iso, "Not an instance of Element")
			}
		}, nil, v8.ReadOnly)
	instanceTemplate := res.GetInstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	return res
}
