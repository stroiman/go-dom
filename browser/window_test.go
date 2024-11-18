package browser_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/browser"
)

var _ = Describe("Window", func() {
	It("Should have a document.documentElement instance of HTMLElement", func() {
		win := NewWindow(nil)
		win.LoadHTML("")
		Expect(win.Document().DocumentElement()).To(BeHTMLElement())
	})
})
