package domain

type Poisoned struct {
	*StateCore
	TTL int
}

func NewPoisoned() *Poisoned {
	return &Poisoned{
		StateCore: &StateCore{
			OrganismCore: nil,
			name:         "Poisoned",
		},
		TTL: 3}
}

func (p *Poisoned) TakeTurn() {
	p.OrganismCore.Injured(10)
	p.StateCore.TakeTurn()
	p.TTL--
	if p.TTL == 0 {
		p.OrganismCore.ResetState()
		return
	}
}
