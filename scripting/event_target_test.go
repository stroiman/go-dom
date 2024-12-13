package scripting_test

import (
	"github.com/stroiman/go-dom/browser"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("EventTarget", func() {
	ctx := InitializeContext()

	It("Doesn't bubble by default", func() {
		ctx := NewTestContext(LoadHTML(`<div id="parent"><div id="target"></div></div>`))
		ctx.MustRunTestScript(`
			var targetCalled = false;
			var parentCalled = false;
			const target = document.getElementById("target")
			target.addEventListener("go:home", e => { targetCalled = true });
			document.getElementById("parent").addEventListener(
				"go:home",
				e => { parentCalled = true });
			target.dispatchEvent(new CustomEvent("go:home", {}))
		`)
		Expect(ctx.RunTestScript("targetCalled")).To(BeTrue())
		Expect(ctx.RunTestScript("parentCalled")).To(BeFalse())
	})

	It("Bubbles when specified in the constructor", func() {
		ctx := NewTestContext(LoadHTML(`<div id="parent"><div id="target"></div></div>`))
		ctx.MustRunTestScript(`
			var targetCalled = false;
			var parentCalled = false;
			const target = document.getElementById("target")
			target.addEventListener("go:home", e => { targetCalled = true });
			document.getElementById("parent").addEventListener(
				"go:home",
				e => { parentCalled = true });
			target.dispatchEvent(new CustomEvent("go:home", { bubbles: true }))
		`)
		Expect(ctx.RunTestScript("targetCalled")).To(BeTrue())
		Expect(ctx.RunTestScript("parentCalled")).To(BeTrue())
	})

	It("Is an EventTarget", func() {
		Expect(ctx.RunTestScript("(new EventTarget()) instanceof EventTarget")).To(BeTrue())
	})

	It("Can call an added event listener", func() {
		Expect(ctx.RunTestScript(`
var callCount = 0
function listener() { callCount++ };
const target = new EventTarget();
target.addEventListener('custom', listener);
target.dispatchEvent(new CustomEvent('custom'));`)).Error().ToNot(HaveOccurred())
		Expect(ctx.RunTestScript("callCount")).To(BeEquivalentTo(1))
	})

	It("Event from Go code will propagate to JS", func() {
		Expect(ctx.RunTestScript(`
var callCount = 0
function listener() { callCount++ };
const target = window;
target.addEventListener('custom', listener);
`)).Error().ToNot(HaveOccurred())
		ctx.Window().DispatchEvent(browser.NewCustomEvent("custom"))
		Expect(ctx.RunTestScript("callCount")).To(BeEquivalentTo(1))
	})

	Describe("Events", func() {
		Describe("Custom events dispatched from Go-code", func() {
			It("Should be of type Event", func() {
				Expect(ctx.RunTestScript(`
var event;
window.addEventListener('custom', e => { event = e });`,
				)).Error().ToNot(HaveOccurred())
				ctx.Window().DispatchEvent(browser.NewCustomEvent("custom"))
				Expect(
					ctx.RunTestScript(`Object.getPrototypeOf(event) === CustomEvent.prototype`),
				).To(BeTrue())
				Expect(
					ctx.RunTestScript(`event instanceof Event`),
				).To(BeTrue())
			})
		})
		It("Should have a type", func() {
			Expect(ctx.RunTestScript(`
var event;
window.addEventListener('custom', e => { event = e });
window.dispatchEvent(new CustomEvent('custom'));
event.type`,
			)).To(Equal("custom"))
			By("Inheriting directly from event")
			Expect(
				ctx.RunTestScript(`Object.getPrototypeOf(event) === CustomEvent.prototype`),
			).To(BeTrue())
		})
	})
})
