// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func init() {
	RegisterJSClass("Node", "EventTarget", CreateNodePrototype)
}

func CreateNodePrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewNodeV8Wrapper(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("getRootNode", v8.NewFunctionTemplateWithError(iso, wrapper.GetRootNode))
	prototypeTmpl.Set("contains", v8.NewFunctionTemplateWithError(iso, wrapper.Contains))
	prototypeTmpl.Set("insertBefore", v8.NewFunctionTemplateWithError(iso, wrapper.InsertBefore))
	prototypeTmpl.Set("appendChild", v8.NewFunctionTemplateWithError(iso, wrapper.AppendChild))
	prototypeTmpl.Set("removeChild", v8.NewFunctionTemplateWithError(iso, wrapper.RemoveChild))

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

func (n NodeV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(n.host.iso, "Illegal Constructor")
}

func (n NodeV8Wrapper) GetRootNode(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.GetInstance(info)
	options, err1 := TryParseArgWithDefault(args, 0, n.DefaultGetRootNodeOptions, n.DecodeGetRootNodeOptions)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.GetRootNode(options)
		return ctx.GetInstanceForNode(result)
	}
	if err0 != nil {
		return nil, err0
	}
	result := instance.GetRootNode()
	return ctx.GetInstanceForNode(result)
}

func (n NodeV8Wrapper) Contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.GetInstance(info)
	other, err1 := TryParseArg(args, 0, n.DecodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Contains(other)
		return n.ToBoolean(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (n NodeV8Wrapper) InsertBefore(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.GetInstance(info)
	node, err1 := TryParseArg(args, 0, n.DecodeNode)
	child, err2 := TryParseArg(args, 1, n.DecodeNode)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.InsertBefore(node, child)
		if callErr != nil {
			return nil, callErr
		} else {
			return ctx.GetInstanceForNode(result)
		}
	}
	return nil, errors.New("Missing arguments")
}

func (n NodeV8Wrapper) AppendChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.GetInstance(info)
	node, err1 := TryParseArg(args, 0, n.DecodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.AppendChild(node)
		if callErr != nil {
			return nil, callErr
		} else {
			return ctx.GetInstanceForNode(result)
		}
	}
	return nil, errors.New("Missing arguments")
}

func (n NodeV8Wrapper) RemoveChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	instance, err0 := n.GetInstance(info)
	child, err1 := TryParseArg(args, 0, n.DecodeNode)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result, callErr := instance.RemoveChild(child)
		if callErr != nil {
			return nil, callErr
		} else {
			return ctx.GetInstanceForNode(result)
		}
	}
	return nil, errors.New("Missing arguments")
}

func (n NodeV8Wrapper) NodeName(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	instance, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.NodeName()
	return n.ToDOMString(ctx, result)
}

func (n NodeV8Wrapper) IsConnected(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	instance, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.IsConnected()
	return n.ToBoolean(ctx, result)
}

func (n NodeV8Wrapper) OwnerDocument(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	instance, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.OwnerDocument()
	return ctx.GetInstanceForNode(result)
}

func (n NodeV8Wrapper) ChildNodes(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	instance, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.ChildNodes()
	return n.ToNodeList(ctx, result)
}

func (n NodeV8Wrapper) FirstChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	instance, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.FirstChild()
	return ctx.GetInstanceForNode(result)
}

func (n NodeV8Wrapper) PreviousSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	instance, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.PreviousSibling()
	return ctx.GetInstanceForNode(result)
}

func (n NodeV8Wrapper) NextSibling(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	instance, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.NextSibling()
	return ctx.GetInstanceForNode(result)
}
