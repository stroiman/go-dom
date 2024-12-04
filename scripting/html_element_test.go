package scripting_test

import (
	"errors"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/onsi/gomega/types"
	. "github.com/stroiman/go-dom/scripting"
)

type BeJSInstanceOfMatcher struct {
	class string
	ctx   *ScriptContext
}

func BeJSInstanceOf(
	expected string,
	ctx *ScriptContext,
) types.GomegaMatcher {
	return BeJSInstanceOfMatcher{expected, ctx}
}

var _ = Describe("V8 HTML Element classes", func() {
	It("Should have right element classes on elements", func() {
		ctx := NewTestContext(LoadHTML(sampleHTML)).ScriptContext
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
	result, err := m.ctx.RunScript(str + " instanceof " + m.class)
	if err == nil {
		success = result.Boolean()
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
