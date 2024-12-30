package matchers

import (
	. "github.com/stroiman/go-dom/browser/dom"

	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/types"
)

func BeHTMLElement() OmegaMatcher { return HtmlElementMatcher{} }

type HtmlElementMatcher struct{}

func (m HtmlElementMatcher) Match(value any) (bool, error) {
	_, ok := value.(HTMLElement)
	return ok, nil
}

func (m HtmlElementMatcher) FailureMessage(actual interface{}) (message string) {
	return "Should be en HTMLElement"
}

func (m HtmlElementMatcher) NegatedFailureMessage(actual interface{}) (message string) {
	return "Should not be an HTMLElement"
}
