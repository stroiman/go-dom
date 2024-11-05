package dom_types

type Element struct {
	tagName string
}

func NewElement(tagName string) Element { return Element{tagName} }

func (e Element) NodeName() string {
	return e.TagName()
}

func (e Element) TagName() string {
	return e.tagName
}
