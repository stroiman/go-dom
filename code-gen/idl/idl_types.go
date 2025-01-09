package idl

import (
	"encoding/json"
	"fmt"
	"strings"
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

type IdlType struct {
	Type     string    `json:"type"`
	ExtAttrs []ExtAttr `json:"extAttrs"`
	Generic  string    `json:"generic"`
	Nullable bool      `json:"nullable"`
	Union    bool      `json:"union"`
	IType    IdlTypes  `json:"idlType"`
}

func (t *IdlTypes) UnmarshalJSON(bytes []byte) error {
	err := json.Unmarshal(bytes, &t.Types)
	if err != nil {
		typ := new(IdlType)
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
	IdlType  IdlTypes  `json:"idlType"`
}

type IdlTypes struct {
	Types    []IdlType
	IdlType  *IdlType
	TypeName string
}

func (i IdlTypes) String() string {
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

type IdlNameMember struct {
	Stuff
	Arguments []ArgumentType `json:"arguments"`
	Special   string         `json:"special"`
	Readonly  bool           `json:"readOnly"`
	Href      string         `json:"href"`
}

type IdlName struct {
	Type    string          `json:"type"`
	Name    string          `json:"name"`
	Members []IdlNameMember `json:"members"`
	Partial bool            `json:"partial"`
	Href    string          `json:"href"`
}

type IdlExtendedName struct {
	Fragment string
	Type     string
	ExtAttrs []ExtAttr `json:"extAttrs"`
	Target   string
	Includes string
}

type IdlExtendedNames []IdlExtendedName

type IdlParsed struct {
	IdlNames         map[string]IdlName
	IdlExtendedNames map[string]IdlExtendedNames
}

type ParsedIdlFile struct {
	IdlParsed `json:"idlParsed"`
}

func FindIdlTypeValue(idl IdlTypes, expectedType string) (IdlType, bool) {
	types := idl.Types
	if len(types) == 0 && idl.IdlType != nil {
		types = []IdlType{*idl.IdlType}
	}
	for _, t := range types {
		if t.Type == expectedType {
			return t, true
		}
	}
	return IdlType{}, false
}

func FindIdlType(idl IdlTypes, expectedType string) (string, bool) {
	if t, ok := FindIdlTypeValue(idl, expectedType); ok {
		return t.IType.TypeName, t.Nullable
	}
	return "", false
}

func FindMemberReturnType(member IdlNameMember) (string, bool) {
	return FindIdlType(member.IdlType, "return-type")
}

func FindMemberAttributeType(member IdlNameMember) (string, bool) {
	return FindIdlType(member.IdlType, "attribute-type")
}

func (member IdlNameMember) IsAttribute() bool {
	if member.Type != "attribute" {
		return false
	}
	t, ok := FindIdlTypeValue(member.IdlType, "attribute-type")
	return ok && !strings.HasSuffix(t.IType.TypeName, "EventHandler")
}
