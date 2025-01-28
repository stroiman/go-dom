package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/gost-dom/browser/browser/dom"
)

type nodeListWrapper struct {
	baseInstanceWrapper[dom.Node]
}

func newNodeListWrapper(instance *GojaContext) wrapper {
	return nodeListWrapper{newBaseInstanceWrapper[dom.Node](instance)}
}

func (w nodeListWrapper) constructor(c g.ConstructorCall, vm *g.Runtime) *g.Object {
	return nil
}

func (w nodeListWrapper) initializePrototype(prototype *g.Object, vm *g.Runtime) {
	prototype.Set("item", vm.ToValue(w.item))
	prototype.Set("keys", vm.ToValue(w.keys))
	prototype.Set("values", vm.ToValue(w.values))
	prototype.Set("forEach", vm.ToValue(w.forEach))
	prototype.Set("length", vm.ToValue(w.length))
}

func (w nodeListWrapper) item(c g.FunctionCall) g.Value {
	return nil
}

func (w nodeListWrapper) keys(c g.FunctionCall) g.Value {
	return nil
}

func (w nodeListWrapper) values(c g.FunctionCall) g.Value {
	return nil
}

func (w nodeListWrapper) forEach(c g.FunctionCall) g.Value {
	return nil
}

func (w nodeListWrapper) length(c g.FunctionCall) g.Value {
	return nil
}
