package html_test

import (
	"github.com/gost-dom/browser/browser/dom"
	. "github.com/gost-dom/browser/browser/html"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Template loaded from body", func() {
	var doc dom.Document
	var template HTMLTemplateElement

	BeforeEach(func() {
		doc = ParseHtmlString(`<body><template id="t"><div id="d"></div></template></body>`)
		var ok bool
		template, ok = doc.Body().FirstChild().(HTMLTemplateElement)
		Expect(ok).To(BeTrue(), "Instance of HTMLTemplateElement")
	})

	It("Should not have any children", func() {
		Expect(template.ChildNodes().Length()).To(Equal(0))
	})

	It("Should render contents in body's OuterHTML", func() {
		Expect(
			doc.Body().OuterHTML(),
		).To(Equal(`<body><template id="t"><div id="d"></div></template></body>`))
	})

	It("Should render contents in OuterHTML", func() {
		Expect(template.OuterHTML()).To(Equal(`<template id="t"><div id="d"></div></template>`))
	})

	It("Should render contents in InnerHTML", func() {
		Expect(template.InnerHTML()).To(Equal(`<div id="d"></div>`))
	})

	Describe("CSS Queries", func() {
		It("Should not find it's children from querying document", func() {
			Expect(doc.QuerySelector("div")).To(BeNil())
		})

		It("Should not find it's children from querying document", func() {
			template, err := doc.QuerySelector("template")
			Expect(err).ToNot(HaveOccurred())
			Expect(template.QuerySelector("div")).To(BeNil())
		})

		It("Should not find it's children from querying template content", func() {
			template, err := doc.QuerySelector("template")
			Expect(template, err).ToNot(BeNil(), "Template was found")
			t, ok := template.(HTMLTemplateElement)
			Expect(ok).To(BeTrue())
			Expect(t.Content().QuerySelector("div")).ToNot(BeNil(), "Div found in template")
		})
	})
})
