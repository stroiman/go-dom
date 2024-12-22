package scripting

import (
	"errors"

	. "github.com/stroiman/go-dom/browser"

	v8 "github.com/tommie/v8go"
)

// ESWrapper serves as a helper for building v8 wrapping code around go objects.
// Generated code assumes that a wrapper type is used with specific helper
// methods implemented.
type ESWrapper[T Entity] struct {
	Converters
	host *ScriptHost
}

func NewESWrapper[T Entity](host *ScriptHost) ESWrapper[T] {
	return ESWrapper[T]{Converters{}, host}
}

func (w ESWrapper[T]) GetInstance(info *v8.FunctionCallbackInfo) (result T, err error) {
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

type Converters struct{}

func (w Converters) DecodeUSVString(ctx *ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w Converters) GetArgDOMString(args []*v8.Value, idx int) (result string, err error) {
	if idx >= len(args) {
		err = errors.New("Index out of range")
		return
	}
	result = args[idx].String()
	return
}

func (w Converters) GetArgByteString(args []*v8.Value, idx int) (result string, err error) {
	return w.GetArgDOMString(args, idx)
}
func (w Converters) GetArgUSVString(args []*v8.Value, idx int) (result string, err error) {
	return w.GetArgDOMString(args, idx)
}
func (w Converters) GetArgBoolean(args []*v8.Value, idx int) (result bool, err error) {
	if idx >= len(args) {
		err = errors.New("Index out of range")
		return
	}
	result = args[idx].Boolean()
	return
}

func (w Converters) ToNullableByteString(ctx *ScriptContext, str *string) (*v8.Value, error) {
	if str == nil {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, str)
}

func (w Converters) ToByteString(ctx *ScriptContext, str string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, str)
}

func (w Converters) ToUSVString(ctx *ScriptContext, str string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, str)
}

func (w Converters) ToUnsignedShort(ctx *ScriptContext, val int) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}

func (w Converters) ToUnsignedLong(ctx *ScriptContext, val int) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}

func (w Converters) ToBoolean(ctx *ScriptContext, val bool) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}
