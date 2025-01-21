package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	// . "github.com/stroiman/go-dom/browser/dom"
	"github.com/stroiman/go-dom/browser"
	app "github.com/stroiman/go-dom/browser/internal/test/htmx-app"
	. "github.com/stroiman/go-dom/browser/testing/gomega-matchers"
)

var _ = Describe("Load from server", Ordered, func() {
	It("Renders HTMX without errors", func() {
		browser := browser.NewBrowserFromHandler(app.CreateServer())
		DeferCleanup(func() {
			browser.Close()
		})
		win, err := browser.Open("/counter/index.html")
		Expect(err).ToNot(HaveOccurred())
		counter := win.Document().GetElementById("counter")
		Expect(counter).To(HaveInnerHTML(Equal("Count: 1")))
		// win.AddEventListener("htmx:after-swap", NewEventHandlerFuncWithoutError(func(e Event) {
		// 	go func() { swap <- true }()
		// }))
		// counter = win.Document().GetElementById("counter")
		// In principle, we should wait for an htmx:load event, currently `Open`
		// doesn't return until the window is fully loaded
		counter.Click()
		// <-swap // Technically, we should wait. But the test seems to have settled
		// allready
		// Again, we should wait for an event, but in practice, the test doesn't
		// continue until the XHR request has been processed
		counter = win.Document().GetElementById("counter")
		Expect(counter).To(HaveInnerHTML(Equal("Count: 2")))
	})
})
