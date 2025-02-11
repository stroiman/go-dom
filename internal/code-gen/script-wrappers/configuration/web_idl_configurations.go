package configuration

// WebIdlConfigurations is a list of specifications for generating ES wrapper
// code. Each key in the map correspond to a specific IDL file
type WebIdlConfigurations map[string](*WebIdlConfiguration)

func NewWrapperGeneratorsSpec() WebIdlConfigurations {
	return make(WebIdlConfigurations)
}

// Module returns the configuration for a specific spec. A new configuration is
// created if it doesn't exist.
func (c WebIdlConfigurations) Module(spec string) *WebIdlConfiguration {
	if mod, ok := c[spec]; ok {
		return mod
	}
	mod := &WebIdlConfiguration{
		Name:       spec,
		Interfaces: make(map[string]*IdlInterfaceConfiguration),
	}
	c[spec] = mod
	return mod
}

func (s *WebIdlConfiguration) Type(typeName string) *IdlInterfaceConfiguration {
	if result, ok := s.Interfaces[typeName]; ok {
		return result
	}
	result := &IdlInterfaceConfiguration{
		DomSpec:  s,
		TypeName: typeName,
	}
	result.ensureMap()
	s.Interfaces[typeName] = result
	return result
}

func (spec *WebIdlConfiguration) SetMultipleFiles(value bool) { spec.MultipleFiles = value }

func (spec WebIdlConfiguration) UseMultipleFiles() bool {
	return spec.MultipleFiles == true
}
