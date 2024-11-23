package scripting

import (
	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

func CreateAttr(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	builder := NewIllegalConstructorBuilder[Attr](host)
	builder.instanceLookup = func(ctx *ScriptContext, this *v8.Object) (Attr, error) {
		instance, ok := ctx.GetCachedNode(this)
		if e, e_ok := instance.(Attr); e_ok && ok {
			return e, nil
		} else {
			return nil, v8.NewTypeError(iso, "Not an instance of NamedNodeMap")
		}
	}
	proto := builder.NewPrototypeBuilder()
	proto.CreateReadonlyProp("name", Attr.Name)
	proto.CreateReadWriteProp("value", Attr.GetValue, Attr.SetValue)
	return builder.constructor
}

func CreateNamedNodeMap(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	builder := NewIllegalConstructorBuilder[NamedNodeMap](host)
	builder.instanceLookup = func(ctx *ScriptContext, this *v8.Object) (NamedNodeMap, error) {
		instance, ok := ctx.GetCachedNode(this)
		if e, e_ok := instance.(NamedNodeMap); e_ok && ok {
			return e, nil
		} else {
			return nil, v8.NewTypeError(iso, "Not an instance of NamedNodeMap")
		}
	}
	proto := builder.NewPrototypeBuilder()
	proto.CreateReadonlyProp2(
		"length",
		func(instance NamedNodeMap, ctx *ScriptContext) (*v8.Value, error) {
			return v8.NewValue(iso, int32(instance.Length()))
		},
	)
	proto.CreateFunction(
		"item",
		func(instance NamedNodeMap, info argumentHelper) (*v8.Value, error) {
			idx, err := info.GetInt32Arg(0)
			item := instance.Item(int(idx))
			if item != nil && err == nil {
				val, err := info.ctx.GetInstanceForNodeByName("Attr", item)
				return val, err
			}
			return nil, err
		},
	)

	return builder.constructor
}
