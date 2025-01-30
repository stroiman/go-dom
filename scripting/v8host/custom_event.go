package v8host

import (
	"errors"
	"runtime/cgo"

	"github.com/gost-dom/browser/dom"
	v8 "github.com/tommie/v8go"
)

func createCustomEvent(host *V8ScriptHost) *v8.FunctionTemplate {
	iso := host.iso
	res := v8.NewFunctionTemplateWithError(
		iso,
		func(info *v8.FunctionCallbackInfo) (*v8.Value, error) {
			ctx := host.mustGetContext(info.Context())
			args := info.Args()
			if len(args) < 1 {
				return nil, v8.NewTypeError(iso, "Must have at least one constructor argument")
			}
			var eventOptions []dom.EventOption
			if len(args) > 1 {
				if options, err := args[1].AsObject(); err == nil {
					bubbles, err1 := options.Get("bubbles")
					cancelable, err2 := options.Get("cancelable")
					err = errors.Join(err1, err2)
					if err != nil {
						return nil, err
					}
					eventOptions = []dom.EventOption{
						dom.EventBubbles(bubbles.Boolean()),
						dom.EventCancelable(cancelable.Boolean()),
					}
				}
			}
			e := dom.NewCustomEvent(args[0].String(), eventOptions...)
			handle := cgo.NewHandle(e)
			ctx.addDisposer(handleDisposable(handle))
			info.This().SetInternalField(0, v8.NewValueExternalHandle(iso, handle))
			return v8.Undefined(iso), nil
		},
	)
	res.InstanceTemplate().SetInternalFieldCount(1)
	return res
}
