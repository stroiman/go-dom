package goja_driver

import "github.com/dop251/goja"

type WindowWrapper struct{}

func (w WindowWrapper) Constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}
