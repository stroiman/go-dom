package html

// Will eventually contain more information, e.g. state
type historyEntry struct {
	href string
}

// History implements the [History API]
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
	newPos := h.currentPos + relative
	if newPos <= 0 || newPos > h.Length() {
		return nil
	}
	h.currentPos = newPos
	return h.window.reload(h.entries[h.currentPos-1].href)
}

func (h *History) pushLoad(href string) {
	if h.currentPos < h.Length() {
		h.entries = h.entries[0:h.currentPos]
	}
	h.currentPos++
	h.entries = append(h.entries, historyEntry{href: href})
}
