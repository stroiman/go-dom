package dom_types

type Element struct {
	tagName string
}

func NewElement(tagName string) Element { return Element{tagName} }

func (e HTMLElement) NodeName() string {
	return e.TagName()
}

func (e HTMLElement) TagName() string {
	return e.tagName
}
