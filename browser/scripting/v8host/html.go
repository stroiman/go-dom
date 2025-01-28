package v8host

import (
	"github.com/gost-dom/browser/browser/dom"
	. "github.com/gost-dom/browser/browser/dom"

	v8 "github.com/tommie/v8go"
)

type esElementContainerWrapper[T ElementContainer] struct {
	nodeV8WrapperBase[T]
}

func newESContainerWrapper[T ElementContainer](host *V8ScriptHost) esElementContainerWrapper[T] {
	return esElementContainerWrapper[T]{newNodeV8WrapperBase[T](host)}
}

func (e esElementContainerWrapper[T]) Install(ft *v8.FunctionTemplate) {
	prototype := ft.PrototypeTemplate()
	prototype.Set(
		"querySelector",
		v8.NewFunctionTemplateWithError(e.scriptHost.iso, e.QuerySelector),
	)
	prototype.Set(
		"querySelectorAll",
		v8.NewFunctionTemplateWithError(e.scriptHost.iso, e.QuerySelectorAll),
	)
	prototype.Set("append", v8.NewFunctionTemplateWithError(e.scriptHost.iso, e.append))
}

func (e esElementContainerWrapper[T]) QuerySelector(
	args *v8.FunctionCallbackInfo,
) (*v8.Value, error) {
	host := e.scriptHost
	iso := e.scriptHost.iso
	ctx := host.mustGetContext(args.Context())
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
	host := e.scriptHost
	iso := e.scriptHost.iso
	ctx := host.mustGetContext(args.Context())
	this, ok := ctx.getCachedNode(args.This())
	if doc, e_ok := this.(ElementContainer); ok && e_ok {
		nodeList, err := doc.QuerySelectorAll(args.Args()[0].String())
		if err != nil {
			return nil, err
		}
		return ctx.getInstanceForNodeByName("NodeList", nodeList)
	}
	return nil, v8.NewTypeError(iso, "Object not a Document")
}

func (w esElementContainerWrapper[T]) append(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	var err error
	ctx := w.scriptHost.mustGetContext(info.Context())
	args := info.Args()
	nodes := make([]dom.Node, len(args))
	for i, a := range args {
		if nodes[i], err = w.decodeNodeOrText(ctx, a); err != nil {
			return nil, err
		}
	}
	i, err := w.getInstance(info)
	if err == nil {
		i.Append(nodes...)
	}
	return nil, nil
}
