package html_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTMLButtonElement", func() {
	Describe("Clicking a type='submit' button outside a form", func() {
		It("Should be ignored", func() {
			doc := ParseHtmlString("<body><button id='btn' type='submit'>Click me!</button></body>")
			button := doc.GetElementById("btn")
			Expect(button.Click()).To(BeTrue())
		})
	})
})
