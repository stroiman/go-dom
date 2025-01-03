package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/browser/scripting"
)

var _ = Describe("V8 Document", func() {
	ctx := InitializeContextWithEmptyHtml()

	itShouldBeADocument := func() {
		It("Should be an instance of Document", func() {
			Expect(ctx.RunTestScript("actual instanceof Document")).To(BeTrue())
		})
		It("Should have nodeType 9", func() {
			Expect(ctx.RunTestScript("actual.nodeType")).To(BeEquivalentTo(9))
		})
		It("Should be an instance of Node", func() {
			Expect(ctx.RunTestScript("actual instanceof Node")).To(BeTrue())
		})
		It("Should be an instance of EventTarget", func() {
			Expect(ctx.RunTestScript("actual instanceof EventTarget")).To(BeTrue())
		})
		It("Should be an instance of Object", func() {
			Expect(ctx.RunTestScript("actual instanceof Object")).To(BeTrue())
		})
		It("Should have a class hierarchy of 4 classes", func() {
			Expect(ctx.RunTestScript(`
        let baseClassCount = 0
        let current = actual
        while(current = Object.getPrototypeOf(current))
          baseClassCount++
        baseClassCount;
      `)).To(BeEquivalentTo(4))
		})
	}

	Describe("Class Hierarchy of new Document()", func() {
		BeforeEach(func() {
			ctx.RunTestScript("const actual = new Document()")
		})
		itShouldBeADocument()
	})

	Describe("Class Hierarchy of `window.document`", func() {
		BeforeEach(func() {
			ctx.RunTestScript("const actual = window.document")
		})
		itShouldBeADocument()
	})

	Describe("Constructor", func() {
		It("Should be instance of Document", func() {
			Expect(ctx.RunTestScript(`
        const doc = new Document();
        doc instanceof Document && doc != document;
      `)).To(BeTrue())
		})

		Describe("createElement", func() {
			It("Should return an HTMLElement", func() {
				Expect(
					ctx.RunTestScript(`const base = document.createElement("base");
						base instanceof HTMLElement`),
				).To(BeTrue())
			})
		})

		It("Should support Document functions", func() {
			Skip("createElement and HTMLElement are missing")
			Expect(
				ctx.RunTestScript(`document.createElement("div") instanceof HTMLElement`),
			).Error().
				ToNot(HaveOccurred())
		})

		It("Should have a documentElement when HTML is loaded", func() {
			ctx := NewTestContext(LoadHTML(` <html> <body> </body> </html>`))
			Expect(
				ctx.RunTestScript(
					"Object.getPrototypeOf(document.documentElement).constructor.name",
				),
			).To(Equal("HTMLElement"))
		})

		It("Should have a documentElement instance of HTMLElement", func() {
			ctx := NewTestContext(LoadHTML(` <html> <body> </body> </html>`))
			Expect(
				ctx.RunTestScript("document.documentElement instanceof HTMLElement"),
			).To(BeTrue())
		})
	})

	Describe("location property", func() {
		It("Should be a Location", func() {
			Expect(ctx.RunTestScript("document.location instanceof Location")).To(BeTrue())
		})

		It("Should equal window.location", func() {
			Expect(ctx.RunTestScript("document.location === location")).To(BeTrue())
		})
	})

	Describe("body and Body", func() {
		It("document.body Should return a <body>", func() {
			ctx := NewTestContext(LoadHTML(`<html><body></body></html>`))
			Expect(ctx.RunTestScript("document.body.tagName")).To(Equal("BODY"))
		})
		It("document.head Should return a <head>", func() {
			ctx := NewTestContext(LoadHTML(`<html><body></body></html>`))
			Expect(ctx.RunTestScript("document.head.tagName")).To(Equal("HEAD"))
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
				ctx.RunTestScript("document.querySelector('[data-key]').outerHTML"),
			).To(Equal(`<div data-key="1">1</div>`))
			Expect(
				ctx.RunTestScript(`document.querySelector('[data-key="2"]').outerHTML`),
			).To(Equal(`<div data-key="2">2</div>`))
			Expect(
				ctx.RunTestScript(`document.querySelector('script')`),
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
				ctx.RunTestScript(
					"Array.from(document.querySelectorAll('[data-key]')).map(x => x.outerHTML).join(',')",
				),
			).To(Equal(`<div data-key="1">1</div>,<div data-key="2">2</div>`))
		})
	})

	Describe("createDocumentFragment", func() {
		It("Should return a DocumentFragment", func() {
			ctx := NewTestContext()
			Expect(ctx.RunTestScript(`
				const fragment = document.createDocumentFragment();
				Object.getPrototypeOf(fragment) === DocumentFragment.prototype
			`)).To(BeTrue())
		})
	})

	It("Should create document fragments", func() {
		Expect(ctx.RunTestScript(
			`Object.getPrototypeOf(document.createDocumentFragment()) === DocumentFragment.prototype`,
		)).To(BeTrue())
	})
})
