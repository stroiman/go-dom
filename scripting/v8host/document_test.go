package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("V8 Document", func() {
	ctx := InitializeContextWithEmptyHtml()

	Describe("Constructor", func() {
		It("Should have a documentElement when HTML is loaded", func() {
			ctx := NewTestContext(LoadHTML(` <html> <body> </body> </html>`))
			Expect(
				ctx.Eval(
					"Object.getPrototypeOf(document.documentElement).constructor.name",
				),
			).To(Equal("HTMLElement"))
		})

		It("Should have a documentElement instance of HTMLElement", func() {
			ctx := NewTestContext(LoadHTML(` <html> <body> </body> </html>`))
			Expect(
				ctx.Eval("document.documentElement instanceof HTMLElement"),
			).To(BeTrue())
		})
	})

	Describe("location property", func() {
		It("Should be a Location", func() {
			Expect(ctx.Eval("document.location instanceof Location")).To(BeTrue())
		})

		It("Should equal window.location", func() {
			Expect(ctx.Eval("document.location === location")).To(BeTrue())
		})
	})

	Describe("body and Body", func() {
		It("document.body Should return a <body>", func() {
			ctx := NewTestContext(LoadHTML(`<html><body></body></html>`))
			Expect(ctx.Eval("document.body.tagName")).To(Equal("BODY"))
		})
		It("document.head Should return a <head>", func() {
			ctx := NewTestContext(LoadHTML(`<html><body></body></html>`))
			Expect(ctx.Eval("document.head.tagName")).To(Equal("HEAD"))
		})
	})

	Describe("querySelector", func() {
		It("can find the right element", func() {
			ctx := NewTestContext(
				LoadHTML(
					`<body><div>0</div><div data-key="1">1</div><div data-key="2">2</div><body>`,
				),
			)
			Expect(
				ctx.Eval("document.querySelector('[data-key]').outerHTML"),
			).To(Equal(`<div data-key="1">1</div>`))
			Expect(
				ctx.Eval(`document.querySelector('[data-key="2"]').outerHTML`),
			).To(Equal(`<div data-key="2">2</div>`))
			Expect(
				ctx.Eval(`document.querySelector('script')`),
			).To(BeNil())
		})
	})

	Describe("querySelectorAll", func() {
		It("can find the right element", func() {
			ctx := NewTestContext(
				LoadHTML(
					`<body><div>0</div><div data-key="1">1</div><div data-key="2">2</div><body>`,
				),
			)
			Expect(
				ctx.Eval(
					"Array.from(document.querySelectorAll('[data-key]')).map(x => x.outerHTML).join(',')",
				),
			).To(Equal(`<div data-key="1">1</div>,<div data-key="2">2</div>`))
		})
	})

	Describe("createDocumentFragment", func() {
		It("Should return a DocumentFragment", func() {
			ctx := NewTestContext()
			Expect(ctx.Eval(`
				const fragment = document.createDocumentFragment();
				Object.getPrototypeOf(fragment) === DocumentFragment.prototype
			`)).To(BeTrue())
		})
	})

	It("Should create document fragments", func() {
		Expect(ctx.Eval(
			`Object.getPrototypeOf(document.createDocumentFragment()) === DocumentFragment.prototype`,
		)).To(BeTrue())
	})
})
