package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 NamedNodeMap", func() {
	ctx := InitializeContextWithEmptyHtml()

	It("Should inherit directly from Object", func() {
		Expect(
			ctx.RunTestScript("Object.getPrototypeOf(NamedNodeMap.prototype) === Object.prototype"),
		).To(BeTrue())
	})

	It("Should allow iterating attributes", func() {
		ctx.Window().LoadHTML(`<body><div id="foo" class="bar" hidden></div></body>`)
		Expect(ctx.RunTestScript(`
const elm = document.getElementById("foo");
const attributes = elm.attributes;
let idAttribute;
for (let i = 0; i < attributes.length; i++) {
  const attr = attributes.item(i)
  if (attr.name === "id") { idAttribute = attr }
}`)).Error().ToNot(HaveOccurred())
		Expect(ctx.RunTestScript("attributes.length")).To(BeEquivalentTo(3))
		Expect(ctx.RunTestScript("idAttribute.value")).To(Equal("foo"))
		ctx.MustRunTestScript("idAttribute.value = 'bar'")
		Expect(ctx.RunTestScript("idAttribute.value")).To(Equal("bar"))
	})

	It("Should allow indexing by number", func() {
		ctx.Window().LoadHTML(`<body><div id="foo" class="bar" hidden></div></body>`)
		Expect(ctx.RunTestScript(`
const elm = document.getElementById("foo");
const attributes = elm.attributes;
attributes[0] instanceof Attr &&
attributes[1] instanceof Attr &&
attributes[2] instanceof Attr
`)).To(BeTrue())
	})

	It("Should return `null` when indexing outside the elements", func() {
		ctx.Window().LoadHTML(`<body><div id="foo" class="bar" hidden></div></body>`)
		Expect(ctx.RunTestScript(`
const elm = document.getElementById("foo");
const attributes = elm.attributes;
attributes[3]
`)).To(BeNil())
	})
})
