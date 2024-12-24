package main

type TypeCustomization []string

// ESClassWrapper contains information about a class in the web specifications,
// and how it is mapped to underlying go code
type ESClassWrapper struct {
	TypeName        string
	InnerTypeName   string
	WrapperTypeName string
	Receiver        string
	Customization   TypeCustomization
}
