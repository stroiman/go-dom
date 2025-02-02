package app

import (
	"fmt"
	"net/http"

	"github.com/gost-dom/browser/internal/test/htmx-app/content"
)

type TestServer struct {
	Mux      *http.ServeMux
	Requests []*http.Request
}

func (s *TestServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	s.Requests = append(s.Requests, r)
	s.Mux.ServeHTTP(w, r)
}

func CreateServer() *TestServer {
	mux := http.NewServeMux()
	count := 1
	mux.Handle("GET /", http.FileServer(http.FS(content.FS)))
	mux.HandleFunc("POST /counter/increment", func(w http.ResponseWriter, r *http.Request) {
		count++
		w.Write([]byte(fmt.Sprintf("Count: %d", count)))
	})
	mux.HandleFunc("POST /forms/form-1", func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		fmt.Println("*** FORM", r.Form, r.Header)
		w.Write([]byte("Form values:<br />"))
		w.Write([]byte(fmt.Sprintf(`<div id="field-value-1">%s</div>`, r.FormValue("field-1"))))
	})
	return &TestServer{Mux: mux}
}
