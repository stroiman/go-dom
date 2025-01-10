package goja_driver

import (
	. "github.com/dop251/goja"
	"github.com/stroiman/go-dom/browser/dom"
)

type DocumentWrapper struct {
}

func (w DocumentWrapper) Constructor(call ConstructorCall, r *Runtime) *Object {
	panic(r.NewTypeError("Illegal Constructor"))
}

func (w DocumentWrapper) InitializePrototype(prototype *Object,
	vm *Runtime) {
	createElement := vm.ToValue(func(c FunctionCall) Value {
		if c.This == nil {
			panic("No this pointer")
		}
		doc, ok := c.This.Export().(dom.Document)
		if !ok {
			panic("Not a document")
		}
		name := c.Argument(0)
		return vm.ToValue(doc.CreateElement(name.String()))
	})

	prototype.Set("createElement", createElement)
}
