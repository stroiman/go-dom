package html_test

import (
	"io"
	"net/http"
	"net/url"

	"github.com/gost-dom/browser/dom"
	. "github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/internal/http"
	. "github.com/gost-dom/browser/testing/gomega-matchers"

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
				Expect(form).To(HaveAttribute("method", "new value"))
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
		var submittedForm url.Values

		AfterEach(func() {
			// Make the values ready for garbage collection
			actualRequest = nil
			form = nil
		})

		BeforeEach(func() {
			actualBody = ""
			DeferCleanup(func() { requests = nil; submittedForm = nil })

			window = NewWindow(WindowOptions{
				HttpClient: NewHttpClientFromHandler(
					http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
						if req.ParseForm() != nil {
							panic("Error parsing form")
						}
						actualRequest = req
						submittedForm = req.Form
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

		Describe("Method and action behaviour", func() {

			Describe("Action", func() {
				Describe("Get", func() {
					It("Should return the document location when not set", func() {
						Expect(form.GetAction()).To(Equal(window.Location().Href()))
					})
				})

				Describe("Set", func() {
					It(
						"Should update the action attribute, and update action property to relative url",
						func() {
							form.SetAction("/foo-bar")
							Expect(
								form,
							).To(HaveAttribute("action", "/foo-bar"))
							Expect(form.GetAction()).To(Equal("http://example.com/foo-bar"))
						},
					)
				})
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

		Describe("ReqeustSubmit with a <input type='submit'>", func() {
			var submitter dom.Element
			BeforeEach(func() {
				submitter = window.Document().CreateElement("input")
				submitter.SetAttribute("type", "submit")
				submitter.SetAttribute("name", "submitter")
				form.AppendChild(submitter)
			})

			It("Should add the name of a submitter, if passed", func() {
				form.RequestSubmit(submitter)
				Expect(submittedForm).To(HaveKey("submitter"))
			})

			It("Should ignore the name of a submitter if not passed", func() {
				form.RequestSubmit(nil)
				Expect(submittedForm).ToNot(HaveKey("submitter"))
			})
		})

		Describe("React to <button> click", func() {
			var button dom.Element

			BeforeEach(func() {
				button = window.Document().CreateElement("button")
				form.Append(button)
			})

			Describe("The button is a type='submit'", func() {
				BeforeEach(func() {
					button.SetAttribute("type", "submit")
				})

				It("Should submit the form", func() {
					button.Click()
					Expect(actualRequest).ToNot(BeNil())
				})

				It("Should also submit the form if 'type' was weird casing", func() {
					button.SetAttribute("type", "sUBmiT")
					button.Click()
					Expect(actualRequest).ToNot(BeNil())
				})

				It("Should not submit if preventDefault is called", func() {
					button.AddEventListener(
						"click",
						dom.NewEventHandlerFunc(func(e dom.Event) error {
							e.PreventDefault()
							return nil
						}),
					)
					button.Click()
					Expect(actualRequest).To(BeNil())
				})

				It("Should include the button name in the form data if set", func() {
					button.SetAttribute("name", "the-button")
					button.Click()
					Expect(submittedForm).To(HaveKey("the-button"))
				})
			})

			Describe("The button is not type='submit'", func() {
				It("should not submit the form", func() {
					button.SetAttribute("type", "reset")
					button.Click()
					Expect(actualRequest).To(BeNil())
				})
			})
		})

		Describe("Click a <input type='submit'>", func() {
			It("Should submit the form", func() {
				input := window.Document().CreateElement("input").(HTMLInputElement)
				input.SetType("submit")
				form.Append(input)
				input.Click()
				Expect(actualRequest).ToNot(BeNil())
			})
		})

		Describe("Click a <input type='reset'>", func() {
			It("Should submit the form", func() {
				input := window.Document().CreateElement("input").(HTMLInputElement)
				input.SetType("reset")
				form.Append(input)
				input.Click()
				Expect(actualRequest).To(BeNil())
			})
		})

		Describe("Dispatched events", func() {
			Describe("Submit event", func() {
				It("Should not be dispatched when form.submit is called", func() {
					var actualEvent dom.Event
					form.AddEventListener(
						"submit",
						dom.NewEventHandlerFunc(func(e dom.Event) error {
							actualEvent = e
							return nil
						}),
					)
					form.Submit()
					Expect(actualEvent).To(BeNil())
				})

				It("Should be dispatched when form.requestSubmit is called", func() {
					var actualEvent dom.Event
					form.AddEventListener(
						"submit",
						dom.NewEventHandlerFunc(func(e dom.Event) error {
							actualEvent = e
							return nil
						}),
					)
					form.RequestSubmit(nil)
					Expect(actualEvent).ToNot(BeNil())
					Expect(actualRequest).ToNot(BeNil())
				})

				It("Should be abort the request on preventDefault()", func() {
					form.AddEventListener(
						"submit",
						dom.NewEventHandlerFunc(func(e dom.Event) error {
							e.PreventDefault()
							return nil
						}),
					)
					form.RequestSubmit(nil)
					Expect(actualRequest).To(BeNil())
				})
			})

			Describe("formdata event", func() {
				It("Should be dispatched when a form is submitted", func() {
					var actualEvent dom.Event
					form.AddEventListener(
						"formdata",
						dom.NewEventHandlerFunc(func(e dom.Event) error {
							actualEvent = e
							return nil
						}),
					)
					form.Submit()
					Expect(actualEvent).ToNot(BeNil())
					Expect(actualEvent.Cancelable()).To(BeFalse())
					Expect(actualEvent.Bubbles()).To(BeTrue())
					formDataEvent, ok := actualEvent.(FormDataEvent)
					Expect(ok).To(BeTrue())
					Expect(formDataEvent.FormData()).ToNot(BeNil())
					Expect(formDataEvent.FormData()).To(HaveFormDataValue("foo", "bar"))
				})
			})
		})
	})
})
