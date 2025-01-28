package matchers

import (
	"errors"

	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/gcustom"
	"github.com/onsi/gomega/types"
	"github.com/gost-dom/browser/browser/html"
)

func HaveFormDataValue(key, expected string) types.GomegaMatcher {
	matcher := Equal(expected)
	var (
		countMismatch bool
		noOfMatches   int
	)

	return gcustom.MakeMatcher(
		func(actual *html.FormData) (bool, error) {
			if actual == nil {
				return false, errors.New("Formdata was nil")
			}
			values := actual.GetAll(key)
			noOfMatches = len(values)
			if 0 == noOfMatches {
				countMismatch = true
				return false, nil
			}
			if 1 < noOfMatches {
				countMismatch = true
				return false, nil
			}
			return matcher.Match(string(values[0]))
		}).WithTemplate("Expected:\n{{.FormattedActual}}\n{{.To}} have one value {{.Data.Key}}: {{.Data.Expected}}\n{{if .Data.CountMismatch }}Found {{.Data.NoOfMatches}}{{else}}{{.Data.Matcher.FailureMessage .Actual }}{{end}}", struct {
		Key           string
		Expected      string
		Matcher       types.GomegaMatcher
		CountMismatch *bool
		NoOfMatches   *int
	}{key, expected, matcher, &countMismatch, &noOfMatches})
}
