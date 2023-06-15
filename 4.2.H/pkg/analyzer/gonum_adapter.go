package analyzer

import (
	"gonum.org/v1/gonum/graph"
	"gonum.org/v1/gonum/graph/simple"
	"gonum.org/v1/gonum/graph/traverse"
)

type GonumAdapter struct {
	graph *simple.UndirectedGraph
	nodes map[string]*Node
}

func NewGonumAdapter(
	graph *simple.UndirectedGraph,
	nodes map[string]*Node) *GonumAdapter {
	return &GonumAdapter{
		graph: graph,
		nodes: nodes,
	}
}

func (adapter *GonumAdapter) HasConnection(name1 string, name2 string) bool {
	node1 := adapter.nodes[name1]
	node2 := adapter.nodes[name2]
	hasConnection := false
	path := traverse.BreadthFirst{
		Visit: func(n graph.Node) {
			if n.ID() == node2.ID() {
				hasConnection = true
			}
		},
	}

	path.Walk(adapter.graph, node1, func(n graph.Node, d int) bool {
		return n.ID() == node2.ID()
	})
	return hasConnection
}
