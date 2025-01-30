package matchers

import (
	"github.com/onsi/gomega"
	"github.com/onsi/gomega/format"
	"github.com/onsi/gomega/gcustom"
	. "github.com/onsi/gomega/types"
	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
)

func FormatDocument(value any) (result string, ok bool) {
	var doc dom.Document
	if doc, ok = value.(dom.Document); ok {
		result = doc.DocumentElement().OuterHTML()
	}
	return
}

func init() {
	format.RegisterCustomFormatter(FormatDocument)
}

func mustQuerySelector(c dom.ElementContainer, pattern string) dom.Element {
	// The error is only because of a bad pattern, which is static in tests; so
	// panic. The test code is wrong, not the system.
	r, err := c.QuerySelector(pattern)
	if err != nil {
		panic(err)
	}
	return r
}

func mustQuerySelectorAll(c dom.ElementContainer, pattern string) []dom.Node {
	// The error is only because of a bad pattern, which is static in tests; so
	// panic. The test code is wrong, not the system.
	r, err := c.QuerySelectorAll(pattern)
	if err != nil {
		panic(err)
	}
	return r.All()
}

func HaveH1(expected string) GomegaMatcher {
	var data = struct {
		Expected       string
		Actual         string
		H1ElementCount int
		Matcher        GomegaMatcher
	}{
		Expected: expected,
		Matcher:  gomega.Equal(expected),
	}
	return gcustom.MakeMatcher(
		func(d html.HTMLDocument) (bool, error) {
			h1Elements := mustQuerySelectorAll(d, "h1")
			data.H1ElementCount = len(h1Elements)
			if data.H1ElementCount != 1 {
				return false, nil
			}
			data.Actual = h1Elements[0].TextContent()
			return data.Matcher.Match(data.Actual)
		}).WithTemplate(`Expected: {{.FormattedActual}}
{{.To}} have an <h1> with text content {{.Data.Expected}}
{{- if (gt .Data.H1ElementCount 1) }}
  Too many <h1> elements found: {{.Data.H1ElementCount}}{{end}}
{{- if (eq .Data.H1ElementCount 0) }}
  No <h1> element found{{end}}
{{- if (eq .Data.H1ElementCount 1) }}
{{.Data.Matcher.FailureMessage .Data.Actual}}{{end}}`, &data)
}
