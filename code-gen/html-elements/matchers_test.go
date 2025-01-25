package htmlelements_test

import (
	"fmt"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	. "github.com/stroiman/go-dom/code-gen/generators"
)

func Render(expected string) types.GomegaMatcher {
	return WithTransform(
		// func(g Generator) string { return GeneratorStringer{Generator: g}.String() },
		func(g Generator) string { return fmt.Sprintf("%#v", g.Generate()) },
		Equal(expected),
	)
}
