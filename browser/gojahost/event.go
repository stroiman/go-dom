package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
)

type EventWrapper struct {
	baseInstanceWrapper[dom.Event]
}

func NewEventWrapper(instance *GojaContext) wrapper { return newEventWrapper(instance) }
func newEventWrapper(instance *GojaContext) EventWrapper {
	return EventWrapper{newBaseInstanceWrapper[dom.Event](instance)}
}

type GojaEvent[T dom.Event] struct {
	Value *g.Object
	Event T
}

func ToBoolean(value g.Value) bool {
	return value != nil && value.ToBoolean()
}

func (w EventWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	arg1 := call.Argument(0).String()
	options := make([]dom.CustomEventOption, 0, 2)
	if arg2 := call.Argument(1); !g.IsUndefined(arg2) {
		if obj, ok := arg2.(*g.Object); ok {
			options = append(options, dom.EventCancelable(ToBoolean(obj.Get("cancelable"))))
			options = append(options, dom.EventBubbles(ToBoolean(obj.Get("bubbles"))))
		}
	}
	newInstance := dom.NewCustomEvent(arg1, options...)
	w.storeInternal(newInstance, call.This)
	return nil
}

func (w EventWrapper) PreventDefault(c g.FunctionCall) g.Value {
	w.getInstance(c).PreventDefault()
	return nil
}

func (w EventWrapper) GetType(c g.FunctionCall) g.Value {
	return w.instance.vm.ToValue(w.getInstance(c).Type())
}

func (w EventWrapper) initializePrototype(prototype *g.Object,
	vm *g.Runtime) {
	prototype.Set("preventDefault", w.PreventDefault)
	prototype.DefineAccessorProperty(
		"type",
		w.instance.vm.ToValue(w.GetType),
		nil,
		g.FLAG_TRUE,
		g.FLAG_TRUE,
	)
}

type CustomEventWrapper struct {
	EventWrapper
}

func NewCustomEventWrapper(instance *GojaContext) wrapper {
	return CustomEventWrapper{newEventWrapper(instance)}
}
