package goja_driver

import . "github.com/dop251/goja"

type NodeWrapper struct{}

func NewNodeWrapper(instance *GojaInstance) Wrapper { return NodeWrapper{} }

func (w NodeWrapper) Constructor(call ConstructorCall, r *Runtime) *Object {
	panic(r.NewTypeError("Illegal Constructor"))
}
