package html_test

import (
	"io"
	"net/http"

	. "github.com/stroiman/go-dom/browser/html"
	. "github.com/stroiman/go-dom/browser/internal/http"
	matchers "github.com/stroiman/go-dom/browser/testing/gomega-matchers"

	//
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("HTML Form", func() {
	Describe("Method property", func() {
		var form HTMLFormElement
		BeforeEach(func() {
			doc := NewHTMLDocument(nil)
			form = doc.CreateElement("form").(HTMLFormElement)
		})

		Describe("Setting the value", func() {
			It("Should update the attribute", func() {
				form.SetMethod("new value")
				Expect(form).To(matchers.HaveAttribute("method", "new value"))
				Expect(form).ToNot(BeNil())
			})
		})

		Describe("Getting the value", func() {
			It("Should return 'get' by default", func() {
				Expect(form.GetMethod()).To(Equal("get"))
			})

			It("Should return 'post' when set to 'post', 'POST', 'PoSt', etc.", func() {
				for _, value := range []string{"post", "POST", "PoSt"} {
					form.SetMethod(value)
					Expect(form.GetMethod()).To(Equal("post"))
				}
			})

			It("Should return 'get' when assigned an invalid value", func() {
				form.SetMethod("post")
				Expect(form.GetMethod()).To(Equal("post"))
				form.SetMethod("invalid")
				Expect(form.GetMethod()).To(Equal("get"))
			})
		})
	})

	Describe("Submit behaviour", func() {
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
				BaseLocation: "http://example.com/forms/example-form.html?original-query=original-value",
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

			It("Should handle path lookup for relative paths", func() {
				form.SetAttribute("action", "submit-target")
				form.Submit()
				Expect(actualRequest.Method).To(Equal("GET"))
				Expect(
					actualRequest.URL.String(),
				).To(Equal("http://example.com/forms/submit-target?foo=bar"))
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
				).To(Equal("http://example.com/forms/example-form.html?original-query=original-value"))
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
})