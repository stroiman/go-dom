package generators

import "github.com/dave/jennifer/jen"

type Struct struct {
	Name            Type
	Members         []StructMember
	DefaultReceiver string
}

type StructMember struct {
	Name Generator
	Type Generator
}

func NewStruct(name Generator) Struct {
	n, ok := name.(Type)
	if !ok {
		n = Type{name}
	}
	return Struct{Name: n}
}

func (s *Struct) Field(name Generator, fieldType Generator) {
	s.Members = append(s.Members, StructMember{name, fieldType})
}

func (s *Struct) Embed(fieldType Generator) {
	s.Field(nil, fieldType)
}

// SetDefaultReceiver sets the name of the receiver when generating methods
// using [Struct.CreateMethod].
func (s *Struct) SetDefaultReceiver(name string) {
	s.DefaultReceiver = name
}

func (s *Struct) MethodName(name string) FunctionDefinition {
	return FunctionDefinition{
		Name: name,
		Receiver: FunctionArgument{
			Name: Id(s.DefaultReceiver),
			Type: s.Name,
		},
	}
}

func (s *Struct) PointerMethodName(name string) *FunctionDefinition {
	return &FunctionDefinition{
		Name: name,
		Receiver: FunctionArgument{
			Name: Id(s.DefaultReceiver),
			Type: s.Name.Pointer(),
		},
	}
}

func (s Struct) Generate() *jen.Statement {
	fields := make([]jen.Code, 0, len(s.Members))
	for _, m := range s.Members {
		if m.Name == nil {
			fields = append(fields, m.Type.Generate())
		} else {
			fields = append(fields, m.Name.Generate().Add(m.Type.Generate()))
		}
	}
	return jen.Type().Add(s.Name.Generate()).Struct(fields...)
}

func InstantiateStruct(t Generator, values ...Generator) Generator {
	return Raw(t.Generate().Values(ToJenCodes(values)...))
}
