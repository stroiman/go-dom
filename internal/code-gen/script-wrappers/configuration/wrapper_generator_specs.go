package configuration

import (
	"cmp"
	"slices"

	"github.com/gost-dom/generators"
)

type WrapperTypeSpec = *ESClassWrapper

type WrapperGeneratorFileSpec struct {
	Name          string
	MultipleFiles bool
	Types         map[string]WrapperTypeSpec
}

func NewWrapperGeneratorsSpec() WrapperGeneratorsSpec {
	return make(WrapperGeneratorsSpec)
}

// WrapperGeneratorsSpec is a list of specifications for generating ES wrapper
// code. Each key in the map correspond to a specific IDL file
type WrapperGeneratorsSpec map[string](*WrapperGeneratorFileSpec)

func (spec WrapperGeneratorFileSpec) GetTypesSorted() []WrapperTypeSpec {
	types := make([]WrapperTypeSpec, len(spec.Types))
	idx := 0
	for _, t := range spec.Types {
		types[idx] = t
		idx++
	}
	slices.SortFunc(types, func(x, y WrapperTypeSpec) int {
		return cmp.Compare(x.TypeName, y.TypeName)
	})
	return types
}

func (spec WrapperGeneratorFileSpec) UseMultipleFiles() bool {
	return spec.MultipleFiles == true
}

func (spec *WrapperGeneratorFileSpec) SetMultipleFiles(value bool) { spec.MultipleFiles = value }

func (g WrapperGeneratorsSpec) Module(spec string) *WrapperGeneratorFileSpec {
	if mod, ok := g[spec]; ok {
		return mod
	}
	mod := &WrapperGeneratorFileSpec{
		Name:  spec,
		Types: make(map[string]WrapperTypeSpec),
	}
	g[spec] = mod
	return mod
}

func (s *WrapperGeneratorFileSpec) Type(typeName string) WrapperTypeSpec {
	if result, ok := s.Types[typeName]; ok {
		return result
	}
	result := &ESClassWrapper{
		DomSpec:  s,
		TypeName: typeName,
		Receiver: generators.DefaultReceiverName(typeName),
	}
	result.ensureMap()
	s.Types[typeName] = result
	return result
}
