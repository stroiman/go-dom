package main

type TypeCustomization []string

type ESMethodWrapper struct {
	NotImplemented bool
}

// ESClassWrapper contains information about a class in the web specifications,
// and how it is mapped to underlying go code
type ESClassWrapper struct {
	TypeName        string
	InnerTypeName   string
	WrapperTypeName string
	Receiver        string
	Customization   map[string]ESMethodWrapper
}

func (w *ESClassWrapper) ensureMap() {
	if w.Customization == nil {
		w.Customization = make(map[string]ESMethodWrapper)
	}
}

func (w *ESClassWrapper) MarkMembersAsNotImplemented(names ...string) {
	w.ensureMap()
	for _, name := range names {
		w.Customization[name] = ESMethodWrapper{NotImplemented: true}
	}
}
