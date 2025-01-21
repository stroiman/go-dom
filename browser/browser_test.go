package browser_test

import (
	"net/http"

	. "github.com/stroiman/go-dom/browser"
	. "github.com/stroiman/go-dom/browser/dom"
	. "github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/testing/gomega-matchers"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Browser", func() {
	It("Should be able to read from an http.Handler instance", func() {
		handler := (http.HandlerFunc)(func(w http.ResponseWriter, r *http.Request) {
			w.WriteHeader(200)
			w.Header().Add("Content-Type", "text/html") // For good measure, not used yet"
			w.Write([]byte("<html></html>"))
		})
		browser := NewBrowserFromHandler(handler)
		result, err := browser.Open("/")
		Expect(err).ToNot(HaveOccurred())
		element := result.Document().DocumentElement()

		Expect(element.NodeName()).To(Equal("HTML"))
		Expect(element.TagName()).To(Equal("HTML"))
	})

	It("Executes scripts", func() {
		// This is not necessarily desired behaviour right now.
		server := http.NewServeMux()
		server.Handle(
			"GET /index.html",
			http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
				res.Write([]byte(`<body>
					<div id='target'></div>
					<script>
						const target = document.getElementById('target');
						target.textContent = "42"
					</script>
				</body>`))
			}),
		)
		browser := NewBrowserFromHandler(server)
		DeferCleanup(func() { browser.Close() })
		win, err := browser.Open("/index.html")
		Expect(err).ToNot(HaveOccurred())
		target := win.Document().GetElementById("target")
		Expect(target).To(HaveOuterHTML(Equal(`<div id="target">42</div>`)))
	})

	Describe("Navigation", func() {
		Describe("Page A has been loaded", func() {
			var (
				window Window
			)

			BeforeEach(func() {
				var err error
				server := newBrowserNavigateTestServer()
				DeferCleanup(func() { server = nil })
				browser := NewBrowserFromHandler(server)
				window, err = browser.Open("/a.html")
				Expect(err).ToNot(HaveOccurred())
			})

			It("Should have Page A loaded", func() {
				heading, _ := window.Document().QuerySelector("h1")
				Expect(heading).To(HaveTextContent(Equal("Page A")))
				Expect(window.ScriptContext().Eval("loadedA")).To(Equal("PAGE A"))
			})

			Describe("Navigate to new page", func() {
				BeforeEach(func() {
					anchor, _ := window.Document().QuerySelector("a")
					anchor.Click()
					// TODO, Wait for load?
				})

				It("Should load a new page when clicking a link", func() {
					heading, _ := window.Document().QuerySelector("h1")
					Expect(heading).To(HaveTextContent(Equal("Page B")))
					Expect(window.ScriptContext().Eval("loadedB")).To(Equal("PAGE B"))
				})

				It("Should have cleared global JS state", func() {
					Expect(window.ScriptContext().Eval("typeof loadedA")).To(Equal("undefined"))
				})
			})

			It("Should NOT load a new page when an eventhandler aborts", func() {
				anchor, _ := window.Document().QuerySelector("a")
				anchor.AddEventListener(
					"click",
					NewEventHandlerFunc(NoError(Event.PreventDefault)),
				)
				anchor.Click()
				heading, _ := window.Document().QuerySelector("h1")
				Expect(heading).To(HaveTextContent(Equal("Page A")))
			})
		})
	})
})

func newBrowserNavigateTestServer() http.Handler {
	server := http.NewServeMux()
	server.HandleFunc("GET /a.html",
		func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(
				`<body>
					<h1>Page A</h1>
					<a href="b.html">Load B</a>
					<script>loadedA = "PAGE A"</script>
				</body>`))
		})

	server.HandleFunc("GET /b.html",
		func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(`
				<body>
					<h1>Page B</h1>
					<script>loadedB = "PAGE B"</script>
				</body>`))
		})

	return server
}
