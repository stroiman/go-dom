//go:generate ../../internal/code-gen/code-gen -g scripting

// The v8host packages provides functionality to execute client-side scripts in
// gost-dom. The engine uses v8, and requires cgo.
//
// This engine is based on tommie's v8go form, which automatically pulls the
// latest v8 changes from the chromium repo.
//
// An alternate script engine that is implemented in pure go is found in the
// gojahost package
//
// See also: https://github.com/tommie/v8go
package v8host
