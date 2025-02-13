package wrappers

import (
	g "github.com/gost-dom/generators"
)

// The ConstructorBuilder is the function that creates the ES constructor
// itself, i.e. starts with a new function template, installs prototypes on the
// template, etc.
type ConstructorBuilder struct {
	v8Iso
	Proto        v8PrototypeTemplate
	InstanceTmpl v8InstanceTemplate
	Wrapper      WrapperInstance
}

func NewConstructorBuilder() ConstructorBuilder {
	return ConstructorBuilder{
		v8Iso:        v8Iso{g.NewValue("iso")},
		Proto:        v8PrototypeTemplate{g.NewValue("prototypeTmpl")},
		InstanceTmpl: v8InstanceTemplate{g.NewValue("instanceTmpl")},
		Wrapper:      WrapperInstance{g.NewValue("wrapper")},
	}
}

func (builder ConstructorBuilder) NewFunctionTemplateOfWrappedMethod(name string) g.Generator {
	return builder.NewFunctionTemplate(builder.Wrapper.Method(name))
}

type PrototypeInstaller struct {
	v8Iso
	Proto v8PrototypeTemplate
	// InstanceTmpl v8InstanceTemplate
	Wrapper WrapperInstance
}

func (builder PrototypeInstaller) InstallFunctionHandlers(
	data ESConstructorData,
) JenGenerator {
	generators := make([]g.Generator, 0, len(data.Operations))
	for _, op := range data.Operations {
		if !op.MethodCustomization.Ignored {
			generators = append(generators,
				builder.Proto.Set(
					op.Name,
					builder.NewFunctionTemplate(builder.Wrapper.Field(op.CallbackMethodName())),
				),
			)
		}
	}
	return g.StatementList(generators...)
}

func (builder PrototypeInstaller) InstallAttributeHandlers(
	data ESConstructorData,
) g.Generator {
	length := len(data.Attributes)
	if length == 0 {
		return g.Noop
	}
	generators := make([]JenGenerator, 1, length+1)
	generators[0] = g.Line
	for op := range data.AttributesToInstall() {
		generators = append(generators, builder.InstallAttributeHandler(op))
	}
	return g.StatementList(generators...)
}

func (builder PrototypeInstaller) InstallAttributeHandler(
	op ESAttribute,
) g.Generator {
	wrapper := builder.Wrapper
	getter := op.Getter
	setter := op.Setter
	if getter == nil {
		return g.Noop
	}
	getterFt := builder.NewFunctionTemplate(wrapper.Field(getter.CallbackMethodName()))
	setterFt := g.Nil
	if setter != nil {
		setterFt = builder.NewFunctionTemplate(wrapper.Field(setter.CallbackMethodName()))
	}
	return builder.Proto.SetAccessorProperty(
		op.Name,
		g.WrapLine(getterFt),
		g.WrapLine(setterFt),
		g.WrapLine(v8None),
	)
}
