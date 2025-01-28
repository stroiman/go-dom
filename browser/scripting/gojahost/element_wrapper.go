package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/browser/dom"
)

type elementWrapper struct {
	baseInstanceWrapper[dom.Element]
}

func newElementWrapper(instance *GojaContext) wrapper {
	return elementWrapper{newBaseInstanceWrapper[dom.Element](instance)}
}
func (w elementWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.DefineAccessorProperty(
		"outerHTML",
		w.ctx.vm.ToValue(w.outerHTML),
		nil,
		g.FLAG_TRUE,
		g.FLAG_TRUE,
	)
}

func (w elementWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

func (w elementWrapper) outerHTML(call g.FunctionCall, r *g.Runtime) g.Value {
	return r.ToValue(w.getInstance(call).OuterHTML())
}
