// The types here are abstractions over concrete types in generated code,
// helping function lookup code being easier to read.

package main

import (
	g "github.com/stroiman/go-dom/code-gen/generators"
)

type v8ArgInfo g.Value

func (info v8ArgInfo) GetV8Context() g.Generator { return g.Value(info).Method("Context").Call() }

type WrapperInstance struct{ g.Value }

func (i WrapperInstance) GetScriptHost() g.Value { return i.Field("host") }

type v8PrototypeTemplate struct{ g.Value }

func (proto v8PrototypeTemplate) Set(name string, handler g.Generator) g.Generator {
	return proto.Value.Method("Set").Call(g.Lit(name), handler)
}

func (proto v8PrototypeTemplate) SetAccessorProperty(
	name string,
	arguments ...g.Generator,
) g.Generator {
	args := make([]g.Generator, len(arguments)+1)
	args[0] = g.Lit(name)
	copy(args[1:], arguments)
	return proto.Method("SetAccessorProperty").Call(args...)
}
