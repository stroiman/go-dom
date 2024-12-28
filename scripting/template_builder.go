package scripting

import (
	"errors"

	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

type ConstructorBuilder[T any] struct {
	host           *ScriptHost
	constructor    *v8.FunctionTemplate
	instanceLookup func(*ScriptContext, *v8.Object) (T, error)
}

func NewConstructorBuilder[T any](
	host *ScriptHost,
	cb v8.FunctionCallbackWithError,
) ConstructorBuilder[T] {
	constructor := v8.NewFunctionTemplateWithError(
		host.iso,
		cb,
	)
	constructor.GetInstanceTemplate().SetInternalFieldCount(1)

	builder := ConstructorBuilder[T]{host: host,
		constructor: constructor,
	}
	return builder
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

func getInstanceFromThis[T any](ctx *ScriptContext, this *v8.Object) (instance T, err error) {
	cachedEntity, ok := ctx.GetCachedNode(this)
	if !ok {
		err = errors.New("No cached entity could be found for `this`")
		return
	}
	if i, e_ok := cachedEntity.(T); e_ok && ok {
		return i, nil
	} else {
		err = v8.NewTypeError(ctx.host.iso, "Not an object of the correct type")
		return
	}
}

func (c *ConstructorBuilder[T]) SetDefaultInstanceLookup() {
	c.instanceLookup = func(ctx *ScriptContext, this *v8.Object) (val T, err error) {
		instance, ok := ctx.GetCachedNode(this)
		if instance, e_ok := instance.(T); e_ok && ok {
			return instance, nil
		} else {
			err = v8.NewTypeError(ctx.host.iso, "Not an instance of NamedNodeMap")
			return
		}
	}
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

func (c ConstructorBuilder[T]) NewInstanceBuilder() PrototypeBuilder[T] {
	if c.instanceLookup == nil {
		panic("Cannot build prototype builder if instance lookup not specified")
	}
	return PrototypeBuilder[T]{
		host:   c.host,
		proto:  c.constructor.GetInstanceTemplate(),
		lookup: c.instanceLookup,
	}
}

type PrototypeBuilder[T any] struct {
	host   *ScriptHost
	proto  *v8.ObjectTemplate
	lookup func(*ScriptContext, *v8.Object) (T, error)
}

func (b ConstructorBuilder[T]) GetInstance(info *v8.FunctionCallbackInfo) (T, error) {
	ctx := b.host.MustGetContext(info.Context())
	return b.instanceLookup(ctx, info.This())
}

func (b PrototypeBuilder[T]) GetInstance(info *v8.FunctionCallbackInfo) (T, error) {
	ctx := b.host.MustGetContext(info.Context())
	return b.lookup(ctx, info.This())
}

type FunctionInfo[T any] struct {
	instance T
	ctx      *ScriptContext
}

func (h PrototypeBuilder[T]) CreateReadonlyProp2(
	name string,
	fn func(T, *ScriptContext) (*v8.Value, error),
) {
	h.proto.SetAccessorPropertyCallback(name,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := h.host.MustGetContext(info.Context())
			instance, err := h.GetInstance(info)
			if err != nil {
				return nil, err
			}
			return fn(instance, ctx)
		}, nil, v8.ReadOnly)
}

func (h PrototypeBuilder[T]) CreateReadonlyProp(name string, fn func(T) string) {
	h.proto.SetAccessorPropertyCallback(name,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			instance, err := h.GetInstance(info)
			if err != nil {
				return nil, err
			}
			value := fn(instance)
			return v8.NewValue(h.host.iso, value)
		}, nil, v8.ReadOnly)
}

func (h PrototypeBuilder[T]) CreateReadWriteProp(
	name string,
	get func(T) string,
	set func(T, string),
) {
	h.proto.SetAccessorPropertyCallback(name,
		func(arg *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := h.host.MustGetContext(arg.Context())
			instance, err := h.lookup(ctx, arg.This())
			if err != nil {
				return nil, err
			}
			value := get(instance)
			return v8.NewValue(h.host.iso, value)
		},
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			instance, err := h.GetInstance(info)
			if err != nil {
				return nil, err
			}
			newVal := info.Args()[0].String()
			set(instance, newVal)
			return nil, nil
		}, v8.None)
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
				instance, err := h.GetInstance(info)
				if err != nil {
					return nil, err
				}
				return fn(instance, argumentHelper{info, ctx, 0})
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

func TryParseArgs[T interface{}](
	ctx *ScriptContext,
	args []*v8.Value,
	index int,
	parsers ...func(*ScriptContext, *v8.Value) (T, error),
) (res T, err error) {
	if index >= len(args) {
		err = errors.New("Index out of range")
		return
	}
	for _, parser := range parsers {
		res, err = parser(ctx, args[index])
		if err == nil {
			return
		}
	}
	err = errors.New("TODO")
	return
}

func TryParseArg[T any](
	args *argumentHelper,
	index int,
	parser func(*ScriptContext, *v8.Value) (T, error),
) (result T, err error) {
	value := args.GetArg(index)
	if value == nil {
		return
	}
	return parser(args.ctx, value)
}

func GetBodyFromDocument(ctx *ScriptContext, val *v8.Value) (*browser.XHRRequestBody, error) {
	if val.IsNull() {
		return nil, nil
	}
	return nil, errors.New("Not supported yet")
}

func GetBodyFromXMLHttpRequestBodyInit(
	ctx *ScriptContext,
	val *v8.Value,
) (*browser.XHRRequestBody, error) {
	if !val.IsObject() {
		return nil, errors.New("Not supported yet")
	}
	obj := val.Object()
	node, ok := ctx.GetCachedNode(obj)
	if !ok {
		return nil, errors.New("Not a node")
	}
	formData, ok := node.(*browser.FormData)
	if ok {
		return browser.NewXHRRequestBodyOfFormData(formData), nil
	}
	return nil, errors.New("Not a node")
}
