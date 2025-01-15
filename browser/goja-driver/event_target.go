package goja_driver

import (
	"github.com/stroiman/go-dom/browser/dom"

	"github.com/dop251/goja"
	. "github.com/dop251/goja"
)

type EventTargetWrapper struct {
	baseInstanceWrapper[dom.EventTarget]
}

func NewEventTargetWrapper(instance *GojaInstance) Wrapper {
	return EventTargetWrapper{newBaseInstanceWrapper[dom.EventTarget](instance)}
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
	obj := h.instance.vm.CreateObject(customEvent.Prototype)
	customEvent.Wrapper.storeInternal(e, obj)
	_, err := h.f(obj, obj)
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
	newInstance := dom.NewEventTarget()
	w.storeInternal(newInstance, call.This)
	return nil
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

func (w EventTargetWrapper) AddEventListener(c FunctionCall) Value {
	instance := w.GetInstance(c)
	name := c.Argument(0).String()
	instance.AddEventListener(name, NewGojaEventListener(w.instance, c.Argument(1)))
	return nil
}

func (w EventTargetWrapper) DispatchEvent(c FunctionCall) Value {
	instance := w.GetInstance(c)
	internal := c.Argument(0).(*Object).GetSymbol(w.instance.wrappedGoObj).Export()
	if event, ok := internal.(dom.Event); ok {
		return w.instance.vm.ToValue(instance.DispatchEvent(event))
	} else {
		panic(w.instance.vm.NewTypeError("Not an event"))
	}
}

func (w EventTargetWrapper) InitializePrototype(prototype *Object,
	vm *Runtime) {
	prototype.Set("addEventListener", w.AddEventListener)
	prototype.Set("dispatchEvent", w.DispatchEvent)
}
