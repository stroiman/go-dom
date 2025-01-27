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

// Spec represents the information stored in Web IDL files.
type Spec struct {
	// ParsedIdlFile is a direct JSON deserialisation of the data.
	//
	// Note: This was the first implementation and is most complete in terms of
	// available data, but has a lower level of abstraction. Use the other
	// properties, When data is available on those, e.g., find an interface from
	// the Interfaces map.
	//
	// This property will eventually be removed
	ParsedIdlFile
	Interfaces map[string]Interface
}

func (s *Spec) createInterface(n Name) Interface {
	includedNames := s.IdlExtendedNames[n.Name].includes()

	jsonAttributes := slices.Collect(n.Attributes())

	intf := Interface{
		InternalSpec: n,
		Name:         n.Name,
		Includes:     make([]Interface, len(includedNames)),
		Attributes:   make([]Attribute, len(jsonAttributes)),
	}
	for i, n := range includedNames {
		intf.Includes[i] = s.createInterface(s.IdlNames[n])
	}
	for i, a := range jsonAttributes {
		name, nullable := FindMemberAttributeType(a)
		intf.Attributes[i] = Attribute{
			InternalSpec: a,
			Name:         a.Name,
			Type:         AttributeType{Name: name, Nullable: nullable},
			Readonly:     a.Readonly,
		}
	}
	return intf
}

// initialize fills out the high-level representations from the low level parsed
// JSON data.
func (s *Spec) initialize() {
	s.Interfaces = make(map[string]Interface)
	for name, spec := range s.IdlNames {
		s.Interfaces[name] = s.createInterface(spec)
	}
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
	Spec         *Spec
	IdlInterface Interface
}

type MemberSpec struct{ NameMember }
type AttributeSpec struct{ NameMember }

func (t *TypeSpec) Members() []NameMember {
	return t.IdlInterface.InternalSpec.Members
}

func (t *TypeSpec) Constructor() (res MemberSpec, found bool) {
	idx := slices.IndexFunc(t.IdlInterface.InternalSpec.Members, func(n NameMember) bool {
		return n.Type == "constructor"
	})
	found = idx >= 0
	if found {
		res = MemberSpec{t.IdlInterface.InternalSpec.Members[idx]}
	}
	return
}

func (t *TypeSpec) Inheritance() string {
	return t.IdlInterface.InternalSpec.Inheritance
}

func (t *TypeSpec) InstanceMethods() iter.Seq[MemberSpec] {
	return func(yield func(v MemberSpec) bool) {
		for i, member := range t.IdlInterface.InternalSpec.Members {
			if member.Special == "static" {
				continue
			}
			if member.Type == "operation" && member.Name != "" {
				// Empty name seems to indicate a named property getter. Not sure yet.
				firstIndex := slices.IndexFunc(
					t.IdlInterface.InternalSpec.Members,
					func(m NameMember) bool {
						return m.Name == member.Name
					},
				)
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
		for _, member := range t.IdlInterface.InternalSpec.Members {
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
	spec.initialize()
	return spec, err
}

func (s *Spec) GetType(name string) (TypeSpec, bool) {
	result, ok := s.Interfaces[name]
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
