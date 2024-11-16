package scripting

import (
	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

func CreateDocumentPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	res := v8.NewFunctionTemplateWithError(
		iso,
		func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
			scriptContext := host.MustGetContext(args.Context())
			return scriptContext.CacheNode(args.This(), NewDocument())
		},
	)
	instanceTemplate := res.GetInstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	proto := res.PrototypeTemplate()
	proto.Set("createElement", v8.NewFunctionTemplate(iso,
		func(info *v8.FunctionCallbackInfo) *v8.Value {
			return v8.Undefined(iso)
		}))

	proto.SetAccessorPropertyCallback("documentElement",
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(arg.Context())
			this, ok := ctx.domNodes[arg.This().GetInternalField(0).Int32()]
			if e, e_ok := this.(Document); ok && e_ok {
				return ctx.GetInstanceForNode(host.htmlElement, e.DocumentElement())
			}
			return nil, v8.NewTypeError(iso, "Object not a Document")
		},
		nil,
		v8.ReadOnly,
	)
	return res
}
