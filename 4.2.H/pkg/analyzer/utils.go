package analyzer

type Node struct {
	id int64
}

func NewNode(id int64) *Node {
	return &Node{id: id}
}

func (n *Node) ID() int64 {
	return n.id
}

type idCreator struct {
	id int64
}

func NewIdCreator() *idCreator {
	return &idCreator{id: 0}
}

func (creator *idCreator) Create() int64 {
	creator.id++
	return creator.id
}
