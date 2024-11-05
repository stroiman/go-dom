package dom_types

type Document struct{}

func (d Document) NodeName() string { return "#document" }
