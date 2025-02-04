package v8host

import (
	"errors"
	"runtime/cgo"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/entity"

	v8 "github.com/tommie/v8go"
)

// nodeV8WrapperBase serves as a helper for building v8 wrapping code around go objects.
// Generated code assumes that a wrapper type is used with specific helper
// methods implemented.
type nodeV8WrapperBase[T entity.Entity] struct {
	converters
	scriptHost *V8ScriptHost
}

func (w nodeV8WrapperBase[T]) iso() *v8.Isolate {
	return w.scriptHost.iso
}

func (w nodeV8WrapperBase[T]) mustGetContext(info *v8.FunctionCallbackInfo) *V8ScriptContext {
	return w.scriptHost.mustGetContext(info.Context())
}

func newNodeV8WrapperBase[T entity.Entity](host *V8ScriptHost) nodeV8WrapperBase[T] {
	return nodeV8WrapperBase[T]{converters{}, host}
}

func (w nodeV8WrapperBase[T]) getInstance(info *v8.FunctionCallbackInfo) (result T, err error) {
	if ctx, ok := w.scriptHost.netContext(info.Context()); ok {
		if instance, ok := ctx.getCachedNode(info.This()); ok {
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

func (w nodeV8WrapperBase[T]) store(
	value T,
	ctx *V8ScriptContext,
	this *v8.Object,
) (*v8.Value, error) {
	val := this.Value
	objectId := value.ObjectId()
	ctx.v8nodes[objectId] = val
	ctx.domNodes[objectId] = value
	internal, err := v8.NewValue(ctx.host.iso, objectId)
	if err != nil {
		return nil, err
	}
	this.SetInternalField(0, internal)
	return val, nil
}

type converters struct{}

func (w converters) decodeUSVString(ctx *V8ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w converters) decodeByteString(ctx *V8ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w converters) decodeDOMString(ctx *V8ScriptContext, val *v8.Value) (string, error) {
	return val.String(), nil
}

func (w converters) decodeBoolean(ctx *V8ScriptContext, val *v8.Value) (bool, error) {
	return val.Boolean(), nil
}

func (w converters) decodeLong(ctx *V8ScriptContext, val *v8.Value) (int, error) {
	return int(val.Int32()), nil
}

func (w converters) decodeUnsignedLong(ctx *V8ScriptContext, val *v8.Value) (int, error) {
	return int(val.Uint32()), nil
}

func (w converters) decodeNode(ctx *V8ScriptContext, val *v8.Value) (dom.Node, error) {
	if val.IsObject() {
		o := val.Object()
		cached, ok_1 := ctx.getCachedNode(o)
		if node, ok_2 := cached.(dom.Node); ok_1 && ok_2 {
			return node, nil
		}
	}
	return nil, v8.NewTypeError(ctx.host.iso, "Must be a node")
}

func (w converters) decodeHTMLElement(
	ctx *V8ScriptContext,
	val *v8.Value,
) (html.HTMLElement, error) {
	if val.IsObject() {
		o := val.Object()
		cached, ok_1 := ctx.getCachedNode(o)
		if node, ok_2 := cached.(html.HTMLElement); ok_1 && ok_2 {
			return node, nil
		}
	}
	return nil, v8.NewTypeError(ctx.host.iso, "Must be a node")
}
func (c converters) defaultHTMLElement() html.HTMLElement { return nil }

func (w converters) decodeNodeOrText(ctx *V8ScriptContext, val *v8.Value) (dom.Node, error) {
	if val.IsString() {
		return NewText(val.String()), nil
	}
	return w.decodeNode(ctx, val)
}

func (w converters) toNullableByteString(ctx *V8ScriptContext, str *string) (*v8.Value, error) {
	if str == nil {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, *str)
}

func (w converters) toByteString(ctx *V8ScriptContext, str string) (*v8.Value, error) {
	if str == "" {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) toDOMString(ctx *V8ScriptContext, str string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) toNullableDOMString(ctx *V8ScriptContext, str *string) (*v8.Value, error) {
	if str == nil {
		return v8.Null(ctx.host.iso), nil
	}
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) toUnsignedLong(ctx *V8ScriptContext, val int) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, uint32(val))
}

func (w converters) toAny(ctx *V8ScriptContext, val string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}

func (w converters) toUSVString(ctx *V8ScriptContext, str string) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, str)
}

func (w converters) toUnsignedShort(ctx *V8ScriptContext, val int) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, uint32(val))
}

func (w converters) toBoolean(ctx *V8ScriptContext, val bool) (*v8.Value, error) {
	return v8.NewValue(ctx.host.iso, val)
}

func (w converters) toNodeList(ctx *V8ScriptContext, val NodeList) (*v8.Value, error) {
	return ctx.getInstanceForNodeByName("NodeList", val)
}

func (w converters) toHTMLFormControlsCollection(
	ctx *V8ScriptContext,
	val NodeList,
) (*v8.Value, error) {
	return w.toNodeList(ctx, val)
}

type handleReffedObject[T any] struct {
	scriptHost *V8ScriptHost
	converters
}

func (o handleReffedObject[T]) mustGetContext(info *v8.FunctionCallbackInfo) *V8ScriptContext {
	return o.scriptHost.mustGetContext(info.Context())
}

func newHandleReffedObject[T any](host *V8ScriptHost) handleReffedObject[T] {
	return handleReffedObject[T]{
		scriptHost: host,
	}
}

func (o handleReffedObject[T]) store(value T, ctx *V8ScriptContext, this *v8.Object) {
	handle := cgo.NewHandle(value)
	ctx.addDisposer(handleDisposable(handle))

	internalField := v8.NewValueExternalHandle(o.scriptHost.iso, handle)
	this.SetInternalField(0, internalField)
}

func getWrappedInstance[T any](object *v8.Object) (res T, err error) {
	field := object.GetInternalField(0)
	handle := field.ExternalHandle()
	var ok bool
	res, ok = handle.Value().(T)
	if !ok {
		err = errors.New("Not a valid type stored in the handle")
	}
	return
}

func (o handleReffedObject[T]) getInstance(info *v8.FunctionCallbackInfo) (T, error) {
	return getWrappedInstance[T](info.This())
}
