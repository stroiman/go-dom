package goja

import (
	. "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
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
	if e, ok := value.(dom.Entity); ok {
		w.instance.cachedNodes[e.ObjectId()] = obj
	}
	// obj.SetSymbol(w.instance.wrappedGoObj, w.instance.vm.ToValue(value))
}

func getInstanceValue[T any](c *GojaContext, v Value) (T, bool) {
	res, ok := v.(*Object).GetSymbol(c.wrappedGoObj).Export().(T)
	return res, ok
}

func (w baseInstanceWrapper[T]) getInstance(c FunctionCall) T {
	if c.This == nil {
		panic("No this pointer")
	}
	if res, ok := getInstanceValue[T](w.instance, c.This); ok {
		return res
	} else {
		panic(w.instance.vm.NewTypeError("Not an entity"))
	}
}

func (w baseInstanceWrapper[T]) getCachedObject(e dom.Entity) Value {
	return w.instance.cachedNodes[e.ObjectId()]
}

func (w baseInstanceWrapper[T]) decodeNode(v Value) dom.Node {
	if r, ok := getInstanceValue[dom.Node](w.instance, v); ok {
		return r
	} else {
		panic("Bad node")
	}
}

func (w baseInstanceWrapper[T]) getPrototype(e dom.Entity) *Object {
	switch e.(type) {
	case dom.Node:
		return w.instance.globals["Node"].Prototype
	}
	panic("Prototype lookup not defined")
}

func (w baseInstanceWrapper[T]) toNode(e dom.Entity) Value {
	if o := w.getCachedObject(e); o != nil {
		return o
	}
	prototype := w.getPrototype(e)
	obj := w.instance.vm.CreateObject(prototype)
	w.storeInternal(e, obj)
	return obj
}

func (w baseInstanceWrapper[T]) toBoolean(b bool) Value {
	return w.instance.vm.ToValue(b)
}

func (w baseInstanceWrapper[T]) toDOMString(b string) Value {
	return w.instance.vm.ToValue(b)
}

func (w baseInstanceWrapper[T]) toDocument(e dom.Entity) Value {
	return w.toNode(e)
}
