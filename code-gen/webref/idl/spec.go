package idl

import (
	"encoding/json"
	"io"
	"iter"
	"log/slog"
	"slices"
	"strings"

	"github.com/stroiman/go-dom/code-gen/webref"
)

type RetType struct {
	TypeName string
	Nullable bool
}

func NewRetTypeUndefined() RetType { return RetType{TypeName: "undefined", Nullable: false} }

type Spec struct {
	ParsedIdlFile
}

// LoadIdlParsed loads a files from the /curated/idlpased directory containing
// specifications of the interfaces.
func LoadIdlParsed(name string) (Spec, error) {
	file, err := webref.OpenIdlParsed(name)
	if err != nil {
		return Spec{}, err
	}
	defer file.Close()
	return ParseIdlJsonReader(file)
}

type TypeSpec struct {
	Spec *Spec
	Type Name
}

type MemberSpec struct{ NameMember }
type AttributeSpec struct{ NameMember }

func (t *TypeSpec) Members() []NameMember {
	return t.Type.Members
}

func (t *TypeSpec) Constructor() (res MemberSpec, found bool) {
	idx := slices.IndexFunc(t.Type.Members, func(n NameMember) bool {
		return n.Type == "constructor"
	})
	found = idx >= 0
	if found {
		res = MemberSpec{t.Type.Members[idx]}
	}
	return
}

func (t *TypeSpec) Inheritance() string {
	return t.Type.Inheritance
}

func (t *TypeSpec) InstanceMethods() iter.Seq[MemberSpec] {
	return func(yield func(v MemberSpec) bool) {
		for i, member := range t.Type.Members {
			if member.Special == "static" {
				continue
			}
			if member.Type == "operation" && member.Name != "" {
				// Empty name seems to indicate a named property getter. Not sure yet.
				firstIndex := slices.IndexFunc(t.Type.Members, func(m NameMember) bool {
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

func (t *TypeSpec) Attributes() iter.Seq[AttributeSpec] {
	return func(yield func(v AttributeSpec) bool) {
		for _, member := range t.Type.Members {
			if member.IsAttribute() {
				yield(AttributeSpec{member})
			}
		}
	}
}

func ParseIdlJsonReader(reader io.Reader) (Spec, error) {
	spec := Spec{}
	b, err := io.ReadAll(reader)
	if err == nil {
		err = json.Unmarshal(b, &spec)
	}
	return spec, err
}

func (s *Spec) GetType(name string) (TypeSpec, bool) {
	result, ok := s.IdlNames[name]
	return TypeSpec{s, result}, ok
}

func (s AttributeSpec) AttributeType() RetType {
	r, n := FindMemberAttributeType(s.NameMember)
	return RetType{r, n}
}

func (s MemberSpec) ReturnType() RetType {
	r, n := FindMemberReturnType(s.NameMember)
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
