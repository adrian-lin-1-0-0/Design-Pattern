package analyzer

type RelationshipAnalyzer interface {
	Init(script string) RelationshipGraph
	GetMutualFriends(name1 string, name2 string) []string
}

type RelationshipGraph interface {
	HasConnection(name1 string, name2 string) bool
}
