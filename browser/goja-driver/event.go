package goja_driver

import (
	. "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
)

type EventWrapper struct {
	BaseInstanceWrapper[dom.Event]
}

func NewEventWrapper(instance *GojaInstance) Wrapper { return newEventWrapper(instance) }
func newEventWrapper(instance *GojaInstance) EventWrapper {
	return EventWrapper{NewBaseInstanceWrapper[dom.Event](instance)}
}

type GojaEvent[T dom.Event] struct {
	Value *Object
	Event T
}

func ToBoolean(value Value) bool {
	return value != nil && value.ToBoolean()
}

func (w EventWrapper) Constructor(call ConstructorCall, r *Runtime) *Object {
	arg1 := call.Argument(0).String()
	options := make([]dom.CustomEventOption, 0, 2)
	if arg2 := call.Argument(1); !IsUndefined(arg2) {
		if obj, ok := arg2.(*Object); ok {
			options = append(options, dom.EventCancelable(ToBoolean(obj.Get("cancelable"))))
			options = append(options, dom.EventBubbles(ToBoolean(obj.Get("bubbles"))))
		}
	}
	newInstance := dom.NewCustomEvent(arg1, options...)
	w.StoreInternal(newInstance, call.This)
	return nil
}

func (w EventWrapper) PreventDefault(c FunctionCall) Value {
	w.GetInstance(c).PreventDefault()
	return nil
}

func (w EventWrapper) GetType(c FunctionCall) Value {
	return w.instance.vm.ToValue(w.GetInstance(c).Type())
}

func (w EventWrapper) InitializePrototype(prototype *Object,
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

func NewCustomEventWrapper(instance *GojaInstance) Wrapper {
	return CustomEventWrapper{newEventWrapper(instance)}
}
