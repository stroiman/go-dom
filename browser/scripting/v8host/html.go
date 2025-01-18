package v8host

import (
	. "github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

type esElementContainerWrapper[T ElementContainer] struct {
	nodeV8WrapperBase[T]
}

func NewESContainerWrapper[T ElementContainer](host *V8ScriptHost) esElementContainerWrapper[T] {
	return esElementContainerWrapper[T]{newNodeV8WrapperBase[T](host)}
}

func (e esElementContainerWrapper[T]) Install(ft *v8.FunctionTemplate) {
	prototype := ft.PrototypeTemplate()
	prototype.Set("querySelector", v8.NewFunctionTemplateWithError(e.host.iso, e.QuerySelector))
	prototype.Set(
		"querySelectorAll",
		v8.NewFunctionTemplateWithError(e.host.iso, e.QuerySelectorAll),
	)
}

func (e esElementContainerWrapper[T]) QuerySelector(
	args *v8.FunctionCallbackInfo,
) (*v8.Value, error) {
	host := e.host
	iso := e.host.iso
	ctx := host.MustGetContext(args.Context())
	this, ok := ctx.getCachedNode(args.This())
	if doc, e_ok := this.(ElementContainer); ok && e_ok {
		node, err := doc.QuerySelector(args.Args()[0].String())
		if err != nil {
			return nil, err
		}
		if node == nil {
			return v8.Null(iso), nil
		}
		return ctx.getInstanceForNode(node)
	}
	return nil, v8.NewTypeError(iso, "Object not a Document")
}

func (e esElementContainerWrapper[T]) QuerySelectorAll(
	args *v8.FunctionCallbackInfo,
) (*v8.Value, error) {
	host := e.host
	iso := e.host.iso
	ctx := host.MustGetContext(args.Context())
	this, ok := ctx.getCachedNode(args.This())
	if doc, e_ok := this.(ElementContainer); ok && e_ok {
		nodeList, err := doc.QuerySelectorAll(args.Args()[0].String())
		if err != nil {
			return nil, err
		}
		return ctx.GetInstanceForNodeByName("NodeList", nodeList)
	}
	return nil, v8.NewTypeError(iso, "Object not a Document")
}