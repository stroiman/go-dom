package dom_test

import (
	. "github.com/stroiman/go-dom/browser/dom"
	. "github.com/stroiman/go-dom/browser/testing/gomega-matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Element", func() {
	Describe("Get/Set attribute", func() {
		It("Should add a new attribute when not existing", func() {
			doc := CreateHTMLDocument()
			elm := doc.CreateElement("div")
			Expect(elm.Attributes().Length()).To(Equal(0))
			elm.SetAttribute("id", "1")
			Expect(elm.Attributes().Length()).To(Equal(1))
		})

		It("Should add overwrite an existing attribute", func() {
			doc := CreateHTMLDocument()
			elm := doc.CreateElement("div")
			elm.SetAttribute("id", "1")
			elm.SetAttribute("id", "2")
			Expect(elm).To(HaveAttribute("id", "2"))
			Expect(elm.Attributes().Length()).To(Equal(1))
		})

		It("Should return nil when the attribute does't exist", func() {
			doc := CreateHTMLDocument()
			elm := doc.CreateElement("div")
			_, ok := elm.GetAttribute("non-existing")
			Expect(ok).To(BeFalse())
		})
	})

	It("Should support Get/SetTextContent", func() {
		Skip("Only setter created yet")
	})

	Describe("InsertAdjacentHTML", func() {
		It("Should insert correctly 'beforeBegin'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el := doc.GetElementById("2")
			Expect(el.InsertAdjacentHTML(
				"beforebegin",
				"<div>1st new child</div><div>2nd new child</div>",
			)).To(Succeed())
			Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div>1st new child</div><div>2nd new child</div><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`))

		})

		It("Should insert correctly 'afterBegin'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el, err := (doc.QuerySelector("[id='2']"))
			Expect(err).ToNot(HaveOccurred())
			Expect(
				el.InsertAdjacentHTML(
					"afterbegin",
					"<div>1st new child</div><div>2nd new child</div>",
				),
			).To(Succeed())
			Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div id="2"><div>1st new child</div><div>2nd new child</div>El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`))

		})

		It("Should insert correctly 'beforeEnd'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el, err := (doc.QuerySelector("[id='2']"))
			Expect(err).ToNot(HaveOccurred())
			Expect(
				el.InsertAdjacentHTML(
					"beforeend",
					"<div>1st new child</div><div>2nd new child</div>",
				),
			).To(Succeed())
			Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  <div>1st new child</div><div>2nd new child</div></div>
  <div id="3">El 1</div>
</body>`))

		})

		It("Should insert correctly 'afterend'", func() {
			doc := ParseHtmlString(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>
  <div id="3">El 1</div>
</body>`)
			el, err := (doc.QuerySelector("[id='2']"))
			Expect(err).ToNot(HaveOccurred())
			Expect(
				el.InsertAdjacentHTML(
					"afterend",
					"<div>1st new child</div><div>2nd new child</div>",
				),
			).To(Succeed())
			Expect(doc.Body()).To(HaveOuterHTML(`<body>
  <div id="1">El 1</div>
  <div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div><div>1st new child</div><div>2nd new child</div>
  <div id="3">El 1</div>
</body>`))
		})
	})

	Describe("HTML Rendering", func() {
		It("Should support OuterHTML", func() {
			doc := ParseHtmlString(`<body><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div></body>`)
			Expect(doc.Body().OuterHTML()).To(Equal(`<body><div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div></body>`))
			Expect(doc.Body().InnerHTML()).To(Equal(`<div id="2">El 2
    <div>El 2-a</div>
    <div>El 2-b</div>
  </div>`))
		})
	})

	Describe("Click", func() {
		It("Is cancelable and bubbles", func() {
			var event Event
			doc := ParseHtmlString(`<body><div id="target"></div></body>`)
			element := doc.GetElementById("target")
			element.AddEventListener("click", NewEventHandlerFuncWithoutError(func(e Event) {
				event = e
			}))
			element.Click()
			Expect(event.Cancelable()).To(BeTrue())
			Expect(event.Bubbles()).To(BeTrue())
		})
	})
})
