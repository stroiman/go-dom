package testing

import (
	"fmt"
	"net/http"
)

// EchoHandler is an [http.Handler] that echos the requested path in the
// document heading, i.e., the <h1> element.
type EchoHandler struct {
	// Requests contains all HTTP requests sent to the handler
	Requests []*http.Request
}

func (h *EchoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.Requests = append(h.Requests, r)
	w.Write([]byte(fmt.Sprintf("<body><h1>%s</h1></body>", r.URL.Path)))
}

// RequestCount returns how many HTTP requests have been made to the handler
func (h *EchoHandler) RequestCount() int {
	return len(h.Requests)
}
