package domain

type Healing struct {
	*StateCore
	TTL int
}

func NewHealing() *Healing {
	return &Healing{
		StateCore: &StateCore{
			OrganismCore: nil,
			name:         "Healing",
		},
		TTL: 5,
	}
}

func (h *Healing) TakeTurn() {
	h.Heal(30)
	h.StateCore.TakeTurn()
	if h.OrganismCore.HP >= h.OrganismCore.MaxHP {
		h.OrganismCore.ResetState()
		return
	}
	h.TTL--
	if h.TTL == 0 {
		h.OrganismCore.ResetState()
		return
	}
}
