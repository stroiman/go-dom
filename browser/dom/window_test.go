package dom_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/stroiman/go-dom/browser/dom"
)

var _ = Describe("Window", func() {
	It("Should have a document.documentElement instance of HTMLElement", func() {
		win, err := NewWindowReader(strings.NewReader("<html><body></body></html>"))
		Expect(err).ToNot(HaveOccurred())
		Expect(win.Document().DocumentElement()).To(BeHTMLElement())
	})

	It("Should respect the <!DOCTYPE>", func() {
		Skip("<!DOCTYPE> should be respected")
		// win, err := NewWindowReader(strings.NewReader("<!DOCTYPE HTML><html><body></body></html>"))
	})
})
