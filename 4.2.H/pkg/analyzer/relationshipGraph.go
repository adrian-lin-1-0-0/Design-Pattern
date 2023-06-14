package analyzer

type RelationshipGraph interface {
	HasConnection(name1 string, name2 string) bool
}
