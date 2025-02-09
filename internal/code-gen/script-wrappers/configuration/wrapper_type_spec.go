package configuration

// ESClassWrapper contains information about how to generate ES wrapper code
// around a class in the web specification.
//
// All classes will be generated using a set of defaults. Data in this structure
// will allow deviating from the defaults.
type ESClassWrapper struct {
	DomSpec                   *WrapperGeneratorFileSpec
	TypeName                  string
	InnerTypeName             string
	WrapperTypeName           string
	Receiver                  string
	RunCustomCode             bool
	WrapperStruct             bool
	SkipPrototypeRegistration bool
	IncludeIncludes           bool
	Customization             map[string]*ESMethodWrapper
}

func (w *ESClassWrapper) ensureMap() {
	if w.Customization == nil {
		w.Customization = make(map[string]*ESMethodWrapper)
	}
}

func (w *ESClassWrapper) MarkMembersAsNotImplemented(names ...string) {
	w.ensureMap()
	for _, name := range names {
		w.Customization[name] = &ESMethodWrapper{NotImplemented: true}
	}
}
func (w *ESClassWrapper) MarkMembersAsIgnored(names ...string) {
	w.ensureMap()
	for _, name := range names {
		w.Customization[name] = &ESMethodWrapper{Ignored: true}
	}
}

func (w *ESClassWrapper) GetMethodCustomization(name string) (result ESMethodWrapper) {
	if val, ok := w.Customization[name]; ok {
		result = *val
	}
	return
}

func (w *ESClassWrapper) Method(name string) (result *ESMethodWrapper) {
	w.ensureMap()
	var ok bool
	if result, ok = w.Customization[name]; !ok {
		result = new(ESMethodWrapper)
		w.Customization[name] = result
	}
	return result
}

func (s *ESClassWrapper) CreateWrapper() {
	s.WrapperStruct = true
}
