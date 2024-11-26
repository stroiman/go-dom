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
			this, ok := ctx.GetCachedNode(arg.This())
			if e, e_ok := this.(Document); ok && e_ok {
				return ctx.GetInstanceForNodeByName("HTMLElement", e.DocumentElement())
			}
			return nil, v8.NewTypeError(iso, "Object not a Document")
		},
		nil,
		v8.ReadOnly,
	)
	proto.SetAccessorPropertyCallback("head",
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(arg.Context())
			this, ok := ctx.GetCachedNode(arg.This())
			if e, e_ok := this.(Document); ok && e_ok {
				return ctx.GetInstanceForNodeByName("HTMLElement", e.Head())
			}
			return nil, v8.NewTypeError(iso, "Object not a Document")
		},
		nil,
		v8.ReadOnly,
	)
	proto.SetAccessorPropertyCallback("body",
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(arg.Context())
			this, ok := ctx.GetCachedNode(arg.This())
			if e, e_ok := this.(Document); ok && e_ok {
				return ctx.GetInstanceForNodeByName("HTMLElement", e.Body())
			}
			return nil, v8.NewTypeError(iso, "Object not a Document")
		},
		nil,
		v8.ReadOnly,
	)
	proto.Set(
		"querySelector",
		v8.NewFunctionTemplateWithError(iso,
			func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(args.Context())
				this, ok := ctx.GetCachedNode(args.This())
				if doc, e_ok := this.(Document); ok && e_ok {
					node, err := doc.QuerySelector(args.Args()[0].String())
					if err != nil {
						return nil, err
					}
					if node == nil {
						return v8.Null(iso), nil
					}
					return ctx.GetInstanceForNode(node)
				}
				return nil, v8.NewTypeError(iso, "Object not a Document")
			}),
	)
	proto.Set(
		"querySelectorAll",
		v8.NewFunctionTemplateWithError(iso,
			func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(args.Context())
				this, ok := ctx.GetCachedNode(args.This())
				if doc, e_ok := this.(Document); ok && e_ok {
					nodeList, err := doc.QuerySelectorAll(args.Args()[0].String())
					if err != nil {
						return nil, err
					}
					return ctx.GetInstanceForNodeByName("NodeList", nodeList)
				}
				return nil, v8.NewTypeError(iso, "Object not a Document")
			}),
	)
	proto.Set(
		"getElementById",
		v8.NewFunctionTemplateWithError(iso,
			func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(args.Context())
				this, ok := ctx.GetCachedNode(args.This())
				if doc, e_ok := this.(Document); ok && e_ok {
					element := doc.GetElementById(args.Args()[0].String())
					return ctx.GetInstanceForNodeByName("Element", element)
				}
				return nil, v8.NewTypeError(iso, "Object not a Document")
			}),
	)
	return res
}
