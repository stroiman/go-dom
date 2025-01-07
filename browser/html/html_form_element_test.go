package html_test

import (
	"io"
	"net/http"

	. "github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/internal/http"

	//
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTML Form", func() {
	var window Window
	var actualBody string
	var requests []*http.Request

	BeforeEach(func() {
		actualBody = ""
		DeferCleanup(func() { requests = nil })

		window = NewWindow(WindowOptions{
			HttpClient: NewHttpClientFromHandler(
				http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					requests = append(requests, req)
					data, err := io.ReadAll(req.Body)
					Expect(err).ToNot(HaveOccurred())
					actualBody = string(data)
				}),
			),
		})
		Expect(
			window.LoadHTML(
				`<body>
					<form method="POST" action="/submitted">
						<input name="foo" value="bar" />
					</form>
				</body>`,
			),
		).To(Succeed())
	})

	It("Should submit form values to the server", func() {
		formElm, _ := window.Document().QuerySelector("form")
		form := formElm.(HTMLFormElement)
		form.Submit()
		Expect(actualBody).To(Equal("foo=bar"))
	})
})
