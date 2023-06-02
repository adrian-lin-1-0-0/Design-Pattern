package domain

type Invincible struct {
	*StateCore
	TTL int
}

func NewInvincible() *Invincible {
	return &Invincible{
		StateCore: &StateCore{
			OrganismCore: nil,
			name:         "Invincible",
		},
		TTL: 2,
	}
}

func (i *Invincible) Injured(power int) {
	// do nothing
}

func (i *Invincible) TakeTurn() {
	i.StateCore.TakeTurn()
	i.TTL--
	if i.TTL == 0 {
		i.OrganismCore.ResetState()
		return
	}
}
