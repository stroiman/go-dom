package configuration

import (
	"cmp"
	"slices"
)

type WebIdlConfiguration struct {
	// Name is the name of the specification, corresponds to the file name.
	Name string
	// MultipleFiles indicate if one, or multiple files should be generated for
	// a specification
	MultipleFiles bool
	// Interfaces defines the names of the specified interfaces for which to
	// generate specifications
	Interfaces map[string]*IdlInterfaceConfiguration
}

func (spec WebIdlConfiguration) GetTypesSorted() []*IdlInterfaceConfiguration {
	types := make([]*IdlInterfaceConfiguration, len(spec.Interfaces))
	idx := 0
	for _, t := range spec.Interfaces {
		types[idx] = t
		idx++
	}
	slices.SortFunc(types, func(x, y *IdlInterfaceConfiguration) int {
		return cmp.Compare(x.TypeName, y.TypeName)
	})
	return types
}
