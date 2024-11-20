package browser

import "github.com/ericchiang/css"

type CSSHelper struct{ Node }

func (h CSSHelper) QuerySelector(pattern string) (Node, error) {
	nodes, err := h.QuerySelectorAll(pattern)
	var result Node
	if len(nodes) > 0 {
		result = nodes[0]
	}
	return result, err
}

func (d CSSHelper) QuerySelectorAll(pattern string) (StaticNodeList, error) {
	sel, err := css.Parse(pattern)
	if err != nil {
		return nil, err
	}
	htmlNode, m := toHtmlNodeAndMap(d)
	nodes := sel.Select(htmlNode)
	result := make(StaticNodeList, len(nodes))
	for i, node := range nodes {
		resultNode := m[node]
		result[i] = resultNode
	}
	return result, nil
}
