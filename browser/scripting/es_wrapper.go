package scripting

import (
	"errors"
	"runtime/cgo"

	"github.com/stroiman/go-dom/browser/dom"
	. "github.com/stroiman/go-dom/browser/dom"

	v8 "github.com/tommie/v8go"
)

// nodeV8WrapperBase serves as a helper for building v8 wrapping code around go objects.
// Generated code assumes that a wrapper type is used with specific helper
// methods implemented.
type nodeV8WrapperBase[T Entity] struct {
	converters
	host *ScriptHost
}

func newNodeV8WrapperBase[T Entity](host *ScriptHost) nodeV8WrapperBase[T] {
	return nodeV8WrapperBase[T]{converters{}, host}
}

func (w nodeV8WrapperBase[T]) getInstance(info *v8.FunctionCallbackInfo) (result T, err error) {
	if ctx, ok := w.host.GetContext(info.Context()); ok {
		if instance, ok := ctx.GetCachedNode(info.This()); ok {
			if typedInstance, ok := instance.(T); ok {
				return typedInstance, nil
			}
		}
		err = v8.NewTypeError(ctx.host.iso, "Not an instance of NamedNodeMap")
		return
	}
	err = errors.New("Could not get context")
	return
}

type converters struct{}

func (w converters) decodeUSVString(ctx *ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w converters) decodeByteString(ctx *ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w converters) decodeDOMString(ctx *ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w converters) decodeBoolean(ctx *ScriptContext, val *v8.Value) (bool, error) {
	return val.Boolean(), nil
}

func (w converters) decodeUnsignedLong(ctx *ScriptContext, val *v8.Value) (int, error) {
	return int(val.Int32()), nil
}

func (w converters) decodeNode(ctx *ScriptContext, val *v8.Value) (dom.Node, error) {
	if val.IsObject() {
		o := val.Object()
		cached, ok_1 := ctx.GetCachedNode(o)
		if node, ok_2 := cached.(dom.Node); ok_1 && ok_2 {
			return node, nil
		}
	}
	return nil, v8.NewTypeError(ctx.host.iso, "Must be a node")
}

func (w converters) getArgDOMString(args []*v8.Value, idx int) (result string, err error) {
	if idx >= len(args) {
		err = errors.New("Index out of range")
		return
	}
	result = args[idx].String()
	return
}

func (w converters) GetArgByteString(args []*v8.Value, idx int) (result string, err error) {
	return w.getArgDOMString(args, idx)
}

func (w converters) GetArgUSVString(args []*v8.Value, idx int) (result string, err error) {
	return w.getArgDOMString(args, idx)
}

func (w converters) GetArgBoolean(args []*v8.Value, idx int) (result bool, err error) {
	if idx >= len(args) {
		err = errors.New("Index out of range")
		return
	}
	result = args[idx].Boolean()
	return
}

func (w converters) ToNullableByteString(ctx *ScriptContext, str *string) (*v8.Value, error) {
	if str == nil {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, *str)
}

func (w converters) ToByteString(ctx *ScriptContext, str string) (*v8.Value, error) {
	if str == "" {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) ToDOMString(ctx *ScriptContext, str string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) ToNullableDOMString(ctx *ScriptContext, str *string) (*v8.Value, error) {
	if str == nil {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) ToUnsignedLong(ctx *ScriptContext, val int) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}

func (w converters) ToAny(ctx *ScriptContext, val string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}

func (w converters) ToUSVString(ctx *ScriptContext, str string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) ToUnsignedShort(ctx *ScriptContext, val int) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, uint32(val))
}

func (w converters) ToBoolean(ctx *ScriptContext, val bool) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}

func (w converters) ToNodeList(ctx *ScriptContext, val NodeList) (*v8.Value, error) {
	return ctx.GetInstanceForNodeByName("NodeList", val)
}

type handleReffedObject[T any] struct {
	host *ScriptHost
	converters
}

func NewHandleReffedObject[T any](host *ScriptHost) handleReffedObject[T] {
	return handleReffedObject[T]{
		host: host,
	}
}

func (o handleReffedObject[T]) Store(value T, ctx *ScriptContext, this *v8.Object) {
	handle := cgo.NewHandle(value)
	ctx.AddDisposer(HandleDisposable(handle))

	internalField := v8.NewValueExternalHandle(o.host.iso, handle)
	this.SetInternalField(0, internalField)
}

func getWrappedInstance[T any](object *v8.Object) T {
	field := object.GetInternalField(0)
	handle := field.ExternalHandle()
	return handle.Value().(T)
}

func (o handleReffedObject[T]) getInstance(info *v8.FunctionCallbackInfo) (T, error) {
	return getWrappedInstance[T](info.This()), nil
}
