package domain

type MapObject interface {
	GetSymbol() string
	Touched() State
	GetType() MapObjectType
}

type MapObjectCore struct {
	Map    *Map
	Symbol string
	Type   MapObjectType
}

type MapObjectType string

const (
	_Monster   MapObjectType = "Monster"
	_Character MapObjectType = "Character"
	_Treasure  MapObjectType = "Treasure"
	_Obstacle  MapObjectType = "Obstacle"
)
