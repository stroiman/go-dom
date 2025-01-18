package dom

/*
TODO:
ancestorOrigins
hash
*/

// Location is the interface of window.Location and document.Location (they return
// the same object).
//
// The interface itself is a superset of the URL, but has setters that when set,
// has the side effect of the browser actually navigating.
//
// The setters are not yet implemented, so the type is just an embedding of the
// URL interface
type Location interface {
	URL
}
