package goja_driver

import (
	. "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
)

type NodeWrapper struct {
	BaseInstanceWrapper[dom.Node]
}

func NewNodeWrapper(instance *GojaInstance) Wrapper {
	return NodeWrapper{NewBaseInstanceWrapper[dom.Node](instance)}
}

func (w NodeWrapper) Constructor(call ConstructorCall, r *Runtime) *Object {
	panic(r.NewTypeError("Illegal Constructor"))
}
