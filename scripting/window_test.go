package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Window", func() {
	ctx := InitializeContext()

	Describe("Constructor", func() {
		It("Should be defined", func() {
			Expect(ctx.RunScript("Window")).ToNot(BeNil())
		})

		It("Should not be callable", func() {
			Expect(ctx.RunTestScript("Window()")).Error().To(HaveOccurred())
		})

		It("Should not be newable", func() {
			Expect(ctx.RunTestScript("new Window()")).Error().To(HaveOccurred())
		})
	})

	Describe("Inheritance", func() {
		It("Should be an EventTarget", func() {
			Expect(ctx.MustRunTestScript("window instanceof EventTarget")).To(BeTrue())
		})
	})
})