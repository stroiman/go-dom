package idl

import (
	"encoding/json"
	"io"
	"iter"
	"log/slog"
	"slices"
	"strings"
)

type RetType struct {
	TypeName string
	Nullable bool
}

func NewRetTypeUndefined() RetType { return RetType{TypeName: "undefined", Nullable: false} }

type IdlSpec struct {
	ParsedIdlFile
}

type IdlTypeSpec struct {
	Spec *IdlSpec
	Type IdlName
}

type MemberSpec struct{ IdlNameMember }
type AttributeSpec struct{ IdlNameMember }

func (t *IdlTypeSpec) Members() []IdlNameMember {
	return t.Type.Members
}

func (t *IdlTypeSpec) Constructor() (res MemberSpec, found bool) {
	idx := slices.IndexFunc(t.Type.Members, func(n IdlNameMember) bool {
		return n.Type == "constructor"
	})
	found = idx >= 0
	if found {
		res = MemberSpec{t.Type.Members[idx]}
	}
	return
}

func (t *IdlTypeSpec) InstanceMethods() iter.Seq[MemberSpec] {
	return func(yield func(v MemberSpec) bool) {
		for i, member := range t.Type.Members {
			if member.Special == "static" {
				continue
			}
			if member.Type == "operation" && member.Name != "" {
				// Empty name seems to indicate a named property getter. Not sure yet.
				firstIndex := slices.IndexFunc(t.Type.Members, func(m IdlNameMember) bool {
					return m.Name == member.Name
				})
				if firstIndex < i {
					slog.Warn("Function overloads", "Name", member.Name)
					continue
				} else {
					yield(MemberSpec{member})
				}
			}
		}
	}
}

func (t *IdlTypeSpec) Attributes() iter.Seq[AttributeSpec] {
	return func(yield func(v AttributeSpec) bool) {
		for _, member := range t.Type.Members {
			if member.IsAttribute() {
				yield(AttributeSpec{member})
			}
		}
	}
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

func (s AttributeSpec) AttributeType() RetType {
	r, n := FindMemberAttributeType(s.IdlNameMember)
	return RetType{r, n}
}

func (s MemberSpec) ReturnType() RetType {
	r, n := FindMemberReturnType(s.IdlNameMember)
	return RetType{r, n}
}

func (t RetType) IsUndefined() bool { return t.TypeName == "undefined" }
func (t RetType) IsDefined() bool   { return !t.IsUndefined() }

func (t RetType) IsNode() bool {
	loweredName := strings.ToLower(t.TypeName)
	switch loweredName {
	case "node":
		return true
	case "document":
		return true
	case "documentfragment":
		return true
	}
	if strings.HasSuffix(loweredName, "element") {
		return true
	}
	return false
}
