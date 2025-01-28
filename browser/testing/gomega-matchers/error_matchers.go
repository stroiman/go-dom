package matchers

import (
	"github.com/gost-dom/browser/browser/dom"

	"github.com/onsi/gomega/gcustom"
	. "github.com/onsi/gomega/types"
)

func BeADOMError() GomegaMatcher {
	return gcustom.MakeMatcher(func(e error) (bool, error) {
		return dom.IsDOMError(e), nil
	})
}
