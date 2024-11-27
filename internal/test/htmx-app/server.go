package app

import (
	"fmt"
	"net/http"

	"github.com/stroiman/go-dom/internal/test/htmx-app/content"
)

func CreateServer() http.Handler {
	server := http.NewServeMux()
	count := 1
	server.Handle("GET /", http.FileServer(http.FS(content.FS)))
	server.HandleFunc("POST /increment", func(res http.ResponseWriter, req *http.Request) {
		count++
		res.Write([]byte(fmt.Sprintf("Count: %d", count)))
	})
	return server
}
