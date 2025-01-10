package goja_driver

import (
	. "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
)

type EventWrapper struct {
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
	result := r.ToValue(dom.NewCustomEvent(arg1, options...)).(*Object)
	result.SetPrototype(call.This.Prototype())
	return result
}

func (w EventWrapper) InitializePrototype(prototype *Object,
	vm *Runtime) {
	preventDefault := vm.ToValue(func(c FunctionCall) Value {
		event, ok := c.This.Export().(dom.Event)
		if !ok {
			panic(vm.NewTypeError("Instance is not an Event"))
		}
		event.PreventDefault()
		return nil
	})
	prototype.Set("preventDefault", preventDefault)
}

type CustomEventWrapper struct {
	Base EventWrapper
}

func (w CustomEventWrapper) Constructor(call ConstructorCall, r *Runtime) *Object {
	return w.Base.Constructor(call, r)
}
