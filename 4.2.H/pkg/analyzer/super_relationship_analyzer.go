package analyzer

import (
	"bufio"
	"strings"

	"gonum.org/v1/gonum/graph/simple"
)

type SuperRelationshipAnalyzer struct {
	Nodes map[string]*Node
	graph *simple.UndirectedGraph
}

func NewSuperRelationshipAnalyzer() *SuperRelationshipAnalyzer {
	return &SuperRelationshipAnalyzer{
		Nodes: make(map[string]*Node),
		graph: simple.NewUndirectedGraph(),
	}
}

/*
*
/ script:
A -- B
A -- C
A -- D
B -- D
B -- E
C -- E
C -- G
C -- K
C -- M
D -- K
D -- P
E -- J
E -- K
E -- L
F -- Z
*/
func (analyzer *SuperRelationshipAnalyzer) Init(script string) {
	reader := strings.NewReader(script)

	scanner := bufio.NewScanner(reader)
	idCreator := NewIdCreator()

	for scanner.Scan() {
		line := scanner.Text()
		names := strings.Split(line, " -- ")
		for _, name := range names {
			if _, ok := analyzer.Nodes[name]; !ok {
				analyzer.Nodes[name] = NewNode(idCreator.Create())
				analyzer.graph.AddNode(analyzer.Nodes[name])
			}
		}
		for _, name := range names {
			for _, otherName := range names {
				if name == otherName {
					continue
				}
				analyzer.graph.SetEdge(analyzer.graph.NewEdge(analyzer.Nodes[name], analyzer.Nodes[otherName]))
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}
}

func (analyzer *SuperRelationshipAnalyzer) IsMutualFriend(targetName string, name2 string, name3 string) bool {
	p1 := analyzer.Nodes[targetName]
	p2 := analyzer.Nodes[name2]
	p3 := analyzer.Nodes[name3]
	return analyzer.graph.HasEdgeBetween(p1.ID(), p2.ID()) && analyzer.graph.HasEdgeBetween(p1.ID(), p3.ID())
}
