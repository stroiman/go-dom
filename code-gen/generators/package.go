// Package generators provides general code generation functionality. This is
// primarily a more composable API on top of the [jen] package.
//
// It also provides some usefule helper functions for the purpose, such as
// generating a default receiver variable name for methods on a specific type.
//
// WARNING: API is not stable. This package contains exported names, and can be
// used on its own, but exists solely for this project; and as such, they API
// may change.
package generators
