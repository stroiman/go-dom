package generators

import "fmt"

// GeneratorStringer is a wrapper around Generator that implements the
// [fmt.Stringer] interface
type GeneratorStringer struct{ Generator }

// Returns a string representing the generated code.
//
// This is intended for test purposes only. You should used [jen.File] to render
// contents to a file or an [io.Writer] for "production" code.
func (g GeneratorStringer) String() string {
	// if s, ok := g.Generator.(fmt.Stringer); ok {
	// 	return s.String()
	// }
	return fmt.Sprintf("%#v", g.Generate())
}
