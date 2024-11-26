package scripting

import (
	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

func CreateNode(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	builder := NewConstructorBuilder[browser.Node](
		host,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return v8.Undefined(iso), nil
		},
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
	return builder.constructor
}
