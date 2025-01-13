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
	var form HTMLFormElement
	var actualRequest *http.Request

	AfterEach(func() {
		// Make the values ready for garbage collection
		actualRequest = nil
		form = nil
	})

	BeforeEach(func() {
		actualBody = ""
		DeferCleanup(func() { requests = nil })

		window = NewWindow(WindowOptions{
			HttpClient: NewHttpClientFromHandler(
				http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					actualRequest = req
					requests = append(requests, req)
					data, err := io.ReadAll(req.Body)
					Expect(err).ToNot(HaveOccurred())
					actualBody = string(data)
				}),
			),
			BaseLocation: "http://example.com/forms/example-form.html",
		})

		Expect(
			window.LoadHTML(
				`<body>
					<form>
						<input name="foo" value="bar" />
					</form>
				</body>`,
			),
		).To(Succeed())

		el, err := window.Document().QuerySelector("form")
		Expect(err).ToNot(HaveOccurred())
		f, ok := el.(HTMLFormElement)
		Expect(ok).To(BeTrue())
		form = f
	})

	Describe("No method of action specified", func() {
		It("Should make a GET request to the original location", func() {
			form.Submit()
			Expect(actualRequest.Method).To(Equal("GET"))
			Expect(
				actualRequest.URL.String(),
			).To(Equal("http://example.com/forms/example-form.html?foo=bar"))
		})
	})

	Describe("The form is a POST", func() {
		BeforeEach(func() {
			form.SetAttribute("method", "POST")
		})

		It("Should make a POST request", func() {
			form.Submit()
			Expect(actualRequest.Method).To(Equal("POST"))
			Expect(
				actualRequest.URL.String(),
			).To(Equal("http://example.com/forms/example-form.html"))
		})

		It("Should store the values in the form body", func() {
			form.Submit()
			Expect(actualBody).To(Equal("foo=bar"))
		})

		It("Should resolve a relative 'action' without a ./.. prefix", func() {
			form.SetAttribute("action", "example-form-post-target")
			form.Submit()
			Expect(
				actualRequest.URL.String(),
			).To(Equal("http://example.com/forms/example-form-post-target"))
		})
	})
})
