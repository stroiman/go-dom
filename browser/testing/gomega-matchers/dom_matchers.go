package matchers

import (
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	. "github.com/onsi/gomega/types"
	"github.com/stroiman/go-dom/browser/dom"
)

func HaveAttribute(name string, expected interface{}) GomegaMatcher {
	var (
		found     bool
		actual    string
		matcher   GomegaMatcher
		isMatcher bool
	)
	if matcher, isMatcher = expected.(GomegaMatcher); !isMatcher {
		return HaveAttribute(name, gomega.Equal(expected))
	}
	return gcustom.MakeMatcher(func(e dom.Element) (bool, error) {
		if actual, found = e.GetAttribute(name); !found {
			return false, nil
		}
		return matcher.Match(actual)
	}).WithTemplate("Expected:\n{{.FormattedActual}}\n{{.To}} have have attribute '{{.Data.Attribute}}' {{if .Data.Found}}{{.Data.Matcher.FailureMessage .Data.Actual}}{{else}}, but it wasn't found{{end}}", struct {
		Attribute string
		Matcher   GomegaMatcher
		Found     *bool
		Actual    *string
	}{name, matcher, &found, &actual})
}
