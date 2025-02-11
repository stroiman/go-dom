package configuration

// IdlInterfaceConfiguration contains information about how to generate
// prototype objects for an interface in the IDL.
//
// All classes will be generated using a set of defaults. Data in this structure
// will allow deviating from the defaults.
type IdlInterfaceConfiguration struct {
	DomSpec                   *WebIdlConfiguration
	TypeName                  string
	RunCustomCode             bool
	WrapperStruct             bool
	SkipPrototypeRegistration bool
	IncludeIncludes           bool
	Customization             map[string]*ESMethodWrapper
}

func (w *IdlInterfaceConfiguration) ensureMap() {
	if w.Customization == nil {
		w.Customization = make(map[string]*ESMethodWrapper)
	}
}

func (w *IdlInterfaceConfiguration) MarkMembersAsNotImplemented(names ...string) {
	w.ensureMap()
	for _, name := range names {
		w.Customization[name] = &ESMethodWrapper{NotImplemented: true}
	}
}
func (w *IdlInterfaceConfiguration) MarkMembersAsIgnored(names ...string) {
	w.ensureMap()
	for _, name := range names {
		w.Customization[name] = &ESMethodWrapper{Ignored: true}
	}
}

func (w *IdlInterfaceConfiguration) GetMethodCustomization(name string) (result ESMethodWrapper) {
	if val, ok := w.Customization[name]; ok {
		result = *val
	}
	return
}

func (w *IdlInterfaceConfiguration) Method(name string) (result *ESMethodWrapper) {
	w.ensureMap()
	var ok bool
	if result, ok = w.Customization[name]; !ok {
		result = new(ESMethodWrapper)
		w.Customization[name] = result
	}
	return result
}

func (s *IdlInterfaceConfiguration) CreateWrapper() {
	s.WrapperStruct = true
}
