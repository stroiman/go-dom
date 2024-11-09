package dom_types

// This doesn't correspond to a specific type in the DOM spec, but both a
// Document and Element are _nodes_, and add children, append, remove, etc,
// functions that work on _element_ instead of _node_ objects
type ElementCollection struct {
	*Node
}
