package dom

import (
	"github.com/ericchiang/css"
)

type CSSHelper struct{ Node }

func (h CSSHelper) QuerySelector(pattern string) (Element, error) {
	nodes, err := h.QuerySelectorAll(pattern)
	if err != nil {
		return nil, err
	}
	// TODO, it should be a list of elements, not nodes, then the cast, and
	// error isn't necessary
	result := nodes.Item(0)
	element, _ := result.(Element)
	return element, nil
}

func (d CSSHelper) QuerySelectorAll(pattern string) (StaticNodeList, error) {
	sel, err := css.Parse(pattern)
	if err != nil {
		return nil, err
	}
	htmlNode, m := toHtmlNodeAndMap(d)
	nodes := sel.Select(htmlNode)
	result := make([]Node, len(nodes))
	for i, node := range nodes {
		resultNode := m[node]
		result[i] = resultNode
	}
	return NewNodeList(result...), nil
}
