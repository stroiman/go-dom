package scripting

import (
	v8 "github.com/tommie/v8go"
)

type ConstructorBuilder[T any] struct {
	host           *ScriptHost
	constructor    *v8.FunctionTemplate
	instanceLookup func(*ScriptContext, *v8.Object) (T, error)
}

func NewIllegalConstructorBuilder[T any](host *ScriptHost) ConstructorBuilder[T] {
	constructor := v8.NewFunctionTemplateWithError(
		host.iso,
		func(args *v8.FunctionCallbackInfo) (*v8.Value, error) {
			return nil, v8.NewTypeError(host.iso, "Illegal Constructor")
		},
	)
	constructor.GetInstanceTemplate().SetInternalFieldCount(1)

	builder := ConstructorBuilder[T]{host: host,
		constructor: constructor,
	}
	return builder
}

func (c ConstructorBuilder[T]) NewPrototypeBuilder() PrototypeBuilder[T] {
	if c.instanceLookup == nil {
		panic("Cannot build prototype builder if instance lookup not specified")
	}
	return PrototypeBuilder[T]{
		host:   c.host,
		proto:  c.constructor.PrototypeTemplate(),
		lookup: c.instanceLookup,
	}
}

type PrototypeBuilder[T any] struct {
	host   *ScriptHost
	proto  *v8.ObjectTemplate
	lookup func(*ScriptContext, *v8.Object) (T, error)
}

func (h PrototypeBuilder[T]) CreateReadonlyProp(name string, fn func(T) string) {
	h.proto.SetAccessorPropertyCallback(name,
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := h.host.MustGetContext(arg.Context())
			instance, err := h.lookup(ctx, arg.This())
			if err != nil {
				return nil, err
			}
			value := fn(instance)
			return v8.NewValue(h.host.iso, value)
		}, nil, v8.ReadOnly)
}

func (h PrototypeBuilder[T]) CreateFunction(
	name string,
	fn func(T, argumentHelper) (*v8.Value, error),
) {
	h.proto.Set(
		name,
		v8.NewFunctionTemplateWithError(
			h.host.iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := h.host.MustGetContext(info.Context())
				instance, err := h.lookup(ctx, info.This())
				if err != nil {
					return nil, err
				}
				return fn(instance, argumentHelper{info})
				// return v8.NewValue(h.host.iso, value)
			},
		),
		v8.ReadOnly,
	)
}

func (h PrototypeBuilder[T]) CreateFunctionStringToString(name string, fn func(T, string) string) {
	h.CreateFunction(name, func(instance T, info argumentHelper) (*v8.Value, error) {
		value := fn(instance, info.Args()[0].String())
		return v8.NewValue(h.host.iso, value)
	})
}
