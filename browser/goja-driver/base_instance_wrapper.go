package goja_driver

import (
	. "github.com/dop251/goja"
)

type BaseInstanceWrapper[T any] struct {
	instance *GojaInstance
}

func NewBaseInstanceWrapper[T any](instance *GojaInstance) BaseInstanceWrapper[T] {
	return BaseInstanceWrapper[T]{instance}
}

func (w BaseInstanceWrapper[T]) StoreInternal(value any, obj *Object) *Object {
	result := w.instance.vm.ToValue(value).(*Object)
	result.SetPrototype(obj.Prototype())
	return result
}

func (w BaseInstanceWrapper[T]) GetInstance(c FunctionCall) T {
	if c.This == nil {
		panic("No this pointer")
	}
	var instance any
	if c.This == w.instance.vm.GlobalObject() {
		instance = w.instance.window
	} else {
		instance = c.This.(*Object).Export()
	}
	if result, ok := instance.(T); ok {
		return result
	} else {
		panic(w.instance.vm.NewTypeError("Not an event target"))
	}
}
