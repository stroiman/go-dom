package dom_types_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/dom-types"
	"github.com/stroiman/go-dom/interfaces"
)

var _ = Describe("Document", func() {
	var doc interfaces.Document

	BeforeEach(func() {
		doc = NewDocument()
	})

	Describe("NewDocument", func() {
		It("Should have an empty body", func() {
			Expect(doc.DocumentElement()).To(BeNil())
		})

		It("Should have zero children", func() {
			Skip("ChildNodes not yet implemented")
		})

		Describe("Building a valid HTML structure", func() {
			Describe("Append an HTML element", func() {
				var html interfaces.HTMLElement

				BeforeEach(func() {
					html = doc.CreateElement("html")
				})
				JustBeforeEach(func() {
					doc.Append(html)
				})

				It("Should set the documentElement", func() {
					Expect(doc.DocumentElement()).To(BeIdenticalTo(html))
				})

				It("Should leave body", func() {
					Expect(doc.Body()).To(BeNil())
				})

				It("Should set the body when added to the html element", func() {
					body := doc.CreateElement("body")
					html.Append(body)
					Expect(doc.Body()).To(BeIdenticalTo(body))
				})

				It("Should set the 'children' of the html element", func() {
					Skip("TODO")
					body := doc.CreateElement("body")
					html.Append(body)
					Expect(html.Children()).To(HaveExactElements(BeIdenticalTo(body)))
				})

				It("Should connect the body", func() {
					Skip("TODO")
					body := doc.CreateElement("body")
					html.Append(body)
					Expect(doc.Body().NodeName()).To(Equal("BODY"))
					Expect(doc.Body().IsConnected()).To(BeTrue())
				})
			})
		})

		Describe("Append bad root", func() {
			It("Should set the document", func() {
				Skip("TODO")
			})
			It("Should generate when adding two children", func() {
				Skip("TODO")
			})
			It("Should not set document.body when adding a body element", func() {
				Skip("TODO")
			})
			It("Should not set document.head when adding a head element", func() {
				Skip("TODO")
			})
		})
	})

	Describe("CreateElement", func() {
		It("Should return a non-connected element", func() {
			Expect(
				NewDocument().CreateElement("html").IsConnected(),
			).To(BeFalse())
		})
	})
})
