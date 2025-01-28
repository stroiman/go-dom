package main

import (
	"net/http"

	app "github.com/gost-dom/browser/browser/internal/test/htmx-app"
)

func main() {
	server := app.CreateServer()
	http.ListenAndServe(":4000", server)
}
