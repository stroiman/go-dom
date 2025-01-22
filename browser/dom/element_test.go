package dom_test

import (
	. "github.com/stroiman/go-dom/browser/dom"
	. "github.com/stroiman/go-dom/browser/testing/gomega-matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Element", func() {
	var doc Document

	BeforeEach(func() {
		doc = CreateHTMLDocument()
		DeferCleanup(func() { doc = nil })
	})

	Describe("Get/Set attribute", func() {
		It("Should add a new attribute when not existing", func() {
			doc := CreateHTMLDocument()
			elm := doc.CreateElement("div")
			Expect(elm.Attributes().Length()).To(Equal(0))
			elm.SetAttribute("id", "1")
			Expect(elm.Attributes().Length()).To(Equal(1))
		})

		It("Should add overwrite an existing attribute", func() {
			elm := doc.CreateElement("div")
			elm.SetAttribute("id", "1")
			elm.SetAttribute("id", "2")
			Expect(elm).To(HaveAttribute("id", "2"))
			Expect(elm.Attributes().Length()).To(Equal(1))
		})

		It("Should return nil when the attribute does't exist", func() {
			elm := doc.CreateElement("div")
			_, ok := elm.GetAttribute("non-existing")
			Expect(ok).To(BeFalse())
		})
	})

	Describe("Matches", func() {
		It("Should return true for a simple string matching the root element", func() {
			d := doc.CreateElement("div")
			p := doc.CreateElement("p")
			d.Append(p)
			Expect(d.Matches("div")).To(BeTrue())
		})

		It("Should return false for a simple string matching a child element", func() {
			d := doc.CreateElement("div")
			p := doc.CreateElement("p")
			d.Append(p)
			Expect(d.Matches("p")).To(BeFalse())
		})

		It("Should return true for an existing attribute", func() {
			d := doc.CreateElement("div")
			d.SetAttribute("known-attribute", "")
			Expect(d.Matches("[known-attribute]")).To(BeTrue())
		})

		It("Should return true if one attribute match", func() {
			d := doc.CreateElement("div")
			d.SetAttribute("known-attribute", "")
			Expect(d.Matches("[unknown-attribute], [known-attribute]")).To(BeTrue())
		})

		It("Should return false for a non-existing attribute", func() {
			d := doc.CreateElement("div")
			Expect(d.Matches("[unknown-attribute]")).To(BeFalse())
		})

		It("Should return true if tagname + attribute key=value has right value", func() {
			d := doc.CreateElement("div")
			d.SetAttribute("a", "right")
			Expect(d.Matches(`div[a="right"]`)).To(BeTrue())
		})

		It("Should return true if tagname + attribute key=value has wrong value", func() {
			d := doc.CreateElement("div")
			d.SetAttribute("a", "right")
			Expect(d.Matches(`div[a="wrong"]`)).To(BeFalse())
		})
	})

	It("Should support Get/SetTextContent", func() {
		d := doc.CreateElement("div")
		d.AppendChild(doc.CreateElement("p"))
		d.SetTextContent("Replace the p")
		Expect(d).To(HaveTextContent("Replace the p"))
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
