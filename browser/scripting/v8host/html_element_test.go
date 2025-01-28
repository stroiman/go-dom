package v8host_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	"github.com/gost-dom/browser/browser/html"
)

type BeJSInstanceOfMatcher struct {
	class string
	ctx   html.ScriptContext
}

func BeJSInstanceOf(
	expected string,
	ctx html.ScriptContext,
) types.GomegaMatcher {
	return BeJSInstanceOfMatcher{expected, ctx}
}

var _ = Describe("V8 HTML Element classes", func() {
	It("Should have right element classes on elements", func() {
		ctx := NewTestContext(LoadHTML(sampleHTML))
		Expect("document.getElementById('a')").To(BeJSInstanceOf("HTMLAnchorElement", ctx))
		Expect("document.getElementById('p')").To(BeJSInstanceOf("HTMLParagraphElement", ctx))
		Expect("document.getElementById('div')").To(BeJSInstanceOf("HTMLDivElement", ctx))
	})
})

func (m BeJSInstanceOfMatcher) Match(actual interface{}) (success bool, err error) {
	str, ok := actual.(string)
	if !ok {
		return false, errors.New("Actual must be a string")
	}
	v, err := m.ctx.Eval(str + " instanceof " + m.class)
	success, ok = v.(bool)
	if !ok {
		panic("Should have received a bool")
	}
	return
}

func (m BeJSInstanceOfMatcher) FailureMessage(actual any) string {
	return "Expected an instance of " + m.class
}
func (m BeJSInstanceOfMatcher) NegatedFailureMessage(actual any) string {
	return "Expected to not be an instance of " + m.class
}

var sampleHTML = `<body>
	<a ref="#foo" id="a">Link</a>
	<p id="p">Paragraph</p>
	<div id="div">Div</div>
</body>`
