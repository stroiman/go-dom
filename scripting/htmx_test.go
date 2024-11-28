package scripting_test

import (
	"log/slog"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/browser"
	app "github.com/stroiman/go-dom/internal/test/htmx-app"
)

var _ = Describe("Load from server", Ordered, func() {
	It("Renders HTMX without errors", func() {
		server := app.CreateServer()
		DeferCleanup(func() { server = nil })

		browser := NewTestBrowserFromHandler(server)
		win, err := browser.OpenWindow("/index.html")
		win.AddEventListener("htmx:error", NewEventHandlerFunc(func(e Event) error {
			slog.Error("ERRROR", "msg", e)
			return nil
		}))
		Expect(err).Error().ToNot(HaveOccurred())

		called := make(chan bool)
		swap := make(chan bool)
		win.AddEventListener(
			"htmx:load",
			NewEventHandlerFuncWithoutError(func(e Event) { go func() { called <- true }() }),
		)
		// <-called // Technically, we should wait. But the test seems to have settled
		counter := win.Document().GetElementById("counter")
		win.AddEventListener("htmx:after-swap", NewEventHandlerFuncWithoutError(func(e Event) {
			go func() { swap <- true }()
		}))
		counter = win.Document().GetElementById("counter")
		counter.Click()
		// <-swap // Technically, we should wait. But the test seems to have settled
		// allready
		counter = win.Document().GetElementById("counter")
		Expect(counter.InnerHTML()).To(Equal("Count: 2"))
	})
})
