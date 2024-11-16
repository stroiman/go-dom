package scripting

import (
	"runtime"

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

func (l v8EventListener) HandleEvent(e browser.Event) {
	f, err := l.val.AsFunction()
	if err != nil {
		panic(err)
	}
	f.Call(l.val, v8.Undefined(l.iso))
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
			args := info.Args()
			if len(args) < 1 {
				return nil, v8.NewTypeError(iso, "Must have at least one constructor argument")
			}
			e := browser.NewCustomEvent(args[0].String())
			t := runtime.Pinner{}
			t.Pin(e)
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
					args := info.Args()
					listener := NewV8EventListener(iso, args[1])
					target.AddEventListener(args[0].String(), listener)
					return v8.Undefined(iso), nil
				} else {
					return nil, v8.NewTypeError(iso, "What?")
				}
			}), v8.ReadOnly)
	proto.Set(
		"dispatchEvent",
		v8.NewFunctionTemplateWithError(iso,
			func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
				ctx := host.MustGetContext(info.Context())
				if target, ok := ctx.domNodes[info.This().GetInternalField(0).Int32()].(browser.EventTarget); ok {
					e := info.Args()[0]
					intf := e.Object().GetInternalField(0).ExternalInterface()
					evt, ok := intf.(browser.Event)
					if !ok {
						panic("Not an event")
					}
					target.DispatchEvent(evt)
					//target.DispatchEvent(browser.NewCustomEvent("custom"))
					return v8.Undefined(iso), nil
				} else {
					return nil, v8.NewTypeError(iso, "What the dispatch?")
				}
			}), v8.ReadOnly)
	instanceTemplate := res.GetInstanceTemplate()
	instanceTemplate.SetInternalFieldCount(1)
	return res
}
