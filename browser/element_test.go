package browser_test

import (
	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Element", func() {
	Describe("SetAttribute", func() {
		It("Should add a new attribute when not existing", func() {
			doc := NewDocument()
			elm := doc.CreateElement("div")
			Expect(elm.GetAttributes().Length()).To(Equal(0))
			elm.SetAttribute("id", "1")
			Expect(elm.GetAttributes().Length()).To(Equal(1))
		})

		It("Should add overwrite an existing attribute", func() {
			doc := NewDocument()
			elm := doc.CreateElement("div")
			elm.SetAttribute("id", "1")
			elm.SetAttribute("id", "2")
			Expect(elm.GetAttribute("id")).To(Equal("2"))
			Expect(elm.GetAttributes().Length()).To(Equal(1))
		})
	})
})
