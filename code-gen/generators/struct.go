package generators

import "github.com/dave/jennifer/jen"

type Struct struct {
	Name    string
	Members []StructMember
}

type StructMember struct {
	Name Generator
	Type Generator
}

func NewStruct(name string) Struct { return Struct{Name: name} }

func (s *Struct) Field(name Generator, fieldType Generator) {
	s.Members = append(s.Members, StructMember{name, fieldType})
}

func (s *Struct) Embed(fieldType Generator) {
	s.Field(nil, fieldType)
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
	return jen.Type().Id(s.Name).Struct(fields...)
}
