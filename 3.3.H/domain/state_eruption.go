package domain

type Eruption struct {
	*StateCore
	TTL int
}

func NewEruption() *Eruption {
	return &Eruption{
		StateCore: &StateCore{
			OrganismCore: nil,
			name:         "Eruption",
		},
		TTL: 3,
	}
}

func (e *Eruption) TakeTurn() {
	e.StateCore.TakeTurn()
	e.TTL--
	if e.TTL == 0 {
		e.OrganismCore.SetState(NewTeleport())
		return
	}
}

func (e *Eruption) Effect(o *OrganismCore) {
	e.StateCore.Effect(o)
	e.OrganismCore.AttackPower = 50
}

func (e *Eruption) Attack() error {
	for _, mapObject := range e.OrganismCore.MapObjectCore.Map.GetAllObjects() {
		if mapObject.GetPosition() != e.OrganismCore.GetPosition() {
			if mapObject.GetType() == _Organism {
				mapObject.(Organism).Injured(e.OrganismCore.AttackPower)
			}
		}
	}
	return nil
}
