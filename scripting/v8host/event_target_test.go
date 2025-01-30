package v8host_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("V8 EventTarget", func() {
	It("Should be able to add and remove event listeners", func() {
		win := scriptTestSuite.NewWindow()
		Expect(win.Eval(`
			let events = [];
			let noOfEvents = []
			const handler = e => { events.push(e) }
			window.addEventListener("gost", handler)
			window.dispatchEvent(new CustomEvent("gost"))
			noOfEvents.push(events.length)
			window.removeEventListener("gost", handler)
			window.dispatchEvent(new CustomEvent("gost"))
			noOfEvents.push(events.length)
			window.addEventListener("gost", handler)
			window.dispatchEvent(new CustomEvent("gost"))
			noOfEvents.push(events.length)
			noOfEvents
		`)).To(HaveExactElements([]int32{1, 1, 2}))
	})
})
