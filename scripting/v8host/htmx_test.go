package v8host_test

import (
	"fmt"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	// . "github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser"
	app "github.com/gost-dom/browser/internal/test/htmx-app"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
)

var _ = Describe("HTMX Tests", Ordered, func() {
	var b *browser.Browser
	BeforeEach(func() {
		b = browser.NewBrowserFromHandler(app.CreateServer())
		DeferCleanup(func() {
			b.Close()
		})
	})

	It("Should increment the counter example", func() {
		win, err := b.Open("/counter/index.html")
		Expect(err).ToNot(HaveOccurred())
		counter := win.Document().GetElementById("counter")
		Expect(counter).To(HaveInnerHTML(Equal("Count: 1")))
		counter.Click()
		counter = win.Document().GetElementById("counter")
		Expect(counter).To(HaveInnerHTML(Equal("Count: 2")))
	})

	It("Should not update the location when a link has hx-get", Focus, func() {
		win, err := b.Open("/navigation/page-a.html")
		Expect(err).ToNot(HaveOccurred())
		Expect(win.ScriptContext().Eval("window.pageA")).To(BeTrue())
		Expect(win.ScriptContext().Eval("window.pageB")).To(BeNil())
		fmt.Println("\n\nClick\n\n--")
		win.Document().GetElementById("link-to-b").Click()
		Expect(win.ScriptContext().Eval("window.pageA")).To(BeTrue())
		Expect(win.ScriptContext().Eval("window.pageB")).To(BeTrue())
		Expect(win.Document()).To(HaveH1("Page B"), "Page heading")
		Expect(win.Location().Pathname()).To(Equal("/navigation/page-a.html"))
	})

	It("Should update the location when a link with href is boosted", func() {
		win, err := b.Open("/navigation/page-a.html")
		Expect(err).ToNot(HaveOccurred())
		Expect(win.ScriptContext().Eval("window.pageA")).To(BeTrue())
		Expect(win.ScriptContext().Eval("window.pageB")).To(BeNil())
		fmt.Println("\n\nClick\n\n--")
		win.Document().GetElementById("link-to-b-boosted").Click()
		Expect(win.ScriptContext().Eval("window.pageA")).ToNot(BeNil(), "pageA exists")
		Expect(win.ScriptContext().Eval("window.pageA")).To(BeTrue(), "pageA value")
		Expect(win.ScriptContext().Eval("window.pageB")).ToNot(BeNil(), "pageB exists")
		Expect(win.ScriptContext().Eval("window.pageB")).To(BeTrue(), "pageB value")
		Expect(win.Document()).ToNot(BeNil())
		Expect(win.Document()).To(HaveH1("Page B"), "Page heading")
		Expect(win.Location().Pathname()).To(Equal("/navigation/page-b.html"))
	})
})
