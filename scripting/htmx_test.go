package scripting_test

import (
	"net/http"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/browser"
)

var _ = Describe("Load from server", Ordered, func() {
	var server *http.ServeMux

	BeforeAll(func() {
		server = http.NewServeMux()
		server.HandleFunc("GET /index.html", func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(indexHTML))
		})
		server.Handle("/public/", http.StripPrefix("/public/",
			http.FileServer(http.Dir("./content"))))
	})
	AfterAll(func() {
		server = nil // Ready for GC
	})

	It("Renders without errors", func() {
		browser := NewTestBrowserFromHandler(server)
		win, err := browser.OpenWindow("/index.html")
		Expect(err).Error().ToNot(HaveOccurred())
		called := make(chan bool)
		handler := NewEventHandlerFuncWithoutError(func(e Event) { called <- true })
		win.AddEventListener("htmx:load", handler)
		Eventually(called).Should(Receive())
	})
})

const indexHTML = `<html><head>
  <script src="/public/xpath.js"></script>
  <script>
  const { XPathExpression, XPathResult } = window;
  const evaluate = XPathExpression.prototype.evaluate;
  XPathExpression.prototype.evaluate = function (context, type, res) {
    return evaluate.call(this, context, type ?? XPathResult.ANY_TYPE, res);
  };
  </script>
  <script src="/public/htmx.js"></script>
  </head><body>
  <div hx-get="/increment">Count: 1</div>
</body></html>`
