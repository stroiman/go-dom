package htmlelements_test

import (
	"fmt"

	. "github.com/gost-dom/generators"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
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
