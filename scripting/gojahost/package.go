//go:generate ../../../code-gen/code-gen -g goja

// The gojahost package provides functionality to execute client-scripts in
// gost-dom. The package is a pure Go implementation by using goja.
//
// The package may not have full ECMAScript support. If this is a problem, the
// [v8host] proves a v8 engine implementation, but requires CGo.
//
// See also: https://pkg.go.dev/github.com/dop251/goja
package gojahost
