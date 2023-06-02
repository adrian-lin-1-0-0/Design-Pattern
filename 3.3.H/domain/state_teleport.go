package domain

import "math/rand"

type Teleport struct {
	*StateCore
	TTL int
}

func NewTeleport() *Teleport {
	return &Teleport{
		StateCore: &StateCore{
			OrganismCore: nil,
			name:         "Teleport",
		},
		TTL: 1,
	}
}

func (t *Teleport) TakeTurn() {
	t.StateCore.TakeTurn()
	t.TTL--
	if t.TTL == 0 {
		t.OrganismCore.SetState(t.OrganismCore.DefaultState)
		t.MoveToRandom()
	}
}

func (t *Teleport) MoveToRandom() {
	emptyPositions := t.OrganismCore.MapObjectCore.Map.GetEmpty()
	t.OrganismCore.MoveTo(emptyPositions[rand.Intn(len(emptyPositions))])
}
