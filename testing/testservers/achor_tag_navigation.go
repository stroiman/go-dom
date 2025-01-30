package testservers

import "net/http"

// Simple test server to help verify havigation
func NewAnchorTagNavigationServer() http.Handler {
	server := http.NewServeMux()
	server.HandleFunc("GET /index",
		func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(
				`<body>
					<h1>Index</h1>
					<a href="products">Products from relative url</a>
					<a href="/products">Products from absolute url</a>
				</body>`))
		})

	server.HandleFunc("GET /products",
		func(res http.ResponseWriter, req *http.Request) {
			res.Write([]byte(
				`<body>
					<h1>Products</h1>
				</body>`))
		})

	return server
}
