package idl

import "iter"

// Interface represents an interface specification in the webref IDL files.
//
// For example, the following interface Animal is represented by an _interface_
//
//	[Exposed=Window]
//	interface Animal {
//		attribute DOMString name;
//	};
type Interface struct {
	// Don't rely on this, it only exists during a refactoring process
	InternalSpec Name
	Name         string
	Attributes   []Attribute
	// Includes represent interfaces included using the includes IDL statement.
	//
	// See also: https://webidl.spec.whatwg.org/#includes-statement
	Includes []Interface
}

// Represents an attribute on an IDL interface
type Attribute struct {
	// Don't rely on this, it only exists during a refactoring process
	InternalSpec NameMember
	Name         string
	Type         AttributeType
	Readonly     bool
}

type AttributeType struct {
	Name     string
	Nullable bool
}

// // NOTE: This will be removed in favour of a slice on the type
// func (i Interface) Attributes() iter.Seq[NameMember] {
// 	return i.InternalSpec.Attributes()
// }

// Attributes iterates and return all attributes from the IDO interface i. If
// included is true, this will also iterate attributes from interfaces that i
// includes.
func (i Interface) AllAttributes(included bool) iter.Seq[Attribute] {
	return func(yield func(Attribute) bool) {
		for _, a := range i.Attributes {
			if !yield(a) {
				return
			}
		}
		if included {
			for _, ii := range i.Includes {
				for _, a := range ii.Attributes {
					if !yield(a) {
						return
					}
				}
			}
		}
	}
}
