package gojahost

import (
	g "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/internal/entity"
)

type genericElementWrapper struct {
	baseInstanceWrapper[entity.Entity]
}

func newGenericElementWrapper(instance *GojaContext) wrapper {
	return genericElementWrapper{newBaseInstanceWrapper[entity.Entity](instance)}
}
func (w genericElementWrapper) constructor(call g.ConstructorCall, r *g.Runtime) *g.Object {
	panic(r.NewTypeError("Illegal Constructor"))
}
