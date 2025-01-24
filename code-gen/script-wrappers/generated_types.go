// The types here are abstractions over concrete types in generated code,
// helping function lookup code being easier to read.

package wrappers

import (
	g "github.com/stroiman/go-dom/code-gen/generators"
)

type v8ArgInfo g.Value

func (info v8ArgInfo) GetV8Context() g.Generator { return g.Value(info).Method("Context").Call() }

type WrapperInstance struct{ g.Value }

func (i WrapperInstance) GetScriptHost() g.Value { return i.Field("host") }

func (i WrapperInstance) MustGetContext(info g.Generator) g.Generator {
	return i.Method("mustGetContext").Call(info)
}

type v8PrototypeTemplate struct{ g.Value }

func (proto v8PrototypeTemplate) Set(name string, handler g.Generator) g.Generator {
	return proto.Value.Method("Set").Call(g.Lit(name), handler)
}

type v8InstanceTemplate struct{ g.Value }

func (tmpl v8InstanceTemplate) SetInternalFieldCount(val int) g.Generator {
	return tmpl.Method("SetInternalFieldCount").Call(g.Lit(val))
}

type v8FunctionTemplate struct{ g.Value }

func (ft v8FunctionTemplate) GetPrototypeTemplate() g.Generator {
	return ft.Method("PrototypeTemplate").Call()
}

func (ft v8FunctionTemplate) GetInstanceTemplate() v8InstanceTemplate {
	return v8InstanceTemplate{ft.Method("InstanceTemplate").Call()}
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

// Provides helpers for functions that needs an iso as the first argument
type v8Iso struct{ g.Value }

func (iso v8Iso) NewFunctionTemplate(cb g.Generator) g.Generator {
	return NewV8FunctionTemplate{iso.Value, cb}
}
