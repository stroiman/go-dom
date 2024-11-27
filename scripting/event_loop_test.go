package scripting_test

import (
	"time"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser"
)

var _ = Describe("EventLoop", func() {
	It("Defers execution", func() {
		ctx := NewTestContext()
		c := make(chan bool)
		defer close(c)

		ctx.Window().
			AddEventListener("go:home", browser.NewEventHandlerFunc(func(e browser.Event) error { c <- true; return nil }))
		Expect(ctx.RunTestScript(`let a; setTimeout(() => { a = 42 }, 1); a`)).To(BeNil())
		ctx.StartEventLoop()
		<-time.After(time.Millisecond * 10)
		Expect(ctx.RunTestScript(`a`)).To(BeEquivalentTo(42))
	})

	It("Should notify of errors", func() {
		Skip("TODO")
	})
})
