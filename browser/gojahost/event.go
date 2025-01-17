package gojahost

import (
	. "github.com/dop251/goja"
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
	Value *Object
	Event T
}

func ToBoolean(value Value) bool {
	return value != nil && value.ToBoolean()
}

func (w EventWrapper) constructor(call ConstructorCall, r *Runtime) *Object {
	arg1 := call.Argument(0).String()
	options := make([]dom.CustomEventOption, 0, 2)
	if arg2 := call.Argument(1); !IsUndefined(arg2) {
		if obj, ok := arg2.(*Object); ok {
			options = append(options, dom.EventCancelable(ToBoolean(obj.Get("cancelable"))))
			options = append(options, dom.EventBubbles(ToBoolean(obj.Get("bubbles"))))
		}
	}
	newInstance := dom.NewCustomEvent(arg1, options...)
	w.storeInternal(newInstance, call.This)
	return nil
}

func (w EventWrapper) PreventDefault(c FunctionCall) Value {
	w.getInstance(c).PreventDefault()
	return nil
}

func (w EventWrapper) GetType(c FunctionCall) Value {
	return w.instance.vm.ToValue(w.getInstance(c).Type())
}

func (w EventWrapper) initializePrototype(prototype *Object,
	vm *Runtime) {
	prototype.Set("preventDefault", w.PreventDefault)
	prototype.DefineAccessorProperty(
		"type",
		w.instance.vm.ToValue(w.GetType),
		nil,
		FLAG_TRUE,
		FLAG_TRUE,
	)
}

type CustomEventWrapper struct {
	EventWrapper
}

func NewCustomEventWrapper(instance *GojaContext) wrapper {
	return CustomEventWrapper{newEventWrapper(instance)}
}
