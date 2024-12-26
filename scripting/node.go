package scripting

import (
	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

type ESNode struct {
	ESWrapper[browser.Node]
}

func NewESNode(host *ScriptHost) ESNode {
	return ESNode{NewESWrapper[browser.Node](host)}
}

func CreateNode(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := NewESNode(host)
	builder := NewConstructorBuilder[browser.Node](
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

	builder.instanceLookup = func(ctx *ScriptContext, this *v8.Object) (browser.Node, error) {
		instance, ok := ctx.GetCachedNode(this)
		if instance, e_ok := instance.(browser.Node); e_ok && ok {
			return instance, nil
		} else {
			return nil, v8.NewTypeError(iso, "Not an instance of NamedNodeMap")
		}
	}
	protoBuilder := builder.NewPrototypeBuilder()
	protoBuilder.CreateReadonlyProp2("nodeType",
		func(instance browser.Node, ctx *ScriptContext) (*v8.Value, error) {
			return v8.NewValue(iso, int32(instance.NodeType()))
		})
	protoBuilder.CreateReadonlyProp2(
		"childNodes",
		func(instance browser.Node, ctx *ScriptContext) (*v8.Value, error) {
			return ctx.GetInstanceForNodeByName("NodeList", instance.ChildNodes())
		},
	)
	protoBuilder.CreateFunction("contains",
		func(instance browser.Node, info argumentHelper) (result *v8.Value, err error) {
			var node browser.Node
			node, err = info.GetNodeArg(0)
			if err == nil {
				result, err = v8.NewValue(info.ctx.host.iso, instance.Contains(node))
			}
			return
		},
	)
	return builder.constructor
}

func (n ESNode) GetFirstChild(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
	ctx := n.host.MustGetContext(info.Context())
	node, err := n.GetInstance(info)
	if err != nil {
		return nil, err
	}
	result := node.FirstChild()
	return ctx.GetInstanceForNode(result)
}
