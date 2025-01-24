// This file is generated. Do not edit.

package v8host

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func init() {
	registerJSClass("Node", "EventTarget", createNodePrototype)
}

func createNodePrototype(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newNodeV8Wrapper(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("getRootNode", v8.NewFunctionTemplateWithError(iso, wrapper.getRootNode))
	prototypeTmpl.Set("contains", v8.NewFunctionTemplateWithError(iso, wrapper.contains))
	prototypeTmpl.Set("insertBefore", v8.NewFunctionTemplateWithError(iso, wrapper.insertBefore))
	prototypeTmpl.Set("appendChild", v8.NewFunctionTemplateWithError(iso, wrapper.appendChild))
	prototypeTmpl.Set("removeChild", v8.NewFunctionTemplateWithError(iso, wrapper.removeChild))

	prototypeTmpl.SetAccessorProperty("nodeType",
		v8.NewFunctionTemplateWithError(iso, wrapper.NodeType),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nodeName",
		v8.NewFunctionTemplateWithError(iso, wrapper.NodeName),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("isConnected",
		v8.NewFunctionTemplateWithError(iso, wrapper.IsConnected),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("ownerDocument",
		v8.NewFunctionTemplateWithError(iso, wrapper.OwnerDocument),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("childNodes",
		v8.NewFunctionTemplateWithError(iso, wrapper.ChildNodes),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("firstChild",
		v8.NewFunctionTemplateWithError(iso, wrapper.FirstChild),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("previousSibling",
		v8.NewFunctionTemplateWithError(iso, wrapper.PreviousSibling),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("nextSibling",
		v8.NewFunctionTemplateWithError(iso, wrapper.NextSibling),
		nil,
		v8.None)

	return constructor
}

func (n nodeV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(n.host.iso, "Illegal Constructor")
}

func (n nodeV8Wrapper) getRootNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.getInstance(info)
	options, err1 := tryParseArgWithDefault(args, 0, n.defaultGetRootNodeOptions, n.decodeGetRootNodeOptions)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetRootNode(options)
		return ctx.getInstanceForNode(result)
	}
	return nil, errors.New("Missing arguments")
}

func (n nodeV8Wrapper) contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.getInstance(info)
	other, err1 := tryParseArg(args, 0, n.decodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Contains(other)
		return n.toBoolean(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (n nodeV8Wrapper) insertBefore(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.getInstance(info)
	node, err1 := tryParseArg(args, 0, n.decodeNode)
	child, err2 := tryParseArg(args, 1, n.decodeNode)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.InsertBefore(node, child)
		if callErr != nil {
			return nil, callErr
		} else {
			return ctx.getInstanceForNode(result)
		}
	}
	return nil, errors.New("Missing arguments")
}

func (n nodeV8Wrapper) appendChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.getInstance(info)
	node, err1 := tryParseArg(args, 0, n.decodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.AppendChild(node)
		if callErr != nil {
			return nil, callErr
		} else {
			return ctx.getInstanceForNode(result)
		}
	}
	return nil, errors.New("Missing arguments")
}

func (n nodeV8Wrapper) removeChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.getInstance(info)
	child, err1 := tryParseArg(args, 0, n.decodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.RemoveChild(child)
		if callErr != nil {
			return nil, callErr
		} else {
			return ctx.getInstanceForNode(result)
		}
	}
	return nil, errors.New("Missing arguments")
}

func (n nodeV8Wrapper) NodeName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.NodeName()
	return n.toDOMString(ctx, result)
}

func (n nodeV8Wrapper) IsConnected(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.IsConnected()
	return n.toBoolean(ctx, result)
}

func (n nodeV8Wrapper) OwnerDocument(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.OwnerDocument()
	return ctx.getInstanceForNode(result)
}

func (n nodeV8Wrapper) ChildNodes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.ChildNodes()
	return n.toNodeList(ctx, result)
}

func (n nodeV8Wrapper) FirstChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.FirstChild()
	return ctx.getInstanceForNode(result)
}

func (n nodeV8Wrapper) PreviousSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling()
	return ctx.getInstanceForNode(result)
}

func (n nodeV8Wrapper) NextSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.mustGetContext(info.Context())
	instance, err := n.getInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling()
	return ctx.getInstanceForNode(result)
}
