package domain

import (
	"errors"
	"math/rand"
	"time"
)

type Orderless struct {
	*StateCore
	TTL         int
	moveOptions int
}

func NewOrderless() *Orderless {
	return &Orderless{
		StateCore: &StateCore{
			OrganismCore: nil,
			name:         "Orderless",
		},
		TTL: 3,
	}
}

func (o *Orderless) TakeTurn() {
	rand.Seed(time.Now().UnixNano())
	o.moveOptions = rand.Intn(2)
	o.StateCore.TakeTurn()
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
