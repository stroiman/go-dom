package goja_driver

import (
	"github.com/stroiman/go-dom/browser/dom"

	"github.com/dop251/goja"
	. "github.com/dop251/goja"
)

type EventTargetWrapper struct {
	instance *GojaInstance
}

func NewEventTargetWrapper(instance *GojaInstance) Wrapper {
	return EventTargetWrapper{instance}
}

type GojaEventListener struct {
	instance *GojaInstance
	v        goja.Value
	f        goja.Callable
}

func NewGojaEventListener(r *GojaInstance, v goja.Value) dom.EventHandler {
	f, ok := AssertFunction(v)
	if !ok {
		panic("TODO")
	}
	return &GojaEventListener{r, v, f}
}

func (h *GojaEventListener) HandleEvent(e dom.Event) error {
	customEvent := h.instance.globals["CustomEvent"]
	value := h.instance.vm.ToValue(e).(*Object)
	value.SetPrototype(customEvent.Prototype)
	_, err := h.f(value, value)
	return err
}

func (h *GojaEventListener) Equals(e dom.EventHandler) bool {
	if ge, ok := e.(*GojaEventListener); ok && ge.v.StrictEquals(h.v) {
		return true
	} else {
		return false
	}
}

func (w EventTargetWrapper) Constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	result := r.ToValue(dom.NewEventTarget()).(*Object)
	result.SetPrototype(call.This.Prototype())
	return result
}

func (w EventTargetWrapper) GetEventTarget(c FunctionCall) dom.EventTarget {
	if c.This == nil {
		panic("No this pointer")
	}
	if c.This == w.instance.vm.GlobalObject() {
		return w.instance.window
	}
	instance, ok := c.This.Export().(dom.EventTarget)
	if !ok {
		panic(w.instance.vm.NewTypeError("Not an event target"))
	}
	return instance

}

func (w EventTargetWrapper) InitializePrototype(prototype *Object,
	vm *Runtime) {
	createElement := vm.ToValue(func(c FunctionCall) Value {
		instance := w.GetEventTarget(c)
		name := c.Argument(0).String()
		instance.AddEventListener(name, NewGojaEventListener(w.instance, c.Argument(1)))
		return nil
	})

	dispatchEvent := vm.ToValue(func(c FunctionCall) Value {
		instance := w.GetEventTarget(c)
		if event, ok := c.Argument(0).Export().(dom.Event); ok {
			return vm.ToValue(instance.DispatchEvent(event))
		} else {
			panic(vm.NewTypeError("Not an event"))
		}
	})

	prototype.Set("addEventListener", createElement)
	prototype.Set("dispatchEvent", dispatchEvent)
}
