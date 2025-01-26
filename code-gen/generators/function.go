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
	Receiver FunctionArgument
	RtnTypes []Generator
	Body     Generator
}

func (f *FunctionDefinition) AddArgument(arg FunctionArgument) *FunctionDefinition {
	f.Args = append(f.Args, arg)
	return f
}

func (f *FunctionDefinition) WithReturnValues(values []Generator) *FunctionDefinition {
	f.RtnTypes = values
	return f
}

func (f *FunctionDefinition) WithReturnValue(value Generator) *FunctionDefinition {
	return f.WithReturnValues([]Generator{value})
}

func (f *FunctionDefinition) WithBody(body Generator) *FunctionDefinition {
	f.Body = body
	return f
}

func (f FunctionDefinition) Generate() *jen.Statement {
	var (
		args     []jen.Code = []jen.Code{}
		rtnTypes []jen.Code = []jen.Code{}
	)
	for _, arg := range f.Args {
		args = append(args, arg.Generate())
	}
	for _, t := range f.RtnTypes {
		rtnTypes = append(rtnTypes, t.Generate())
	}
	stmt := jen.Func()
	if f.Receiver.Name != nil {
		stmt.Params(f.Receiver.Generate())
	}
	if f.Name != "" {
		stmt.Id(f.Name)
	}
	stmt.Params(args...)
	if len(rtnTypes) > 0 {
		stmt.Params(rtnTypes...)
	}
	if f.Body != nil {
		stmt.Block(f.Body.Generate())
	}
	return stmt
}
