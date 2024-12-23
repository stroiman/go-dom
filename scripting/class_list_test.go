package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 Element.classList", func() {
	It("Should support .add", func() {
		ctx := NewTestContext(LoadHTML("<div id='target' class='c1'></div>"))
		ctx.MustRunTestScript(`
			const list = document.getElementById('target').classList;
			list.add('c2')
		`)
		element := ctx.Window().Document().GetElementById("target")
		Expect(element.GetAttribute("class")).To(Equal("c1 c2"))
	})
})
