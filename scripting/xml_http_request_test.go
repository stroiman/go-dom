package scripting_test

import (
	"io"
	"net/http"

	"github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	// . "github.com/stroiman/go-dom/scripting"
)

var _ = Describe("V8 XmlHttpRequest", func() {
	var server http.Handler
	var window browser.Window
	var evt chan bool
	var body []byte
	var actualPath string

	BeforeEach(func() {
		evt = make(chan bool)
		server = http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
			if req.Body != nil {
				var err error
				body, err = io.ReadAll(req.Body)
				if err != nil {
					panic(err)
				}
			}
			actualPath = req.URL.Path
		})
		br := NewTestBrowserFromHandler(server)
		var err error
		window, err = br.NewWindow("http://example.com")
		Expect(err).ToNot(HaveOccurred())
		DeferCleanup(func() {
			window.Dispose()
			br.Dispose()
			server = nil
			close(evt)
			evt = nil
		})
	})

	It("Should inherit from EventTarget", func() {
		ctx := NewTestContext()
		Expect(ctx.RunTestScript("new XMLHttpRequest() instanceof EventTarget")).To(BeTrue())
	})

	It("Should dispatch 'load' event", func() {
		window.AddEventListener("go:home", browser.NewEventHandlerFunc(func(e browser.Event) error {
			go func() {
				evt <- true
			}()
			return nil
		}))
		Expect(window.Run(`
			const xhr = new XMLHttpRequest();
			let loadEvent;
			let loadendEvent;
			xhr.addEventListener("load", e => {
				loadEvent = e
				window.dispatchEvent(new CustomEvent("go:home"));
			})
			xhr.addEventListener("loadend", e => {
				loadendEvent = e
				window.dispatchEvent(new CustomEvent("go:home"));
			})
			xhr.open("GET", "/", true);
			xhr.send()
		`)).To(Succeed())
		<-evt
	})

	It("Should dispatch 'load' to `onload` function", func() {
		window.AddEventListener("go:home", browser.NewEventHandlerFunc(func(e browser.Event) error {
			go func() {
				evt <- true
			}()
			return nil
		}))
		Expect(window.Run(`
			const xhr = new XMLHttpRequest();
			let loadEvent;
			let loadendEvent;
			xhr.onload = function() {
				window.dispatchEvent(new CustomEvent("go:home"));
				loadEvent = e
			}
			xhr.onloadend = function() {
				loadendEvent = e
			}
			xhr.open("GET", "/PATH", true);
			xhr.send()
		`)).To(Succeed())
		<-evt
		Expect(actualPath).To(Equal("/PATH"))
	})

	Describe("Send", func() {
		It("Should send with `null`", func() {
			Expect(window.Eval(`
				const xhr = new XMLHttpRequest();
				xhr.open("GET", "/");
				xhr.send(null)
				xhr.status
			`)).To(BeEquivalentTo(200))
		})

		It("Should be able to send formdata", func() {
			Expect(window.Eval(`
				const xhr = new XMLHttpRequest();
				const data = new FormData()
				data.append("k1", "v1")
				data.append("k2", "v2")
				xhr.open("GET", "/");
				xhr.send(data)
				xhr.status
			`)).To(BeEquivalentTo(200))
			Expect(string(body)).To(Equal("k1=v1&k2=v2"))
		})

		It("Should be able to send a string", func() {
			Expect(window.Eval(`
				const xhr = new XMLHttpRequest();
				xhr.send("body contents")
				xhr.status
			`)).To(BeEquivalentTo(200))
			Expect(string(body)).To(Equal("body contents"))
		})
	})
})
