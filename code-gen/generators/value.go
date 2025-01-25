package generators

import "github.com/dave/jennifer/jen"

// Value represents a value in code, e.g. a local/global variable or constant,
// the return value of a function call, a field lookup on a struct, etc. The
// type os a wrapper on top of [Generator] to provide easy access to operations
// relevant for values.
//
// If the value is a struct values, you can get a Value representing field using
// [Value.Field].
//
// If the value is a function, you can call the function using [Value.Call]
//
// You can take a reference of a value (structs, strings, ints, etc) using
// [Value.Reference]
//
// The local package alias is automatically created based on the import
// specifications of the generated file.
type Value struct{ Generator }

// NewValue create a Value representing a value identified by identifier in the
// current scope, i.e., a global or local value
func NewValue(identifier string) Value { return Value{Id(identifier)} }

// NewValueOf create a Value expressed by an existing Generator instance.
func ValueOf(g Generator) Value { return Value{g} }

// NewValuePackage create a Value representing a global var or constant in an
// imported package. Identifier is the exported name, or identifier, and pkg is
// the fully qualified package name.
//
// The local package alias is automatically created based on the import
// specifications of the generated file.
func NewValuePackage(identifier string, pkg string) Value {
	return Value{Raw(jen.Qual(pkg, identifier))}
}

// Reference takes the references of a value using operator &
func (v Value) Reference() Generator {
	return Raw(jen.Op("&").Add(v.Generate()))
}

// Assign creates an assignment and initialization, i.e. using operator :=
func (v Value) Assign(expr Generator) Generator {
	return GeneratorFunc(func() *jen.Statement {
		return v.Generate().Op(":=").Add(expr.Generate())
	})
}

func (v Value) Field(name string) Value {
	return Value{Raw(v.Generate().Dot(name))}
}

func (v Value) Method(name string) Value {
	return Value{Raw(v.Generate().Dot(name))}
}

func (v Value) TypeParam(g Generator) Value {
	return Value{Raw(v.Generate().Index(g.Generate()))}
}

func (m Value) Call(args ...Generator) Value {
	return Value{Raw(m.Generate().CallFunc(func(g *jen.Group) {
		for _, arg := range args {
			g.Add(arg.Generate())
		}
	}))}
}
