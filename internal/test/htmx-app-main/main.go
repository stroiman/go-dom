package main

import (
	"net/http"

	app "github.com/stroiman/go-dom/internal/test/htmx-app"
)

func main() {
	server := app.CreateServer()
	http.ListenAndServe(":4000", server)
}
