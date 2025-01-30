package html

import "github.com/gost-dom/browser/dom"

const HistoryEventPopState = "popstate"

// The PopStateEvent is emitted after navigating to the same document, and will
// contain possible state passed to [History.ReplaceState] or
// [History.PushState].
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/PopStateEvent
type PopStateEvent interface {
	dom.Event
	State() HistoryState
}

// Will eventually contain more information, e.g. state
type historyEntry struct {
	// remote is true when the history entry was caused by a normal navigation,
	// e.g., clicking a link. remote is false when a history entry was added
	// through [pushState]
	remote bool
	state  HistoryState
	href   string
}

// History implements the [History API]
//
// Note: Currently, state is represented by string, but may change to a
// different type representing a JSON object.
//
// [History API]: https://developer.mozilla.org/en-US/docs/Web/API/History
type History struct {
	window     *window
	entries    []historyEntry
	currentPos int
}

// Length returns the number of entries in the history. When navigating back,
// the length doesn't change, as they last viewed page is still in history.
//
// Navigating to a new location truncates future history later than the current
// page.
func (h History) Length() int {
	return len(h.entries)
}

// Back calls Go(-1).
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/History/back
func (h *History) Back() error { return h.Go(-1) }

// Forward calls Go(1).
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/History/forward
func (h *History) Forward() error { return h.Go(1) }

// Go moves back or forward through the history, possibly reloading the page if
// necessary. A negative value goes back in history; a positive value moves
// forward if possible. A value of 0 will trigger a reload.
//
// See also: https://developer.mozilla.org/en-US/docs/Web/API/History/go
func (h *History) Go(relative int) error {
	prevEntry := h.entries[h.currentIdx()]
	if relative == 0 {
		return h.window.reload(prevEntry.href)
	}

	shouldReload := false
	newPos := h.currentPos + relative
	if newPos <= 0 || newPos > h.Length() {
		return nil
	}
	if relative < 0 { // go back
		for i := h.currentPos; i > newPos; i-- {
			shouldReload = shouldReload || h.entries[i-1].remote
		}
	} else { // go forward
		for i := h.currentPos; i < newPos; i++ {
			shouldReload = shouldReload || h.entries[i].remote
		}
	}

	h.currentPos = newPos
	newCurrentEntry := h.entries[h.currentIdx()]
	newHref := newCurrentEntry.href
	if shouldReload {
		return h.window.reload(newHref)
	} else {
		h.window.baseLocation = newHref
		h.window.DispatchEvent(newPopStateEvent(newCurrentEntry.state))
		return nil
	}
}

// ReplaceState the current [Window.Location] and history entry without making a
// new request.
//
// The function corresponds to [replaceState on the History API] with the
// following notes. If href is empty, the URL will not be updated; as if the
// argument was not specified in the JS API.
//
// [replaceState on the History API]: https://developer.mozilla.org/en-US/docs/Web/API/History/replaceState
func (h *History) ReplaceState(state HistoryState, href string) error {
	idx := h.currentIdx()
	h.entries[idx].href = h.window.setBaseLocation(href)
	h.entries[idx].state = state
	return nil
}

func (h History) currentIdx() int {
	return h.currentPos - 1
}

func (h *History) PushState(state HistoryState, href string) error {
	newHref := h.window.setBaseLocation(href)
	h.pushHistoryEntry(historyEntry{state: state, href: newHref, remote: false})
	return nil
}

func (h *History) pushHistoryEntry(entry historyEntry) {
	if h.currentPos < h.Length() {
		h.entries = h.entries[0:h.currentPos]
	}
	h.currentPos++
	h.entries = append(h.entries, entry)
}

func (h *History) pushLoad(href string) {
	entry := historyEntry{href: href, remote: true}
	h.pushHistoryEntry(entry)
}

func (h History) State() HistoryState {
	return h.entries[h.currentIdx()].state
}

type popStateEvent struct {
	dom.Event
	state HistoryState
}

func newPopStateEvent(state HistoryState) PopStateEvent {
	return popStateEvent{dom.NewEvent(HistoryEventPopState), state}
}

func (e popStateEvent) State() HistoryState {
	return e.state
}

type HistoryState string

const EMPTY_STATE HistoryState = ""
