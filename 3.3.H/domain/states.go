package domain

import (
	"errors"
	"math/rand"
	"time"
)

type Accelerated struct {
	State
	TTL          int
	OrganismCore *OrganismCore
}

func NewAccelerated() *Accelerated {
	return &Accelerated{
		TTL: 2}
}

func (a *Accelerated) Effect(o *OrganismCore) {
	a.OrganismCore = o
	a.OrganismCore.ActionCount = 2
	a.State = o.DefaultState
	a.State.Effect(o)
}

func (a *Accelerated) TakeTurn() {
	a.State.TakeTurn()
	a.TTL--
	if a.TTL == 0 {
		a.OrganismCore.ResetState()
	}
	a.Attack()
}

func (a *Accelerated) Injured(power int) {
	a.State.Injured(power)
	a.OrganismCore.ResetState()
}

type Poisoned struct {
	State
	TTL          int
	OrganismCore *OrganismCore
}

func NewPoisoned() *Poisoned {
	return &Poisoned{
		TTL: 3}
}

func (p *Poisoned) TakeTurn() {
	p.OrganismCore.Injured(10)
	p.State.TakeTurn()
	p.TTL--
	if p.TTL == 0 {
		p.OrganismCore.ResetState()
		return
	}
}

type Invincible struct {
	State
	TTL          int
	OrganismCore *OrganismCore
}

func NewInvincible() *Invincible {
	return &Invincible{
		TTL: 2,
	}
}

func (i *Invincible) Injured(power int) {
	// do nothing
}

func (i *Invincible) TakeTurn() {
	i.State.TakeTurn()
	i.TTL--
	if i.TTL == 0 {
		i.OrganismCore.ResetState()
		return
	}
}

type Healing struct {
	State
	TTL          int
	OrganismCore *OrganismCore
}

func NewHealing() *Healing {
	return &Healing{
		TTL: 5,
	}
}

func (h *Healing) TakeTurn() {
	h.Heal(30)
	h.State.TakeTurn()
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

type Orderless struct {
	State
	TTL          int
	OrganismCore *OrganismCore
	moveOptions  int
}

func NewOrderless() *Orderless {
	return &Orderless{
		TTL: 3,
	}
}

func (o *Orderless) TakeTurn() {
	rand.Seed(time.Now().UnixNano())
	o.moveOptions = rand.Intn(2)
	o.State.TakeTurn()
	o.TTL--
	if o.TTL == 0 {
		o.OrganismCore.ResetState()
		return
	}
}

func (o *Orderless) Attack() error {
	return errors.New("cannot attack")
}

func (o *Orderless) Move(direction Dircetion) error {
	p := o.OrganismCore.Position
	if o.moveOptions == 0 {
		switch direction {
		case Up:
			p.Y--
		case Down:
			p.Y++
		case Left:
			return errors.New("cannot move left")
		case Right:
			return errors.New("cannot move right")
		}
	} else {
		switch direction {
		case Up:
			return errors.New("cannot move up")
		case Down:
			return errors.New("cannot move down")
		case Left:
			p.X--
		case Right:
			p.X++
		}
	}

	return o.MoveTo(p)
}

type Stockpile struct {
	State
	TTL          int
	OrganismCore *OrganismCore
}

func NewStockpile() *Stockpile {
	return &Stockpile{
		TTL: 2,
	}
}

func (s *Stockpile) TakeTurn() {
	s.State.TakeTurn()
	s.TTL--
	if s.TTL == 0 {
		s.OrganismCore.ResetState()
		return
	}
}

func (s *Stockpile) Injured(power int) {
	s.State.Injured(power)
	s.OrganismCore.SetState(NewEruption())
}

type Eruption struct {
	State
	TTL          int
	OrganismCore *OrganismCore
}

func NewEruption() *Eruption {
	return &Eruption{
		TTL: 3,
	}
}

func (e *Eruption) TakeTurn() {
	e.State.TakeTurn()
	e.TTL--
	if e.TTL == 0 {
		e.OrganismCore.SetState(NewTeleport())
		return
	}
}

func (e *Eruption) Effect(o *OrganismCore) {
	e.State.Effect(o)
	e.OrganismCore.AttackPower = 50
}

type Teleport struct {
	*Normal
	TTL int
}

func NewTeleport() *Teleport {
	return &Teleport{
		Normal: NewNormal(),
		TTL:    1,
	}
}
