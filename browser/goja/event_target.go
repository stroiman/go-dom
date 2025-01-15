package goja_driver

import (
	"github.com/stroiman/go-dom/browser/dom"

	"github.com/dop251/goja"
	. "github.com/dop251/goja"
)

type eventTargetWrapper struct {
	baseInstanceWrapper[dom.EventTarget]
}

func newEventTargetWrapper(instance *GojaContext) wrapper {
	return eventTargetWrapper{newBaseInstanceWrapper[dom.EventTarget](instance)}
}

type gojaEventListener struct {
	instance *GojaContext
	v        goja.Value
	f        goja.Callable
}

func newGojaEventListener(r *GojaContext, v goja.Value) dom.EventHandler {
	f, ok := AssertFunction(v)
	if !ok {
		panic("TODO")
	}
	return &gojaEventListener{r, v, f}
}

func (h *gojaEventListener) HandleEvent(e dom.Event) error {
	customEvent := h.instance.globals["CustomEvent"]
	obj := h.instance.vm.CreateObject(customEvent.Prototype)
	customEvent.Wrapper.storeInternal(e, obj)
	_, err := h.f(obj, obj)
	return err
}

func (h *gojaEventListener) Equals(e dom.EventHandler) bool {
	if ge, ok := e.(*gojaEventListener); ok && ge.v.StrictEquals(h.v) {
		return true
	} else {
		return false
	}
}

func (w eventTargetWrapper) constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	newInstance := dom.NewEventTarget()
	w.storeInternal(newInstance, call.This)
	return nil
}

func (w eventTargetWrapper) getEventTarget(c FunctionCall) dom.EventTarget {
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

func (w eventTargetWrapper) addEventListener(c FunctionCall) Value {
	instance := w.getInstance(c)
	name := c.Argument(0).String()
	instance.AddEventListener(name, newGojaEventListener(w.instance, c.Argument(1)))
	return nil
}

func (w eventTargetWrapper) dispatchEvent(c FunctionCall) Value {
	instance := w.getInstance(c)
	internal := c.Argument(0).(*Object).GetSymbol(w.instance.wrappedGoObj).Export()
	if event, ok := internal.(dom.Event); ok {
		return w.instance.vm.ToValue(instance.DispatchEvent(event))
	} else {
		panic(w.instance.vm.NewTypeError("Not an event"))
	}
}

func (w eventTargetWrapper) initializePrototype(prototype *Object,
	vm *Runtime) {
	prototype.Set("addEventListener", w.addEventListener)
	prototype.Set("dispatchEvent", w.dispatchEvent)
}
