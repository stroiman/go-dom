package gojahost

import (
	"strings"

	g "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/scripting"
)

type baseInstanceWrapper[T any] struct {
	ctx *GojaContext
}

func newBaseInstanceWrapper[T any](instance *GojaContext) baseInstanceWrapper[T] {
	return baseInstanceWrapper[T]{instance}
}

func (w baseInstanceWrapper[T]) storeInternal(value any, obj *g.Object) {
	obj.DefineDataPropertySymbol(
		w.ctx.wrappedGoObj,
		w.ctx.vm.ToValue(value),
		g.FLAG_FALSE,
		g.FLAG_FALSE,
		g.FLAG_FALSE,
	)
	if e, ok := value.(dom.Entity); ok {
		w.ctx.cachedNodes[e.ObjectId()] = obj
	}
	// obj.SetSymbol(w.instance.wrappedGoObj, w.instance.vm.ToValue(value))
}

func getInstanceValue[T any](c *GojaContext, v g.Value) (T, bool) {
	res, ok := v.(*g.Object).GetSymbol(c.wrappedGoObj).Export().(T)
	return res, ok
}

func (w baseInstanceWrapper[T]) getInstance(c g.FunctionCall) T {
	if c.This == nil {
		panic("No this pointer")
	}
	if res, ok := getInstanceValue[T](w.ctx, c.This); ok {
		return res
	} else {
		panic(w.ctx.vm.NewTypeError("Not an entity"))
	}
}

func (w baseInstanceWrapper[T]) getCachedObject(e dom.Entity) g.Value {
	return w.ctx.cachedNodes[e.ObjectId()]
}

func (w baseInstanceWrapper[T]) decodeNode(v g.Value) dom.Node {
	if r, ok := getInstanceValue[dom.Node](w.ctx, v); ok {
		return r
	} else {
		panic("Bad node")
	}
}

func (w baseInstanceWrapper[T]) getPrototype(e dom.Entity) *g.Object {
	switch v := e.(type) {
	case dom.Element:
		className, found := scripting.HtmlElements[strings.ToLower(v.TagName())]
		if found {
			return w.ctx.globals[className].Prototype
		}
		return w.ctx.globals["Element"].Prototype
	case dom.Node:
		return w.ctx.globals["Node"].Prototype
	}
	panic("Prototype lookup not defined")
}

func (w baseInstanceWrapper[T]) toNode(e dom.Entity) g.Value {
	if o := w.getCachedObject(e); o != nil {
		return o
	}
	prototype := w.getPrototype(e)
	obj := w.ctx.vm.CreateObject(prototype)
	w.storeInternal(e, obj)
	return obj
}

func (w baseInstanceWrapper[T]) toBoolean(b bool) g.Value {
	return w.ctx.vm.ToValue(b)
}

func (w baseInstanceWrapper[T]) toDOMString(b string) g.Value {
	return w.ctx.vm.ToValue(b)
}

func (w baseInstanceWrapper[T]) toDocument(e dom.Entity) g.Value {
	return w.toNode(e)
}

func (w baseInstanceWrapper[T]) toUnsignedShort(i int) g.Value {
	return w.ctx.vm.ToValue(i)
}
