package v8host

import (
	. "github.com/gost-dom/browser/browser/dom"

	v8 "github.com/tommie/v8go"
)

type documentV8Wrapper struct {
	esElementContainerWrapper[Document]
}

func newDocumentV8Wrapper(host *V8ScriptHost) documentV8Wrapper {
	return documentV8Wrapper{newESContainerWrapper[Document](host)}
}

func (w documentV8Wrapper) BuildInstanceTemplate(constructor *v8.FunctionTemplate) {
	tmpl := constructor.InstanceTemplate()
	tmpl.SetAccessorProperty(
		"location",
		v8.NewFunctionTemplateWithError(
			w.scriptHost.iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := w.scriptHost.mustGetContext(info.Context())
				return ctx.v8ctx.Global().Get("location")
			},
		),
		nil,
		v8.None,
	)
}

func createDocumentPrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newDocumentV8Wrapper(host)
	builder := newConstructorBuilder[Document](
		host,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			scriptContext := host.mustGetContext(info.Context())
			return scriptContext.cacheNode(info.This(), NewDocument(nil))
		},
	)
	wrapper.Install(builder.constructor)
	builder.SetDefaultInstanceLookup()
	protoBuilder := builder.NewPrototypeBuilder()
	instanceBuilder := builder.NewInstanceBuilder()
	wrapper.BuildInstanceTemplate(builder.constructor)
	instanceTemplate := instanceBuilder.proto
	instanceTemplate.SetInternalFieldCount(1)
	proto := builder.constructor.PrototypeTemplate()
	protoBuilder.CreateFunction(
		"createElement",
		func(instance Document, args argumentHelper) (val *v8.Value, err error) {
			var name string
			name, err = args.getStringArg(0)
			if err == nil {
				e := instance.CreateElement(name)
				val, err = args.ctx.getInstanceForNode(e)
			}
			return
		},
	)
	protoBuilder.CreateFunction(
		"createDocumentFragment",
		func(instance Document, args argumentHelper) (val *v8.Value, err error) {
			e := instance.CreateDocumentFragment()
			return args.ctx.getInstanceForNode(e)
		},
	)

	proto.SetAccessorPropertyCallback("documentElement",
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.mustGetContext(arg.Context())
			this, ok := ctx.getCachedNode(arg.This())
			if e, e_ok := this.(Document); ok && e_ok {
				return ctx.getInstanceForNodeByName("HTMLElement", e.DocumentElement())
			}
			return nil, v8.NewTypeError(iso, "Object not a Document")
		},
		nil,
		v8.ReadOnly,
	)
	proto.SetAccessorPropertyCallback("head",
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.mustGetContext(arg.Context())
			this, ok := ctx.getCachedNode(arg.This())
			if e, e_ok := this.(Document); ok && e_ok {
				return ctx.getInstanceForNodeByName("HTMLElement", e.Head())
			}
			return nil, v8.NewTypeError(iso, "Object not a Document")
		},
		nil,
		v8.ReadOnly,
	)
	proto.SetAccessorPropertyCallback("body",
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.mustGetContext(arg.Context())
			this, ok := ctx.getCachedNode(arg.This())
			if e, e_ok := this.(Document); ok && e_ok {
				return ctx.getInstanceForNodeByName("HTMLElement", e.Body())
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
				ctx := host.mustGetContext(args.Context())
				this, ok := ctx.getCachedNode(args.This())
				if doc, e_ok := this.(Document); ok && e_ok {
					element := doc.GetElementById(args.Args()[0].String())
					return ctx.getInstanceForNode(element)
				}
				return nil, v8.NewTypeError(iso, "Object not a Document")
			}),
	)
	return builder.constructor
}
