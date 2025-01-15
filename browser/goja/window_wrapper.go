package goja

import (
	"github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/html"
)

type windowWrapper struct {
	baseInstanceWrapper[html.Window]
}

func (w windowWrapper) constructor(call goja.ConstructorCall, r *goja.Runtime) *goja.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

func newWindowWrapper(instance *GojaContext) wrapper {
	return windowWrapper{newBaseInstanceWrapper[html.Window](instance)}
}
