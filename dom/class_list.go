package dom

import (
	"iter"
	"slices"
	"strings"
)

type DOMTokenList struct {
	element Element
}

func NewClassList(element Element) DOMTokenList {
	return DOMTokenList{element}
}

func (l DOMTokenList) Add(tokens ...string) error {
	classes := l.getTokens()
	for _, token := range tokens {
		if token == "" {
			return newDomErrorCode("Empty token", domErrorSyntaxError)
		}
		if strings.Contains(token, " ") {
			return newDomErrorCode("Empty token", domErrorInvalidCharacter)
		}
		if !slices.Contains(classes, token) {
			classes = append(classes, token)
		}
	}
	l.setTokens(classes)
	return nil
}

func (l DOMTokenList) All() iter.Seq[string] {
	return func(yield func(string) bool) {
		tokens := l.getTokens()
		for _, token := range tokens {
			if !yield(token) {
				return
			}
		}
	}
}

func (l DOMTokenList) Contains(token string) bool {
	classes := l.getTokens()
	return slices.Contains(classes, token)
}

func (l DOMTokenList) Length() int {
	return len(l.getTokens())
}

func (l DOMTokenList) Value() string {
	a, _ := l.element.GetAttribute("class")
	return a
}

func (l DOMTokenList) SetValue(val string) {
	l.element.SetAttribute("class", val)
}

func (l DOMTokenList) Item(index int) *string {
	classes := l.getTokens()
	if index >= len(classes) {
		return nil
	}
	return &classes[index]
}

func (l DOMTokenList) Remove(token string) {
	tokens := l.getTokens()
	itemIndex := slices.Index(tokens, token)
	if itemIndex >= 0 {
		newList := slices.Delete(tokens, itemIndex, itemIndex+1)
		l.setTokens(newList)
	}
}

func (l DOMTokenList) Replace(oldToken string, newToken string) bool {
	if l.Contains(oldToken) {
		l.Remove(oldToken)
		l.Add(newToken)
		return true
	} else {
		return false
	}
}

func (l DOMTokenList) Toggle(token string) bool {
	if l.Contains(token) {
		l.Remove(token)
		return false
	} else {
		l.Add(token)
		return true
	}
}

func (l DOMTokenList) getTokens() []string {
	class, found := l.element.GetAttribute("class")
	if !found {
		return []string{}
	}
	return strings.Split(class, " ")
}

func (l DOMTokenList) setTokens(tokens []string) {
	l.element.SetAttribute("class", strings.Join(tokens, " "))
}
