package idl

// Sanitize converts property names into valid Go identifiers, e.g., adding an
// underscore to reserved words like go and type.
func SanitizeName(name string) string {
	switch name {
	case "go":
		return "go_"
	case "type":
		return "type_"
	default:
		return name
	}
}
