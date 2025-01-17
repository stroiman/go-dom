package v8host

import (
	"errors"
	"runtime/cgo"

	"github.com/stroiman/go-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

type v8EventListener struct {
	ctx *V8ScriptContext
	val *v8.Value
}

func NewV8EventListener(ctx *V8ScriptContext, val *v8.Value) dom.EventHandler {
	return v8EventListener{ctx, val}
}

func (l v8EventListener) HandleEvent(e dom.Event) error {
	f, err := l.val.AsFunction()
	if err == nil {
		var event *v8.Value
		event, err = l.ctx.getInstanceForNode(e)
		if err == nil {
			_, err = f.Call(l.val, event)
		}
	}
	return err
}

func (l v8EventListener) Equals(other dom.EventHandler) bool {
	x, ok := other.(v8EventListener)
	return ok && x.val == l.val
}

func createEvent(host *V8ScriptHost) *v8.FunctionTemplate {
	result := NewIllegalConstructorBuilder[dom.Event](host)
	result.SetDefaultInstanceLookup()
	protoBuilder := result.NewPrototypeBuilder()
	protoBuilder.CreateReadonlyProp("type", dom.Event.Type)
	protoBuilder.CreateFunction(
		"preventDefault",
		func(instance dom.Event, a argumentHelper) (*v8.Value, error) {
			instance.PreventDefault()
			return nil, nil
		},
	)
	return result.constructor
}

func createCustomEvent(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	res := v8.NewFunctionTemplateWithError(
		iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			args := info.Args()
			if len(args) < 1 {
				return nil, v8.NewTypeError(iso, "Must have at least one constructor argument")
			}
			var eventOptions []dom.CustomEventOption
			if len(args) > 1 {
				if options, err := args[1].AsObject(); err == nil {
					bubbles, err1 := options.Get("bubbles")
					cancelable, err2 := options.Get("cancelable")
					err = errors.Join(err1, err2)
					if err != nil {
						return nil, err
					}
					eventOptions = []dom.CustomEventOption{
						dom.EventBubbles(bubbles.Boolean()),
						dom.EventCancelable(cancelable.Boolean()),
					}
				}
			}
			e := dom.NewCustomEvent(args[0].String(), eventOptions...)
			handle := cgo.NewHandle(e)
			ctx.AddDisposer(HandleDisposable(handle))
			info.This().SetInternalField(0, v8.NewValueExternalHandle(iso, handle))
			return v8.Undefined(iso), nil
		},
	)
	res.InstanceTemplate().SetInternalFieldCount(1)
	return res
}

type eventTargetV8Wrapper struct {
	handleReffedObject[dom.EventTarget]
}

func newEventTargetV8Wrapper(host *V8ScriptHost) eventTargetV8Wrapper {
	return eventTargetV8Wrapper{newHandleReffedObject[dom.EventTarget](host)}
}

func (w eventTargetV8Wrapper) getInstance(info *v8.FunctionCallbackInfo) (dom.EventTarget, error) {
	if info.This().GetInternalField(0).IsExternal() {
		return w.handleReffedObject.getInstance(info)
	} else {
		ctx := w.host.MustGetContext(info.Context())
		entity, ok := ctx.getCachedNode(info.This())
		if ok {
			if target, ok := entity.(dom.EventTarget); ok {
				return target, nil
			}
		}
		return nil, errors.New("Stored object is not an event target")
	}
}

func createEventTarget(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	wrapper := newEventTargetV8Wrapper(host)
	res := v8.NewFunctionTemplate(
		iso,
		func(info *v8.FunctionCallbackInfo) *v8.Value {
			ctx := host.MustGetContext(info.Context())
			wrapper.store(dom.NewEventTarget(), ctx, info.This())
			return v8.Undefined(iso)
		},
	)
	proto := res.PrototypeTemplate()
	proto.Set(
		"addEventListener",
		v8.NewFunctionTemplateWithError(iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(info.Context())
				target, err := wrapper.getInstance(info)
				if err != nil {
					return nil, err
				}
				args := newArgumentHelper(host, info)
				eventType, e1 := args.getStringArg(0)
				fn, e2 := args.getFunctionArg(1)
				err = errors.Join(e1, e2)
				if err == nil {
					listener := NewV8EventListener(ctx, fn.Value)
					target.AddEventListener(eventType, listener)
				}
				return v8.Undefined(iso), err
			}), v8.ReadOnly)
	proto.Set(
		"dispatchEvent",
		v8.NewFunctionTemplateWithError(iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				target, err := wrapper.getInstance(info)
				if err != nil {
					return nil, err
				}
				e := info.Args()[0]
				handle := e.Object().GetInternalField(0).ExternalHandle()
				if evt, ok := handle.Value().(dom.Event); ok {
					return v8.NewValue(iso, target.DispatchEvent(evt))
				} else {
					return nil, v8.NewTypeError(iso, "Not an Event")
				}
			}), v8.ReadOnly)
	instanceTemplate := res.InstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	return res
}
