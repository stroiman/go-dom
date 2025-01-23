package html

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
// Note: Currently, state is represented by an interface{} type. THIS MAY
// CHANGE. The data must be JSON serializable data, and once stored, it must not
// be changed. Another datatype will eventually be used.
//
// [History API]: https://developer.mozilla.org/en-US/docs/Web/API/History
type History struct {
	window     *window
	entries    []historyEntry
	currentPos int
}

type HistoryState = interface{}

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
	prevEntry := h.entries[h.currentPos-1]
	if relative == 0 {
		return h.window.reload(prevEntry.href)
	}

	shouldReload := false
	newPos := h.currentPos + relative
	if newPos <= 0 || newPos > h.Length() {
		return nil
	}
	if relative < 0 {
		// go back
		for i := h.currentPos; i > newPos; i-- {
			shouldReload = shouldReload || h.entries[i-1].remote
		}
	} else {
		// go forward
		for i := h.currentPos; i < newPos; i++ {
			shouldReload = shouldReload || h.entries[i].remote
		}
	}

	h.currentPos = newPos
	newCurrentEntry := h.entries[h.currentPos-1]
	newHref := newCurrentEntry.href
	if shouldReload {
		return h.window.reload(newHref)
	} else {
		h.window.baseLocation = newHref
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
func (h *History) ReplaceState(state interface{}, href string) error {
	idx := h.currentPos - 1
	if href != "" {
		newHref := h.window.setBaseLocation(href)
		h.entries[idx].href = newHref
	}
	h.entries[idx].state = state
	return nil
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
