package scripting

import (
	v8 "github.com/tommie/v8go"
)

type PrototypeBuilder[T any] struct {
	host   *ScriptHost
	proto  *v8.ObjectTemplate
	lookup func(*ScriptContext) (T, error)
}

func (h PrototypeBuilder[T]) CreateReadonlyProp(name string, fn func(T) string) {
	h.proto.SetAccessorPropertyCallback(name,
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := h.host.MustGetContext(arg.Context())
			instance, err := h.lookup(ctx)
			if err != nil {
				return nil, err
			}
			value := fn(instance)
			return v8.NewValue(h.host.iso, value)
		}, nil, 0)
}
