package goja_driver

import . "github.com/dop251/goja"

type DocumentWrapper struct{}

func (w DocumentWrapper) Constructor(call ConstructorCall, r *Runtime) *Object {
	panic(r.NewTypeError("Illegal Constructor"))
}
