package gojahost

import (
	"strings"

	g "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser/html"
	"github.com/stroiman/go-dom/browser/scripting"
)

type baseInstanceWrapper[T any] struct {
	ctx *GojaContext
}

func newBaseInstanceWrapper[T any](instance *GojaContext) baseInstanceWrapper[T] {
	return baseInstanceWrapper[T]{instance}
}

func (c *GojaContext) storeInternal(value any, obj *g.Object) {
	obj.DefineDataPropertySymbol(
		c.wrappedGoObj,
		c.vm.ToValue(value),
		g.FLAG_FALSE,
		g.FLAG_FALSE,
		g.FLAG_FALSE,
	)
	if e, ok := value.(dom.Entity); ok {
		c.cachedNodes[e.ObjectId()] = obj
	}
	// obj.SetSymbol(w.instance.wrappedGoObj, w.instance.vm.ToValue(value))
}

func (w baseInstanceWrapper[T]) storeInternal(value any, obj *g.Object) {
	w.ctx.storeInternal(value, obj)
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

func (c *GojaContext) getPrototype(e dom.Entity) function {
	switch v := e.(type) {
	case html.HTMLDocument:
		return c.globals["HTMLDocument"]
	case dom.Document:
		return c.globals["Document"]
	case dom.Element:
		className, found := scripting.HtmlElements[strings.ToLower(v.TagName())]
		if found {
			return c.globals[className]
		}
		return c.globals["Element"]
	case dom.Node:
		return c.globals["Node"]
	}
	panic("Prototype lookup not defined")
}

func (c *GojaContext) toNode(e dom.Entity) g.Value {
	if o, ok := c.cachedNodes[e.ObjectId()]; ok {
		return o
	}
	data := c.getPrototype(e)
	obj := c.vm.CreateObject(data.Prototype)
	c.storeInternal(e, obj)
	if initializer, ok := data.Wrapper.(instanceInitializer); ok {
		initializer.initObject(obj)
	}
	return obj
}

func (w baseInstanceWrapper[T]) toNode(e dom.Entity) g.Value {
	return w.ctx.toNode(e)
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
