package scripting

import (
	"errors"

	"github.com/stroiman/go-dom/browser"
	v8 "github.com/tommie/v8go"
)

type v8EventListener struct {
	iso *v8.Isolate
	val *v8.Value
}

func NewV8EventListener(iso *v8.Isolate, val *v8.Value) browser.EventHandler {
	return v8EventListener{iso, val}
}

func (l v8EventListener) HandleEvent(e browser.Event) error {
	f, err := l.val.AsFunction()
	if err == nil {
		_, err = f.Call(l.val, v8.Undefined(l.iso))
	}
	return err
}

func (l v8EventListener) Equals(other browser.EventHandler) bool {
	x, ok := other.(v8EventListener)
	return ok && x.val == l.val
}

func CreateCustomEvent(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	res := v8.NewFunctionTemplateWithError(
		iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.MustGetContext(info.Context())
			args := info.Args()
			if len(args) < 1 {
				return nil, v8.NewTypeError(iso, "Must have at least one constructor argument")
			}
			e := browser.NewCustomEvent(args[0].String())
			// TODO: Better memory management
			ctx.pinner.Pin(e)
			info.This().SetInternalField(0, v8.NewExternalFromInterface(iso, e))
			return v8.Undefined(iso), nil
		},
	)
	res.GetInstanceTemplate().SetInternalFieldCount(1)
	return res
}

func CreateEventTarget(host *ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	res := v8.NewFunctionTemplate(
		iso,
		func(info *v8.FunctionCallbackInfo) *v8.Value {
			ctx := host.MustGetContext(info.Context())
			ctx.CacheNode(info.This(), browser.NewEventTarget())
			return v8.Undefined(iso)
		},
	)
	proto := res.PrototypeTemplate()
	proto.Set(
		"addEventListener",
		v8.NewFunctionTemplateWithError(iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(info.Context())
				if target, ok := ctx.domNodes[info.This().GetInternalField(0).Int32()].(browser.EventTarget); ok {
					args := newArgumentHelper(host, info)
					eventType, e1 := args.GetStringArg(0)
					fn, e2 := args.GetFunctionArg(1)
					err := errors.Join(e1, e2)
					if err == nil {
						listener := NewV8EventListener(iso, fn.Value)
						target.AddEventListener(eventType, listener)
					}
					return v8.Undefined(iso), err
				} else {
					return nil, v8.NewTypeError(iso, "What?")
				}
			}), v8.ReadOnly)
	proto.Set(
		"dispatchEvent",
		v8.NewFunctionTemplateWithError(iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(info.Context())
				this, _ := ctx.GetCachedNode(info.This())
				target, ok := this.(browser.EventTarget)
				if !ok {
					return nil, v8.NewTypeError(iso, "Object not an EventTarget")
				}
				e := info.Args()[0]
				intf := e.Object().GetInternalField(0).ExternalInterface()
				if evt, ok := intf.(browser.Event); ok {
					target.DispatchEvent(evt)
				} else {
					return nil, v8.NewTypeError(iso, "Not an Event")
				}
				return v8.Undefined(iso), nil
			}), v8.ReadOnly)
	instanceTemplate := res.GetInstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	return res
}
