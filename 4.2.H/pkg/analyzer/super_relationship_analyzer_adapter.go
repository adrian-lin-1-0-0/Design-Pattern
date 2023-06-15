package analyzer

import (
	"fmt"
	"strings"
)

type SuperRelationshipAnalyzerAdapter struct {
	analyzer *SuperRelationshipAnalyzer
	edges    map[string][]string
}

func NewSuperRelationshipAnalyzerAdapter() *SuperRelationshipAnalyzerAdapter {
	return &SuperRelationshipAnalyzerAdapter{
		analyzer: NewSuperRelationshipAnalyzer(),
		edges:    make(map[string][]string),
	}
}

/*
*
A: B C D
B: A D E
C: A E G K M
D: A B K P
E: B C J K L
F: Z
*/
func (sa *SuperRelationshipAnalyzerAdapter) Init(script string) RelationshipGraph {
	newScript := sa.striptTransfer(script)
	sa.analyzer.Init(newScript)
	return NewGonumAdapter(sa.analyzer.graph, sa.analyzer.Nodes)
}

func (sa *SuperRelationshipAnalyzerAdapter) GetMutualFriends(name1 string, name2 string) []string {
	mutualFriends := []string{}
	for key := range sa.edges {
		if key == name1 || key == name2 {
			continue
		}
		if sa.analyzer.IsMutualFriend(name1, name2, key) {
			mutualFriends = append(mutualFriends, key)
		}
	}
	return mutualFriends
}

func (sa *SuperRelationshipAnalyzerAdapter) striptTransfer(script string) string {

	lines := strings.Split(strings.TrimSpace(script), "\n")
	for _, line := range lines {
		parts := strings.Split(line, ":")
		node := strings.TrimSpace(parts[0])
		neighbors := strings.Fields(strings.TrimSpace(parts[1]))

		sa.edges[node] = append(sa.edges[node], neighbors...)
		for _, neighbor := range neighbors {
			sa.edges[neighbor] = append(sa.edges[neighbor], node)
		}
	}

	newString := ""
	for node, neighbors := range sa.edges {
		for _, neighbor := range neighbors {
			newString += fmt.Sprintf("%s -- %s\n", node, neighbor)
		}
	}
	return newString
}
