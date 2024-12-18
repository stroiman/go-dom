package generators

import "github.com/dave/jennifer/jen"

type FunctionArgument struct {
	Name Generator
	Type Generator
}

func (f FunctionArgument) Generate() st {
	return f.Name.Generate().Add(f.Type.Generate())
}

type FunctionArgumentList []FunctionArgument

func Arg(name Generator, t Generator) FunctionArgumentList {
	return FunctionArgumentList([]FunctionArgument{{name, t}})
}

func (a FunctionArgumentList) Arg(name Generator, t Generator) FunctionArgumentList {
	return append(a, Arg(name, t)...)
}

// FunctionDefinition defines a named function in global scope.
type FunctionDefinition struct {
	Name     string
	Args     FunctionArgumentList
	RtnTypes []Generator
	Body     Generator
}

func (f FunctionDefinition) Generate() *jen.Statement {
	var (
		args     []jen.Code
		rtnTypes []jen.Code
	)
	for _, arg := range f.Args {
		args = append(args, arg.Generate())
	}
	for _, t := range f.RtnTypes {
		rtnTypes = append(rtnTypes, t.Generate())
	}
	stmt := jen.Func()
	if f.Name != "" {
		stmt.Id(f.Name)
	}
	stmt.Params(args...)
	if len(rtnTypes) > 0 {
		stmt.Params(rtnTypes...)
	}
	return stmt.Block(f.Body.Generate())
}
