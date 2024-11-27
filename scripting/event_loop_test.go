package scripting_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/stroiman/go-dom/browser"
)

var _ = Describe("EventLoop", func() {
	It("Defers execution", func() {
		ctx := NewTestContext()
		ctx.StartEventLoop()
		c := make(chan bool)
		defer close(c)

		ctx.Window().
			AddEventListener("go:home",
				NewEventHandlerFunc(func(e Event) error {
					c <- true
					return nil
				}))
		Expect(
			ctx.RunTestScript(`
				let val; setTimeout(() => {
					val = 42;
					window.dispatchEvent(new CustomEvent("go:home"));
				}, 1); 
				val`,
			),
		).To(BeNil())
		<-c
		Expect(ctx.RunTestScript(`val`)).To(BeEquivalentTo(42))
	})

	It("Dispatches an 'error' event on unhandled error", func() {
		ctx := NewTestContext(IgnoreUnhandledErrors)
		ctx.StartEventLoop()
		c := make(chan bool)
		defer close(c)
		Expect(
			ctx.RunTestScript(`
				let val;
				window.addEventListener('error', () => {
					val = 42;
				});
				setTimeout(() => {
					throw new Error()
				}, 1); 
				val`,
			),
		).To(BeNil())
		<-time.After(10 * time.Millisecond)
		Expect(ctx.RunTestScript(`val`)).To(BeEquivalentTo(42))
	})
})
