// Package browser is the main entry point for Gost, helping create a window
// initialized with a script enging, connected to a server.
//
// Important!
//
// This package depends on two 3rd party components that needs some custom
// modifications to work.
//
//	go mod edit -replace="github.com/tommie/v8go=github.com/stroiman/v8go@go-dom-support"
//	go mod edit -replace="github.com/ericchiang/css=github.com/gost-dom/css@latest"
//	go mod tidy
//
// I hope that all my changes will make it to the original repos, eliminating
// the need for replace (or maintaining a new set of forks).
package browser
