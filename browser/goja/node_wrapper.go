package goja_driver

import (
	. "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
)

type nodeWrapper struct {
	baseInstanceWrapper[dom.Node]
}

func NewNodeWrapper(instance *GojaContext) wrapper {
	return nodeWrapper{newBaseInstanceWrapper[dom.Node](instance)}
}

func (w nodeWrapper) constructor(call ConstructorCall, r *Runtime) *Object {
	panic(r.NewTypeError("Illegal Constructor"))
}
