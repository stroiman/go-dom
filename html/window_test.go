package html_test

import (
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/gost-dom/browser/dom"
	"github.com/gost-dom/browser/html"
	. "github.com/gost-dom/browser/html"
	"github.com/gost-dom/browser/internal/domslices"
	"github.com/gost-dom/browser/internal/testing"
	. "github.com/gost-dom/browser/testing/gomega-matchers"
	"github.com/gost-dom/browser/testing/testservers"
)

var _ = Describe("Window", func() {
	It("Should have a document.documentElement instance of HTMLElement", func() {
		win, err := NewWindowReader(strings.NewReader("<html><body></body></html>"))
		Expect(err).ToNot(HaveOccurred())
		Expect(win.Document().DocumentElement()).To(BeHTMLElement())
	})

	It("Should respect the <!DOCTYPE>", func() {
		win, err := NewWindowReader(strings.NewReader("<!DOCTYPE HTML><html><body></body></html>"))
		Expect(err).ToNot(HaveOccurred())
		Expect(win.Document().FirstChild().NodeType()).To(Equal(dom.NodeTypeDocumentType))
	})

	Describe("History()", func() {
		var win html.Window
		var h *testing.EchoHandler

		BeforeEach(func() {
			h = new(testing.EchoHandler)
			win = html.NewWindow(windowOptionHandler(h))
			DeferCleanup(func() { win = nil; h = nil }) // Allow GC
		})

		It("Should have a length of one when starting", func() {
			Expect(win.History().Length()).To(Equal(1))
		})

		It("Should have a length of two when navigating", func() {
			Expect(win.Navigate("/page-2")).To(Succeed())
			Expect(win.History().Length()).To(Equal(2))
		})

		It("Should reload, but keep the length on Go(0)", func() {
			Expect(win.Navigate("/page-2")).To(Succeed())
			Expect(win.History().Length()).To(Equal(2))
			Expect(h.RequestCount()).To(Equal(1)) // about:blank wasn't a request
			Expect(win.History().Go(0)).To(Succeed())
			Expect(win.History().Length()).To(Equal(2))
			Expect(h.RequestCount()).To(Equal(2)) // about:blank wasn't a request
		})

		It("Should have a length of two when navigating", func() {
			Expect(win.Navigate("/page-2")).To(Succeed())
			Expect(win.History().Length()).To(Equal(2))
			Expect(win.Document()).To(HaveH1("/page-2"))
		})

		It("Should go back, but keep the length", func() {
			Expect(win.Navigate("/page-2")).To(Succeed())
			Expect(win.History().Go(-1)).To(Succeed())
			Expect(win.Document()).To(HaveH1("Gost-DOM"))
			Expect(win.Location().Href()).To(Equal("about:blank"))
		})

		It("Should truncate history when going forward", func() {
			Expect(win.Navigate("/page-2")).To(Succeed())
			Expect(win.Navigate("/page-3")).To(Succeed())
			Expect(win.Navigate("/page-4")).To(Succeed())
			Expect(win.Navigate("/page-5")).To(Succeed())
			Expect(win.History().Length()).To(Equal(5))
			Expect(win.History().Go(-3)).To(Succeed())
			Expect(win.History().Length()).To(Equal(5))
			Expect(win.Navigate("/page-6")).To(Succeed())
			Expect(win.History().Length()).To(Equal(3))
			Expect(win.Location().Pathname()).To(Equal("/page-6"))
		})

		Describe("popstate event", func() {
			// Initial state is about:blank
			BeforeEach(func() {
				Expect(win.Navigate("/page-2")).To(Succeed())
			})

			It("Should do what when the target entry has a state, but was reloaded?", func() {
				Skip("Research")
			})

			Describe("Call replaceState with state, then pustState", func() {
				BeforeEach(func() {
					Expect(win.History().ReplaceState("page-2 state", "")).To(Succeed())
					Expect(win.History().PushState(EMPTY_STATE, "/page-3")).To(Succeed())
				})

				It("Should dispatch a popstate event with the state", func() {
					var actualEvent dom.Event
					win.AddEventListener(
						"popstate",
						dom.NewEventHandlerFunc(func(e dom.Event) error {
							actualEvent = e
							return nil
						}),
					)
					Expect(win.History().Go(-1)).To(Succeed())

					Expect(actualEvent).ToNot(BeNil(), "Event was dispatched")
					popEvent, ok := actualEvent.(PopStateEvent)
					Expect(ok).To(BeTrue(), "Event is a popstateevent")
					Expect(popEvent.State()).To(BeEquivalentTo("page-2 state"), "Event state")
				})
			})
		})

		It("Go should return a Security error if the document is not fully active", func() {
			Skip("TODO")
			/*
			   https://developer.mozilla.org/en-US/docs/Web/API/History/go

			   Thrown if the associated document is not fully active, or if the
			   provided url parameter is not a valid URL. Browsers also throttle navigations
			   and may throw this error, generate a warning, or ignore the call if it's called
			   too frequently
			*/
		})

		Describe("ReplaceState", func() {
			It("Should change 'location' but keep stack count", func() {
				Expect(win.Navigate("/page-2")).To(Succeed())
				Expect(win.History().Length()).To(Equal(2))
				Expect(win.History().ReplaceState(EMPTY_STATE, "/page-3"))
				Expect(h.RequestCount()).To(Equal(1))
				Expect(win.History().Length()).To(Equal(2))
				Expect(win.Location().Pathname()).To(Equal("/page-3"))
			})

			It("Should return an error on different origin", func() {
				Skip("TODO")
			})

			It("Should keep the same URL when the href is empty", func() {
				Skip("TODO")
			})
		})

		Describe("PushState", func() {
			It("Should return a Security error if the document is not fully active", func() {
				Skip("TODO")
				/*
				   https://developer.mozilla.org/en-US/docs/Web/API/History/pushState

				   Thrown if the associated document is not fully active, or if the
				   provided url parameter is not a valid URL. Browsers also throttle navigations
				   and may throw this error, generate a warning, or ignore the call if it's called
				   too frequently
				*/
			})

			Describe("Simple push and back", func() {
				BeforeEach(func() {
					Expect(win.Navigate("/page-2")).To(Succeed())
					Expect(win.Navigate("/page-3")).To(Succeed())

					Expect(win.History().Length()).To(Equal(3))
					Expect(win.History().PushState(EMPTY_STATE, "/page-4"))
				})

				It("Should not emit a hashchange event when just adding a hash", func() {
					eventDispatched := false
					win.AddEventListener(
						"hashchange",
						dom.NewEventHandlerFunc(func(e dom.Event) error {
							eventDispatched = true
							return nil
						}),
					)
					Expect(win.History().PushState(EMPTY_STATE, "/page-4#target"))
					Expect(eventDispatched).To(BeFalse())
				})

				It("Should change 'location' and increase stack count, without a request", func() {
					Expect(win.History().Length()).To(Equal(4))
					Expect(win.History().Length()).To(Equal(4))
					Expect(win.Location().Pathname()).To(Equal("/page-4"))
					Expect(h.RequestCount()).To(Equal(2), "No of request _after_ replaceState")
				})

				It("Navigates back without a request", func() {
					Expect(win.History().Back()).To(Succeed())
					Expect(win.History().Length()).To(Equal(4))
					Expect(win.Location().Pathname()).To(Equal("/page-3"))
					Expect(h.RequestCount()).To(Equal(2), "No of request _after_ back")
				})

				It("Should push the current URL if called with empty URL", func() {
					Skip("TODO")
				})
			})

			Describe("History with multiple pushState and navigation intermixed", func() {
				BeforeEach(func() {
					Expect(win.Navigate("/page-2")).To(Succeed())
					Expect(win.Navigate("/page-3")).To(Succeed())
					Expect(win.History().Length()).To(Equal(3))
					Expect(win.History().PushState(EMPTY_STATE, "/page-4"))
					Expect(win.History().PushState(EMPTY_STATE, "/page-5"))
					Expect(win.Navigate("/page-6")).To(Succeed())
					Expect(win.Navigate("/page-7")).To(Succeed())
					Expect(win.History().PushState(EMPTY_STATE, "/page-8"))
					Expect(win.History().PushState(EMPTY_STATE, "/page-9"))
					Expect(win.History().Length()).To(Equal(9))
					Expect(h.RequestCount()).To(Equal(4))
				})

				It("Should not issue an HTTP request on go(-2)", func() {
					Expect(win.History().Go(-2)).To(Succeed())
					Expect(win.History().Length()).To(Equal(9))
					Expect(h.RequestCount()).To(Equal(4))
					Expect(win.Document().GetElementById("heading")).To(HaveTextContent("/page-7"))
				})

				It("Should issue an HTTP request on go(-3)", func() {
					Expect(win.History().Go(-3)).To(Succeed())
					Expect(win.History().Length()).To(Equal(9))
					Expect(h.RequestCount()).To(Equal(5))
					Expect(win.Document().GetElementById("heading")).To(HaveTextContent("/page-6"))
				})

				Describe("After history.go(-5)", func() {
					BeforeEach(func() {
						Expect(win.History().Go(-5)).To(Succeed())
					})

					It("Should have loaded page 4", func() {
						Expect(
							win.Document().GetElementById("heading"),
						).To(HaveTextContent("/page-4"))
					})

					It("Should have issued a new request", func() {
						Expect(win.History().Length()).To(Equal(9))
						Expect(h.RequestCount()).To(Equal(5))
					})

					It("Should not issue a new request on go(1)", func() {
						Expect(win.History().Go(1)).To(Succeed())
						Expect(win.History().Length()).To(Equal(9))
						Expect(h.RequestCount()).To(Equal(5))
					})

					It("Should issue a new request on go(2)", func() {
						Expect(win.History().Go(2)).To(Succeed())
						Expect(win.History().Length()).To(Equal(9))
						Expect(h.RequestCount()).To(Equal(6))
						Expect(
							win.Document().GetElementById("heading"),
						).To(HaveTextContent("/page-6"))
					})
				})
			})
		})
	})

	Describe("Location()", func() {
		var window Window

		BeforeEach(func() {
			server := testservers.NewAnchorTagNavigationServer()
			DeferCleanup(func() { server = nil })
			window = NewWindowFromHandler(server)
		})

		It("Should be about:blank", func() {
			Expect(window.Location().Href()).To(Equal("about:blank"))
		})

		It("Should return the path loaded from", func() {
			Expect(window.Navigate("/index")).To(Succeed())
			Expect(window.Location().Pathname()).To(Equal("/index"))
		})

		Describe("Navigate", func() {
			It("Should load a blank page when loading about:blank", func() {
				Expect(window.Navigate("about:blank")).To(Succeed())
				Expect(window.Document()).To(HaveH1("Gost-DOM"))
			})
		})

		Describe("User navigation (clicking links)", func() {
			var links []dom.Node

			BeforeEach(func() {
				Expect(window.Navigate("/index")).To(Succeed())
				nodes, err := window.Document().QuerySelectorAll("a")
				Expect(err).ToNot(HaveOccurred())
				links = nodes.All()
			})

			It("Should update when using a link with absolute url", func() {
				link, ok := domslices.SliceFindFunc(links, func(n dom.Node) bool {
					return n.TextContent() == "Products from absolute url"
				})
				Expect(ok).To(BeTrue())
				link.(dom.Element).Click()
				Expect(window.Location().Pathname()).To(Equal("/products"))
			})

			It("Should update when using a link with relative url", func() {
				link, ok := domslices.SliceFindFunc(links, func(n dom.Node) bool {
					return n.TextContent() == "Products from relative url"
				})
				Expect(ok).To(BeTrue())
				link.(dom.Element).Click()
				Expect(window.Location().Pathname()).To(Equal("/products"))
			})
		})
	})
})
