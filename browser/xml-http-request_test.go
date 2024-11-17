package browser_test

import (
	"net/http"

	. "github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func newFromHandlerFunc(f func(http.ResponseWriter, *http.Request)) *XmlHttpRequest {
	client := http.Client{
		Transport: TestRoundTripper{http.HandlerFunc(f)},
	}
	return NewXmlHttpRequest(client)
}

var _ = Describe("XmlHTTPRequest", func() {
	Describe("Synchronous calls", func() {
		// Most browsers have deprecated this on the main thread, so throrough
		// support is not necessary until the code support WebWorkers.

		// It was written as the first test as it's the easier case to deal with
		Describe("Request succeeds", func() {
			It("Can make a request", func() {
				handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
					w.Header().Add("Content-Type", "text/plain")
					w.Write([]byte("Hello, World!"))
				})
				client := http.Client{
					Transport: TestRoundTripper{handler},
				}
				r := NewXmlHttpRequest(client)
				r.Open("GET", "/dummy")
				Expect(r.Status()).To(Equal(0))
				Expect(r.Send()).To(Succeed())
				Expect(r.Status()).To(Equal(200))
				// This is the only place we test StatusText; it's dumb wrapper and may
				// be removed.
				Expect(r.StatusText()).To(Equal("OK"))
				Expect(r.ResponseText()).To(Equal("Hello, World!"))
			})
		})
	})

	Describe("Asynchronous calls", func() {
		It("Emits a 'loadStart' event", func() {
			var (
				loadStarted bool
				loadEnded   bool
				loaded      bool
			)
			handler := http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
				w.Header().Add("Content-Type", "text/plain")
				w.Write([]byte("Hello, World!"))
			})
			client := http.Client{
				Transport: TestRoundTripper{handler},
			}
			r := NewXmlHttpRequest(client)
			r.Open("GET", "/dummy", RequestOptionAsync(true))
			r.AddEventListener(XHREventLoadstart, NewEventHandlerFunc(func(e Event) error {
				loadStarted = true
				return nil
			}))
			r.AddEventListener(XHREventLoadend, NewEventHandlerFunc(func(e Event) error {
				loadEnded = true
				return nil
			}))
			ended := make(chan bool)
			r.AddEventListener(XHREventLoad, NewEventHandlerFunc(func(e Event) error {
				loaded = true
				ended <- true
				return nil
			}))
			By("Sending the request")
			Expect(r.Send()).To(Succeed())
			Expect(r.Status()).To(Equal(0), "Response should not have been received yet")
			Expect(loadStarted).To(BeTrue(), "loadstart emitted")
			Expect(loadEnded).To(BeFalse(), "loadend emitted")
			Expect(loaded).To(BeFalse(), "load emitted")

			By("Receiving load event")
			<-ended

			By("The response should be a success")
			Expect(r.Status()).To(Equal(200))
			Expect(r.ResponseText()).To(Equal("Hello, World!"))
		})
	})

	It("Should handle redirects 'correctly'", func() {
		Skip(
			"Don't know what's the proper handling of redirects; I assume that it's to not do it. But Go will follow them by default",
		)
	})
})
