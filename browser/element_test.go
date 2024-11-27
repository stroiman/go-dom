package browser_test

import (
	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Element", func() {
	Describe("SetAttribute", func() {
		It("Should add a new attribute when not existing", func() {
			doc := NewWindow(nil).Document()
			elm := doc.CreateElement("div")
			Expect(elm.GetAttributes().Length()).To(Equal(0))
			elm.SetAttribute("id", "1")
			Expect(elm.GetAttributes().Length()).To(Equal(1))
		})

		It("Should add overwrite an existing attribute", func() {
			doc := NewWindow(nil).Document()
			elm := doc.CreateElement("div")
			elm.SetAttribute("id", "1")
			elm.SetAttribute("id", "2")
			Expect(elm.GetAttribute("id")).To(Equal("2"))
			Expect(elm.GetAttributes().Length()).To(Equal(1))
		})
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
			node := doc.GetElementById("2")
			if el, ok := node.(Element); ok {
				Element(
					el,
				).InsertAdjacentHTML("beforebegin", "<div>1st new child</div><div>2nd new child</div>")
			}
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
			node, err := (doc.QuerySelector("[id='2']"))
			Expect(err).ToNot(HaveOccurred())
			if el, ok := node.(Element); ok {
				Element(
					el,
				).InsertAdjacentHTML("afterbegin", "<div>1st new child</div><div>2nd new child</div>")
			}
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
			node, err := (doc.QuerySelector("[id='2']"))
			Expect(err).ToNot(HaveOccurred())
			if el, ok := node.(Element); ok {
				Element(
					el,
				).InsertAdjacentHTML("beforeend", "<div>1st new child</div><div>2nd new child</div>")
			}
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
			node, err := (doc.QuerySelector("[id='2']"))
			Expect(err).ToNot(HaveOccurred())
			if el, ok := node.(Element); ok {
				Element(
					el,
				).InsertAdjacentHTML("afterend", "<div>1st new child</div><div>2nd new child</div>")
			}
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
})
