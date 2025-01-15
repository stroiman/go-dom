package goja_driver

import (
	. "github.com/dop251/goja"
)

type baseInstanceWrapper[T any] struct {
	instance *GojaContext
}

func newBaseInstanceWrapper[T any](instance *GojaContext) baseInstanceWrapper[T] {
	return baseInstanceWrapper[T]{instance}
}

func (w baseInstanceWrapper[T]) storeInternal(value any, obj *Object) {
	obj.DefineDataPropertySymbol(
		w.instance.wrappedGoObj,
		w.instance.vm.ToValue(value),
		FLAG_FALSE,
		FLAG_FALSE,
		FLAG_FALSE,
	)
	// obj.SetSymbol(w.instance.wrappedGoObj, w.instance.vm.ToValue(value))
}

func (w baseInstanceWrapper[T]) getInstance(c FunctionCall) T {
	if c.This == nil {
		panic("No this pointer")
	}
	if result, ok := c.This.(*Object).GetSymbol(w.instance.wrappedGoObj).Export().(T); ok {
		return result
	} else {
		panic(w.instance.vm.NewTypeError("Not an event target"))
	}
}
