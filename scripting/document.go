package scripting

import (
	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

type ESDocumentWrapper struct {
	ESElementContainerWrapper[Document]
}

func NewDocumentWrapper(host *ScriptHost) ESDocumentWrapper {
	return ESDocumentWrapper{NewESContainerWrapper[Document](host)}
}

func CreateDocumentPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewDocumentWrapper(host)
	builder := NewConstructorBuilder[Document](
		host,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			scriptContext := host.MustGetContext(info.Context())
			return scriptContext.CacheNode(info.This(), NewDocument(nil))
		},
	)
	wrapper.Install(builder.constructor)
	builder.SetDefaultInstanceLookup()
	protoBuilder := builder.NewPrototypeBuilder()
	protoBuilder.CreateReadonlyProp2(
		"location",
		func(instance Document, ctx *ScriptContext) (*v8.Value, error) {
			return ctx.v8ctx.Global().Get("location")
		},
	)
	instanceBuilder := builder.NewInstanceBuilder()
	instanceTemplate := instanceBuilder.proto
	instanceTemplate.SetInternalFieldCount(1)
	proto := builder.constructor.PrototypeTemplate()
	protoBuilder.CreateFunction(
		"createElement",
		func(instance Document, args argumentHelper) (val *v8.Value, err error) {
			var name string
			name, err = args.GetStringArg(0)
			if err == nil {
				e := instance.CreateElement(name)
				val, err = args.ctx.GetInstanceForNode(e)
			}
			return
		},
	)
	protoBuilder.CreateFunction(
		"createDocumentFragment",
		func(instance Document, args argumentHelper) (val *v8.Value, err error) {
			e := instance.CreateDocumentFragment()
			return args.ctx.GetInstanceForNode(e)
		},
	)

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
		"getElementById",
		v8.NewFunctionTemplateWithError(iso,
			func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(args.Context())
				this, ok := ctx.GetCachedNode(args.This())
				if doc, e_ok := this.(Document); ok && e_ok {
					element := doc.GetElementById(args.Args()[0].String())
					return ctx.GetInstanceForNode(element)
				}
				return nil, v8.NewTypeError(iso, "Object not a Document")
			}),
	)
	return builder.constructor
}
