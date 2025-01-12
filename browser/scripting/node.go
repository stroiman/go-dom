package scripting

import (
	"errors"

	"github.com/stroiman/go-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

type NodeV8Wrapper struct {
	NodeV8WrapperBase[dom.Node]
}

func NewNodeV8Wrapper(host *ScriptHost) NodeV8Wrapper {
	return NodeV8Wrapper{NewNodeV8WrapperBase[dom.Node](host)}
}

func CreateNode(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewNodeV8Wrapper(host)
	builder := NewConstructorBuilder[dom.Node](
		host,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return v8.Undefined(iso), nil
		},
	)

	prototype := builder.constructor.PrototypeTemplate()
	prototype.SetAccessorProperty("firstChild",
		v8.NewFunctionTemplateWithError(iso, wrapper.GetFirstChild),
		nil,
		v8.ReadOnly,
	)
	prototype.Set("removeChild", v8.NewFunctionTemplateWithError(iso, wrapper.RemoveChild))

	builder.instanceLookup = func(ctx *ScriptContext, this *v8.Object) (dom.Node, error) {
		instance, ok := ctx.GetCachedNode(this)
		if instance, e_ok := instance.(dom.Node); e_ok && ok {
			return instance, nil
		} else {
			return nil, v8.NewTypeError(iso, "Not an instance of NamedNodeMap")
		}
	}
	protoBuilder := builder.NewPrototypeBuilder()
	protoBuilder.CreateReadonlyProp2("nodeType",
		func(instance dom.Node, ctx *ScriptContext) (*v8.Value, error) {
			return v8.NewValue(iso, int32(instance.NodeType()))
		})
	protoBuilder.CreateReadonlyProp2(
		"childNodes",
		func(instance dom.Node, ctx *ScriptContext) (*v8.Value, error) {
			return ctx.GetInstanceForNodeByName("NodeList", instance.ChildNodes())
		},
	)
	protoBuilder.CreateFunction("contains",
		func(instance dom.Node, info argumentHelper) (result *v8.Value, err error) {
			var node dom.Node
			node, err = info.GetNodeArg(0)
			if err == nil {
				result, err = v8.NewValue(info.ctx.host.iso, instance.Contains(node))
			}
			return
		},
	)
	protoBuilder.CreateFunction("appendChild",
		func(instance dom.Node, info argumentHelper) (result *v8.Value, err error) {
			var node dom.Node
			if node, err = info.GetNodeArg(0); err == nil {
				result = info.This().Value
				instance.AppendChild(node)
			}
			return
		},
	)
	protoBuilder.CreateFunction("insertBefore",
		func(instance dom.Node, info argumentHelper) (result *v8.Value, err error) {
			var resNode dom.Node
			node, err0 := info.GetNodeArg(0)
			refNode, err1 := info.GetNodeArg(1)
			if err = errors.Join(err0, err1); err == nil {
				resNode, err = instance.InsertBefore(node, refNode)
			}
			if err == nil {
				return info.ctx.GetInstanceForNode(resNode)
			}

			return
		},
	)
	return builder.constructor
}

func (n NodeV8Wrapper) GetFirstChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	node, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := node.FirstChild()
	return ctx.GetInstanceForNode(result)
}

func (n NodeV8Wrapper) RemoveChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	args := newArgumentHelper(n.host, info)
	child, err0 := args.GetNodeArg(0)
	parent, err1 := n.GetInstance(info)
	err := errors.Join(err0, err1)
	if err != nil {
		return nil, err
	}
	if result, err := parent.RemoveChild(child); err == nil {
		return ctx.GetInstanceForNode(result)
	} else {
		return nil, err
	}
}
