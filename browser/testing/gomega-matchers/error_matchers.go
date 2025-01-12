package matchers

import (
	"github.com/stroiman/go-dom/browser/dom"

	"github.com/onsi/gomega/gcustom"
	. "github.com/onsi/gomega/types"
)

func BeADOMError() GomegaMatcher {
	return gcustom.MakeMatcher(func(e error) (bool, error) {
		return dom.IsDOMError(e), nil
	})
}
