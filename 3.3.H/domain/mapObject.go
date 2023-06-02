package domain

type MapObject interface {
	GetSymbol() string
	Touched() State
	GetType() MapObjectType
	GetPosition() Position
	SetPosition(Position)
}

type MapObjectCore struct {
	Map      *Map
	Symbol   string
	Type     MapObjectType
	Position Position
}

func (m MapObjectCore) GetPosition() Position {
	return m.Position
}

func (m MapObjectCore) GetSymbol() string {
	return m.Symbol
}

func (m MapObjectCore) GetType() MapObjectType {
	return m.Type
}

func (m MapObjectCore) Touched() State {
	return nil
}

func (m *MapObjectCore) SetPosition(p Position) {
	m.Position = p
}

type MapObjectType string

const (
	_Organism MapObjectType = "Organism"
	_Treasure MapObjectType = "Treasure"
	_Obstacle MapObjectType = "Obstacle"
)
