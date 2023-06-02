package domain

import (
	"sync"
)

type Organism interface {
	MapObject
	Attack() error
	MoveTo(Position) error
	IsDead() bool
	Action()
	Move(Dircetion) error
	Touch(MapObject)
	Injured(int)
	GetOrganismType() OrganismType
}

type OrganismCore struct {
	*MapObjectCore
	HP              int
	MaxHP           int
	CurrentState    State
	DefaultState    State
	ActionCount     int
	AttackPower     int
	AttackDirection string
	OrganismType    OrganismType
	once            sync.Once
	Dircetion       Dircetion
}

type OrganismType string

const (
	_Monster   OrganismType = "Monster"
	_Character OrganismType = "Character"
)

func (o *OrganismCore) ResetState() {
	o.SetState(o.DefaultState)
}

func NewOrganismCore() *OrganismCore {
	o := &OrganismCore{
		MapObjectCore: &MapObjectCore{},
	}
	o.SetState(NewNormal())
	return o
}

func (o *OrganismCore) Attack() error {
	return o.CurrentState.Attack()
}

func (o *OrganismCore) Move(direction Dircetion) error {
	return o.CurrentState.Move(direction)
}

func (o *OrganismCore) MoveTo(p Position) error {
	return o.CurrentState.MoveTo(p)
}

func (o *OrganismCore) Heal(health int) {
	o.CurrentState.Heal(health)
}

func (o *OrganismCore) Injured(power int) {
	o.CurrentState.Injured(power)
}

func (o *OrganismCore) IsDead() bool {
	return o.CurrentState.IsDead()
}

func (o *OrganismCore) Action() {
	o.CurrentState.Action()
}

func (o *OrganismCore) Touch(obj MapObject) {
	o.CurrentState.Touch(obj)
}

func (o *OrganismCore) SetState(state State) {
	o.once.Do(func() {
		o.DefaultState = state
		o.DefaultState.Effect(o)
	})
	o.CurrentState = state
	o.CurrentState.Effect(o)
}

func (o *OrganismCore) GetSymbol() string {
	return o.MapObjectCore.Symbol
}

func (o *OrganismCore) TakeTurn() {
	o.CurrentState.TakeTurn()
}

func (o *OrganismCore) GetType() MapObjectType {
	return o.MapObjectCore.Type
}

func (o *OrganismCore) Touched() State {
	return nil
}

func (o *OrganismCore) GetOrganismType() OrganismType {
	return o.OrganismType
}
