package main

import (
	"encoding/json"
	"fmt"
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

type IdlParsed struct {
	IdlNames map[string]IdlName
}

type ParsedIdlFile struct {
	IdlParsed `json:"idlParsed"`
}

func FindIdlType(idl IdlTypes) (string, bool) {
	types := idl.Types
	if len(types) == 0 && idl.IdlType != nil {
		types = []IdlType{*idl.IdlType}
	}
	for _, t := range types {
		if t.Type == "return-type" {
			return t.IType.TypeName, t.Nullable
		}
	}
	return "", false

}
