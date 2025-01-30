package domslices

import "slices"

// SliceFindFunc finds a the first element in a slice for which function f
// returns true. If not found ok will be false.
func SliceFindFunc[T any](s []T, f func(T) bool) (e T, ok bool) {
	idx := slices.IndexFunc(s, f)
	if ok = idx >= 0; ok {
		e = s[idx]
	}
	return
}
