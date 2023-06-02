package domain

type Stockpile struct {
	*StateCore
	TTL int
}

func NewStockpile() *Stockpile {
	return &Stockpile{
		StateCore: &StateCore{
			OrganismCore: nil,
			name:         "Stockpile",
		},
		TTL: 2,
	}
}

func (s *Stockpile) TakeTurn() {
	s.StateCore.TakeTurn()
	s.TTL--
	if s.TTL == 0 {
		s.OrganismCore.SetState(NewEruption())
		return
	}
}

func (s *Stockpile) Injured(power int) {
	s.StateCore.Injured(power)
	s.OrganismCore.ResetState()
}
