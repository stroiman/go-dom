//go:generate ../../code-gen/code-gen -g html-elements -o html_elements.go
//go:generate ../../code-gen/code-gen -g dom -o dom_generated.go
//go:generate ../../code-gen/code-gen -g xhr -o xml_http_request_generated.go
//go:generate ../../code-gen/code-gen -g url -o url_generated.go

// The scripting package implements ECMAScript execution through the v8 engine.
package scripting
