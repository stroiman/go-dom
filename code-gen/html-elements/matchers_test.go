package htmlelements_test

import (
	"fmt"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	. "github.com/stroiman/go-dom/code-gen/generators"
)

func HaveRendered(expected interface{}) types.GomegaMatcher {
	matcher, ok := expected.(types.GomegaMatcher)
	if !ok {
		return HaveRendered(Equal(expected))
	}
	return WithTransform(
		func(g Generator) string { return fmt.Sprintf("%#v", g.Generate()) },
		matcher,
	)
}
