//go:generate ../../code-gen/code-gen -g html-elements -o html_elements.go
//go:generate ../../code-gen/code-gen -g scripting -o dummy

// The scripting package implements ECMAScript execution through the v8 engine.
package scripting
