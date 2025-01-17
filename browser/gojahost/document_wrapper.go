package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
)

type documentWrapper struct {
	baseInstanceWrapper[dom.Document]
}

func newDocumentWrapper(instance *GojaContext) wrapper {
	return documentWrapper{newBaseInstanceWrapper[dom.Document](instance)}
}

func (w documentWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

func (w documentWrapper) initializePrototype(prototype *g.Object,
	vm *g.Runtime) {
	prototype.Set("createElement", w.createElement)
	prototype.Set("getElementById", w.getElementById)
}

func (w documentWrapper) createElement(c g.FunctionCall) g.Value {
	if c.This == nil {
		panic("No this pointer")
	}
	doc, ok := c.This.Export().(dom.Document)
	if !ok {
		panic("Not a document")
	}
	name := c.Argument(0)
	return w.ctx.vm.ToValue(doc.CreateElement(name.String()))
}

func (w documentWrapper) getElementById(c g.FunctionCall) g.Value {
	if c.This == nil {
		panic("No this pointer")
	}
	doc, ok := c.This.Export().(dom.Document)
	if !ok {
		panic("Not a document")
	}
	name := c.Argument(0)
	return w.toNode(doc.GetElementById(name.String()))
}
