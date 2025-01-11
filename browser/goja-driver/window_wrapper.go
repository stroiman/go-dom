package goja_driver

import (
	"github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/html"
)

type WindowWrapper struct {
	BaseInstanceWrapper[html.Window]
}

func (w WindowWrapper) Constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

func NewWindowWrapper(instance *GojaInstance) Wrapper {
	return WindowWrapper{NewBaseInstanceWrapper[html.Window](instance)}
}
