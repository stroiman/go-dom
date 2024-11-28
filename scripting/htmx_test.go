package scripting_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/browser"
	app "github.com/stroiman/go-dom/internal/test/htmx-app"
)

var _ = Describe("Load from server", Focus, Ordered, func() {
	It("Renders HTMX without errors", func() {
		server := app.CreateServer()
		DeferCleanup(func() { server = nil })
		browser := NewTestBrowserFromHandler(server)
		win, err := browser.OpenWindow("/index.html")
		Expect(err).Error().ToNot(HaveOccurred())
		called := make(chan bool)
		handler := NewEventHandlerFuncWithoutError(func(e Event) { called <- true })
		win.AddEventListener("htmx:load", handler)
		<-called
		// Eventually(called).Should(Receive())
		counter := win.Document().GetElementById("counter")
		Expect(counter.Click()).To(Succeed())
	})

	//	It("Has the right console functions", func() {
	//		server := app.CreateServer()
	//		DeferCleanup(func() { server = nil })
	//		browser := NewTestBrowserFromHandler(server)
	//		win, err := browser.OpenWindow("/index.html")
	//		Expect(err).ToNot(HaveOccurred())
	//		fmt.Println()
	//		fmt.Println(win.Eval(`Object.getOwnPropertyNames(console).join(',')`))
	//		win.Eval(`
	//
	// console.log('𐀀');
	// console.log('a');
	// console.error('error');
	// `)
	//
	//		return
	//	})
})
