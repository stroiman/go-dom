package app

import (
	"net/http"

	"github.com/stroiman/go-dom/internal/test/htmx-app/content"
)

func CreateServer() http.Handler {
	server := http.NewServeMux()
	server.Handle("/", http.FileServer(http.FS(content.FS)))
	return server
}
