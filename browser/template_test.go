package browser_test

import (
	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Template loaded from body", func() {
	var doc Document
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

	It("Should render contents in InnerHTML", Pending, func() {
		Expect(template.InnerHTML()).To(Equal(`<div id="d"></div>`))
	})
})
