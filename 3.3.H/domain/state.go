package domain

type State interface {
	Effect(*OrganismCore)
	Injured(power int)
	Heal(health int)
	Attack() error
	MoveTo(p Position) error
	Touch(obj MapObject)
	TakeTurn()
	IsDead() bool
	Move(direction Dircetion) error
	Action()
	GetName() string
}

type StateCore struct {
	OrganismCore *OrganismCore
	name         string
}

func (s *StateCore) TakeTurn() {
	s.OrganismCore.DefaultState.TakeTurn()
}

func (s *StateCore) Touch(obj MapObject) {
	s.OrganismCore.DefaultState.Touch(obj)
}

func (s *StateCore) Effect(o *OrganismCore) {
	s.OrganismCore = o
}

func (s *StateCore) IsDead() bool {
	return s.OrganismCore.DefaultState.IsDead()
}

func (s *StateCore) Heal(health int) {
	s.OrganismCore.DefaultState.Heal(health)
}

func (s *StateCore) Injured(power int) {
	s.OrganismCore.DefaultState.Injured(power)
}

func (s *StateCore) Attack() error {
	return s.OrganismCore.DefaultState.Attack()
}

func (s *StateCore) Move(direction Dircetion) error {
	return s.OrganismCore.DefaultState.Move(direction)
}

func (s *StateCore) MoveTo(p Position) error {
	return s.OrganismCore.DefaultState.MoveTo(p)
}

func (s *StateCore) Action() {
	s.OrganismCore.DefaultState.Action()
}

func (s *StateCore) GetName() string {
	return s.name
}
