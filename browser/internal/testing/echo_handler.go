package testing

import (
	"fmt"
	"net/http"
)

// EchoHandler is an [http.Handler] that echos the requested path in the
// document heading, i.e., the <h1> element.
var EchoHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("<body><h1>%s</h1></body>", r.URL.Path)))
})
