package configuration

// type TypeCustomization []string

type ESMethodArgument struct {
	Required     bool
	HasDefault   bool
	DefaultValue string
	Ignored      bool
	encoder      string
	decoder      string
}

func (a *ESMethodArgument) SetRequired() *ESMethodArgument   { a.Required = true; return a }
func (a *ESMethodArgument) SetHasDefault() *ESMethodArgument { a.HasDefault = true; return a }

func (a *ESMethodArgument) HasDefaultValue(value string) {
	a.HasDefault = true
	a.DefaultValue = value
}

func (a *ESMethodArgument) Ignore() {
	a.Ignored = true
}

func (a *ESMethodArgument) SetEncoder(e string) *ESMethodArgument {
	a.encoder = e
	return a
}
func (a *ESMethodArgument) SetDecoder(d string) *ESMethodArgument {
	a.decoder = d
	return a
}
