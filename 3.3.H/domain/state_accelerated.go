package domain

type Accelerated struct {
	*StateCore
	TTL int
}

func NewAccelerated() *Accelerated {
	return &Accelerated{
		StateCore: &StateCore{
			OrganismCore: nil,
			name:         "Accelerated",
		},
		TTL: 2}
}

func (a *Accelerated) Effect(o *OrganismCore) {
	a.StateCore.Effect(o)
	a.OrganismCore.ActionCount = 2
}

func (a *Accelerated) TakeTurn() {
	a.StateCore.TakeTurn()
	a.TTL--
	if a.TTL == 0 {
		a.OrganismCore.ResetState()
	}
}

func (a *Accelerated) Injured(power int) {
	a.StateCore.Injured(power)
	a.OrganismCore.ResetState()
}
