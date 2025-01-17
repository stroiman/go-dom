package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
)

type genericElementWrapper struct {
	baseInstanceWrapper[dom.Element]
}

func newGenericElementWrapper(instance *GojaContext) wrapper {
	return genericElementWrapper{newBaseInstanceWrapper[dom.Element](instance)}
}
func (w genericElementWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}
