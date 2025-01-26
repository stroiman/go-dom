package idl

import (
	"encoding/json"
	"fmt"
	"iter"
	"strings"
)

// IdlNameType represent the value of "type" of an exported name in an IDL
// files, and affects how the data is to be interpreted. Corresponds to the
// values in the json path:
//
//	idlParsed.idlNames[name].type
type IdlNameType string

const (
	IdlNameInterface IdlNameType = "interface"
)

type ValueType struct {
	Value ValueTypes `json:"value"`
}

type ValueTypes struct {
	Values    []ValueType
	Value     *ValueType
	ValueName string
}

type ExtAttr struct {
	Type string `json:"type"`
	Name string `json:"name"`
	Rhs  struct {
		Type  string     `json:"type"`
		Value ValueTypes `json:"value"`
	} `json:"rhs"`
}

type Type struct {
	Type     string    `json:"type"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	Generic  string    `json:"generic"`
	Nullable bool      `json:"nullable"`
	Union    bool      `json:"union"`
	IType    Types     `json:"idlType"`
}

func (t *Types) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &t.Types)
	if err != nil {
		typ := new(Type)
		err = json.Unmarshal(bytes, &typ)
		if err == nil {
			t.IdlType = typ
		}
	}
	if err != nil {
		err = json.Unmarshal(bytes, &t.TypeName)
	}
	return err

}

type Stuff struct {
	Type     string    `json:"type"`
	Name     string    `json:"name"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	IdlType  Types     `json:"idlType"`
}

type Types struct {
	Types    []Type
	IdlType  *Type
	TypeName string
}

func (i Types) String() string {
	if len(i.Types) > 0 {
		return fmt.Sprintf("%v", i.Types)
	}
	if i.IdlType != nil {
		return fmt.Sprintf("%v", *i.IdlType)
	}
	return i.TypeName
}

func (t *ValueTypes) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &t.Values)
	if err != nil {
		val := new(ValueType)
		err = json.Unmarshal(bytes, val)
		if err == nil {
			t.Value = val
		}
	}
	if err != nil {
		err = json.Unmarshal(bytes, &t.ValueName)
	}
	return err
}

type ArgumentType struct {
	Stuff
	Default  any  `json:"default"`
	Optional bool `json:"optional"`
	Variadic bool `json:"variadic"`
}

type NameMember struct {
	Stuff
	Arguments []ArgumentType `json:"arguments"`
	Special   string         `json:"special"`
	Readonly  bool           `json:"readOnly"`
	Href      string         `json:"href"`
}

type Name struct {
	Type        string       `json:"type"`
	Name        string       `json:"name"`
	Members     []NameMember `json:"members"`
	Partial     bool         `json:"partial"`
	Href        string       `json:"href"`
	Inheritance string       `json:"Inheritance"`
}

func (n Name) Attributes() iter.Seq[NameMember] {
	return func(yield func(NameMember) bool) {
		for _, m := range n.Members {
			if m.Type == "attribute" {
				if !yield(m) {
					return
				}
			}
		}
	}
}

type ExtendedName struct {
	Fragment string
	Type     string
	ExtAttrs []ExtAttr `json:"extAttrs"`
	Target   string
	Includes string
}

type ExtendedNames []ExtendedName

func (nn ExtendedNames) includes() []string {
	res := make([]string, 0)
	for _, n := range nn {
		if n.Type == "includes" {
			res = append(res, n.Includes)
		}
	}
	return res
}

type Parsed struct {
	IdlNames         map[string]Name
	IdlExtendedNames map[string]ExtendedNames
}

type ParsedIdlFile struct {
	Parsed `json:"idlParsed"`
}

func FindIdlTypeValue(idl Types, expectedType string) (Type, bool) {
	types := idl.Types
	if len(types) == 0 && idl.IdlType != nil {
		types = []Type{*idl.IdlType}
	}
	for _, t := range types {
		if t.Type == expectedType {
			return t, true
		}
	}
	return Type{}, false
}

func FindIdlType(idl Types, expectedType string) (string, bool) {
	if t, ok := FindIdlTypeValue(idl, expectedType); ok {
		return t.IType.TypeName, t.Nullable
	}
	return "", false
}

func FindMemberReturnType(member NameMember) (string, bool) {
	return FindIdlType(member.IdlType, "return-type")
}

func FindMemberAttributeType(member NameMember) (string, bool) {
	return FindIdlType(member.IdlType, "attribute-type")
}

func (member NameMember) IsAttribute() bool {
	if member.Type != "attribute" {
		return false
	}
	t, ok := FindIdlTypeValue(member.IdlType, "attribute-type")
	return ok && !strings.HasSuffix(t.IType.TypeName, "EventHandler")
}
