// This file is generated. Do not edit.

package scripting

import (
	"errors"
	v8 "github.com/tommie/v8go"
)

func CreateDOMTokenListPrototype(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewDOMTokenListV8Wrapper(host)
	constructor := v8.NewFunctionTemplateWithError(iso, wrapper.Constructor)

	instanceTmpl := constructor.InstanceTemplate()
	instanceTmpl.SetInternalFieldCount(1)

	prototypeTmpl := constructor.PrototypeTemplate()
	prototypeTmpl.Set("item", v8.NewFunctionTemplateWithError(iso, wrapper.Item))
	prototypeTmpl.Set("contains", v8.NewFunctionTemplateWithError(iso, wrapper.Contains))
	prototypeTmpl.Set("add", v8.NewFunctionTemplateWithError(iso, wrapper.Add))
	prototypeTmpl.Set("remove", v8.NewFunctionTemplateWithError(iso, wrapper.Remove))
	prototypeTmpl.Set("toggle", v8.NewFunctionTemplateWithError(iso, wrapper.Toggle))
	prototypeTmpl.Set("replace", v8.NewFunctionTemplateWithError(iso, wrapper.Replace))
	prototypeTmpl.Set("supports", v8.NewFunctionTemplateWithError(iso, wrapper.Supports))

	prototypeTmpl.SetAccessorProperty("length",
		v8.NewFunctionTemplateWithError(iso, wrapper.Length),
		nil,
		v8.None)
	prototypeTmpl.SetAccessorProperty("value",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetValue),
		v8.NewFunctionTemplateWithError(iso, wrapper.SetValue),
		v8.None)

	wrapper.CustomInitialiser(constructor)
	return constructor
}

func (u DOMTokenListV8Wrapper) Constructor(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, v8.NewTypeError(u.host.iso, "Illegal Constructor")
}

func (u DOMTokenListV8Wrapper) Item(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.GetInstance(info)
	index, err1 := TryParseArg(args, 0, u.DecodeUnsignedLong)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Item(index)
		return u.ToNullableDOMString(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (u DOMTokenListV8Wrapper) Contains(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.GetInstance(info)
	token, err1 := TryParseArg(args, 0, u.DecodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		result := instance.Contains(token)
		return u.ToBoolean(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (u DOMTokenListV8Wrapper) Add(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.GetInstance(info)
	tokens, err1 := TryParseArg(args, 0, u.DecodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		callErr := instance.Add(tokens)
		return nil, callErr
	}
	return nil, errors.New("Missing arguments")
}

func (u DOMTokenListV8Wrapper) Remove(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.GetInstance(info)
	tokens, err1 := TryParseArg(args, 0, u.DecodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.Remove(tokens)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
}

func (u DOMTokenListV8Wrapper) Replace(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.GetInstance(info)
	token, err1 := TryParseArg(args, 0, u.DecodeDOMString)
	newToken, err2 := TryParseArg(args, 1, u.DecodeDOMString)
	if args.noOfReadArguments >= 2 {
		err := errors.Join(err0, err1, err2)
		if err != nil {
			return nil, err
		}
		result := instance.Replace(token, newToken)
		return u.ToBoolean(ctx, result)
	}
	return nil, errors.New("Missing arguments")
}

func (u DOMTokenListV8Wrapper) Supports(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	return nil, errors.New("Not implemented: DOMTokenList.supports")
}

func (u DOMTokenListV8Wrapper) Length(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.Length()
	return u.ToUnsignedLong(ctx, result)
}

func (u DOMTokenListV8Wrapper) GetValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := u.host.MustGetContext(info.Context())
	instance, err := u.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := instance.GetValue()
	return u.ToDOMString(ctx, result)
}

func (u DOMTokenListV8Wrapper) SetValue(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	args := newArgumentHelper(u.host, info)
	instance, err0 := u.GetInstance(info)
	val, err1 := TryParseArg(args, 0, u.DecodeDOMString)
	if args.noOfReadArguments >= 1 {
		err := errors.Join(err0, err1)
		if err != nil {
			return nil, err
		}
		instance.SetValue(val)
		return nil, nil
	}
	return nil, errors.New("Missing arguments")
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
