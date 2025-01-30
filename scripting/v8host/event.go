package v8host

import (
	"errors"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/internal/entity"
	v8 "github.com/tommie/v8go"
)

func (w eventV8Wrapper) decodeEventInit(
	ctx *V8ScriptContext,
	v *v8.Value,
) (dom.EventOption, error) {
	var eventOptions []dom.EventOption
	options, err0 := v.AsObject()

	bubbles, err1 := options.Get("bubbles")
	cancelable, err2 := options.Get("cancelable")
	err := errors.Join(err0, err1, err2)
	if err == nil {
		eventOptions = []dom.EventOption{
			dom.EventBubbles(bubbles.Boolean()),
			dom.EventCancelable(cancelable.Boolean()),
		}
	}
	return dom.EventOptions(eventOptions), nil
}

func (w eventV8Wrapper) defaultEventInit() dom.EventOption {
	return dom.EventOptions(nil)
}

func (w eventV8Wrapper) CreateInstance(
	ctx *V8ScriptContext,
	this *v8.Object,
	type_ string,
	o dom.EventOption,
) (*v8.Value, error) {
	e := dom.NewEvent(type_, o)
	return w.store(e, ctx, this)
}

func (w eventV8Wrapper) toNullableEventTarget(
	ctx *V8ScriptContext,
	e dom.EventTarget,
) (*v8.Value, error) {
	if e == nil {
		return v8.Null(w.scriptHost.iso), nil
	}
	if entity, ok := e.(entity.Entity); ok {
		return ctx.getInstanceForNode(entity)
	}
	return nil, v8.NewError(w.iso(), "TODO, Not yet supported")
}
