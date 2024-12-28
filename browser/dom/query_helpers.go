package dom

import "github.com/ericchiang/css"

type CSSHelper struct{ Node }

func (h CSSHelper) QuerySelector(pattern string) (Node, error) {
	nodes, err := h.QuerySelectorAll(pattern)
	return nodes.Item(0), err
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
