package idl

import (
	"encoding/json"
	"io"
)

type IdlSpec struct {
	ParsedIdlFile
}

type IdlTypeSpec struct {
	Spec *IdlSpec
	Type IdlName
}

func (t *IdlTypeSpec) Members() []IdlNameMember {
	return t.Type.Members
}

func ParseIdlJsonReader(reader io.Reader) (IdlSpec, error) {
	decoder := json.NewDecoder(reader)
	spec := IdlSpec{}
	err := decoder.Decode(&spec.ParsedIdlFile)
	return spec, err
}

func (s *IdlSpec) GetType(name string) (IdlTypeSpec, bool) {
	result, ok := s.IdlNames[name]
	return IdlTypeSpec{s, result}, ok
}
