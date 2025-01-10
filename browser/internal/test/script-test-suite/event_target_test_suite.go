package suite

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.com/stroiman/go-dom/browser/dom"
)

func (suite *ScriptTestSuite) CreateEventTargetTests() {
	prefix := suite.Prefix

	Describe(prefix+"EventTarget", func() {
		var ctx *ScriptTestContext

		BeforeEach(func() {
			ctx = suite.NewContext()
		})

		It("Isn't cancellable by default", func() {
			Expect(ctx.Eval(`
				const target = new EventTarget();
				target.addEventListener("custom", e => { e.preventDefault() });
				target.dispatchEvent(new CustomEvent("custom"))
			`)).To(BeTrue())
		})

		It("Can be cancelled", func() {
			Expect(ctx.Eval(`
				const target = new EventTarget();
				target.addEventListener("custom", e => { e.preventDefault() });
				target.dispatchEvent(new CustomEvent("custom", {cancelable: true }))
			`)).To(BeFalse())
		})

		It("Doesn't bubble by default", func() {
			ctx.Window.LoadHTML(`<div id="parent"><div id="target"></div></div>`)
			ctx.Eval(`
				var targetCalled = false;
				var parentCalled = false;
				const target = document.getElementById("target")
				target.addEventListener("go:home", e => { targetCalled = true });
				document.getElementById("parent").addEventListener(
					"go:home",
					e => { parentCalled = true });
				target.dispatchEvent(new CustomEvent("go:home", {}))
			`)
			Expect(ctx.Eval("targetCalled")).To(BeTrue())
			Expect(ctx.Eval("parentCalled")).To(BeFalse())
		})

		It("Bubbles when specified in the constructor", func() {
			ctx.Window.LoadHTML(`<div id="parent"><div id="target"></div></div>`)
			Expect(ctx.Run(`
				var targetCalled = false;
				var parentCalled = false;
				const target = document.getElementById("target")
				target.addEventListener("go:home", e => { targetCalled = true });
				document.getElementById("parent").addEventListener(
					"go:home",
					e => { parentCalled = true });
				target.dispatchEvent(new CustomEvent("go:home", { bubbles: true }))
			`)).To(Succeed())
			Expect(ctx.Eval("targetCalled")).To(BeTrue())
			Expect(ctx.Eval("parentCalled")).To(BeTrue())
		})

		It("Is an EventTarget", func() {
			Expect(ctx.Eval("(new EventTarget()) instanceof EventTarget")).To(BeTrue())
		})

		It("Can call an added event listener", func() {
			Expect(ctx.Eval(`
	var callCount = 0
	function listener() { callCount++ };
	const target = new EventTarget();
	target.addEventListener('custom', listener);
	target.dispatchEvent(new CustomEvent('custom'));`)).Error().ToNot(HaveOccurred())
			Expect(ctx.Eval("callCount")).To(BeEquivalentTo(1))
		})

		It("Event from Go code will propagate to JS", func() {
			Expect(ctx.Eval(`
	var callCount = 0
	function listener() { callCount++ };
	const target = window;
	target.addEventListener('custom', listener);
	`)).Error().ToNot(HaveOccurred())
			ctx.Window.DispatchEvent(dom.NewCustomEvent("custom"))
			Expect(ctx.Eval("callCount")).To(BeEquivalentTo(1))
		})

		Describe("Events", func() {
			Describe("Custom events dispatched from Go-code", func() {
				It("Should be of type Event", func() {
					Expect(ctx.Eval(`
	var event;
	window.addEventListener('custom', e => { event = e });`,
					)).Error().ToNot(HaveOccurred())
					ctx.Window.DispatchEvent(dom.NewCustomEvent("custom"))
					Expect(
						ctx.Eval(`Object.getPrototypeOf(event) === CustomEvent.prototype`),
					).To(BeTrue())
					Expect(ctx.Eval(`event instanceof Event`)).To(BeTrue())
				})
			})
			It("Should have a type", func() {
				Expect(ctx.Eval(`
	var event;
	window.addEventListener('custom', e => { event = e });
	window.dispatchEvent(new CustomEvent('custom'));
	event.type`,
				)).To(Equal("custom"))
				By("Inheriting directly from event")
				Expect(
					ctx.Eval(`Object.getPrototypeOf(event) === CustomEvent.prototype`),
				).To(BeTrue())
			})
		})
	})
}
