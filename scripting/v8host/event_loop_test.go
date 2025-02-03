package v8host_test

import (
	"time"

	. "github.com/gost-dom/browser/dom"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventLoop", func() {
	It("Defers execution", func() {
		Skip("Hmmm")
		ctx := NewTestContext()
		c := make(chan bool)
		defer close(c)

		ctx.Window().
			AddEventListener("go:home",
				NewEventHandlerFunc(func(e Event) error {
					c <- true
					return nil
				}))
		Expect(
			ctx.Eval(`
				let val; setTimeout(() => {
					val = 42;
					window.dispatchEvent(new CustomEvent("go:home"));
				}, 1); 
				val`,
			),
		).To(BeNil())
		<-c
		Expect(ctx.Eval(`val`)).To(BeEquivalentTo(42))
	})

	It("Dispatches an 'error' event on unhandled error", func() {
		ctx := NewTestContext(IgnoreUnhandledErrors)
		Expect(
			ctx.Eval(`
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
		Expect(ctx.Eval(`val`)).To(BeEquivalentTo(42))
	})
})
