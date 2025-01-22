package html

// Will eventually contain more information, e.g. state
type historyEntry struct {
	url string
}

type History struct {
	entries []historyEntry
}

func (h History) Length() int {
	return len(h.entries)
}

func (h *History) pushLoad(url string) {
	h.entries = append(h.entries, historyEntry{url: url})
}
