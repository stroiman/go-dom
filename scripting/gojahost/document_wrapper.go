package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/dom"
)

type documentWrapper struct {
	baseInstanceWrapper[dom.Document]
}

type htmlDocumentWrapper struct {
	documentWrapper
}

func newDocumentWrapper(instance *GojaContext) wrapper {
	return documentWrapper{newBaseInstanceWrapper[dom.Document](instance)}
}

func newHTMLDocumentWrapper(instance *GojaContext) wrapper {
	return htmlDocumentWrapper{documentWrapper{newBaseInstanceWrapper[dom.Document](instance)}}
}

func (w documentWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	window := w.ctx.window
	newInstance := dom.NewDocument(window)
	w.storeInternal(newInstance, call.This)
	return nil
}

func (w htmlDocumentWrapper) initObject(o *g.Object) {
	o.DefineAccessorProperty(
		"location",
		w.ctx.vm.ToValue(w.location),
		nil,
		g.FLAG_TRUE,
		g.FLAG_TRUE,
	)
}
func (w documentWrapper) initObject(o *g.Object) {
	o.DefineAccessorProperty(
		"location",
		w.ctx.vm.ToValue(w.location),
		nil,
		g.FLAG_TRUE,
		g.FLAG_TRUE,
	)
}

func (w documentWrapper) initializePrototype(prototype *g.Object,
	vm *g.Runtime) {
	prototype.Set("createElement", w.createElement)
	prototype.Set("getElementById", w.getElementById)
}

func (w documentWrapper) location(c g.FunctionCall) g.Value {
	return w.ctx.vm.GlobalObject().Get("location")
}

func (w documentWrapper) createElement(c g.FunctionCall) g.Value {
	doc := w.getInstance(c)
	name := c.Argument(0)
	return w.toNode(doc.CreateElement(name.String()))
}

func (w documentWrapper) getElementById(c g.FunctionCall) g.Value {
	doc := w.getInstance(c)
	name := c.Argument(0)
	return w.toNode(doc.GetElementById(name.String()))
}
