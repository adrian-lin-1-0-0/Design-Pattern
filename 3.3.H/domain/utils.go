package domain

func GetCross(p Position) []Position {
	return []Position{
		{X: p.X, Y: p.Y - 1},
		{X: p.X + 1, Y: p.Y},
		{X: p.X, Y: p.Y + 1},
		{X: p.X - 1, Y: p.Y},
	}
}

func IsCharacter(mapObject MapObject) bool {
	if mapObject == nil {
		return false
	}
	return mapObject.GetType() == _Organism && mapObject.(Organism).GetOrganismType() == _Character
}

func IsMonster(mapObject MapObject) bool {
	if mapObject == nil {
		return false
	}
	return mapObject.GetType() == _Organism && mapObject.(Organism).GetOrganismType() == _Monster
}

func IsOrganism(mapObject MapObject) bool {
	if mapObject == nil {
		return false
	}
	return mapObject.GetType() == _Organism
}
