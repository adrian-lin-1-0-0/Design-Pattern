package domain

type OrganismCore struct {
	HP              int
	MaxHP           int
	CurrentState    State
	DefaultState    State
	AttackRange     int
	ActionCount     int
	AttackPower     int
	AttackDirection string
	Position        Position
	MapObjectCore   *MapObjectCore
}

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
	o.CurrentState = state
	o.CurrentState.Effect(o)
}

func (o *OrganismCore) GetSymbol() string {
	return o.MapObjectCore.Symbol
}

func (o *OrganismCore) TakeTurn() {
	o.CurrentState.TakeTurn()
}

type Organism interface {
	Attack(Organism)
	MoveTo(Position)
	IsDead() bool
	Action()
	Touch(MapObject)
	GetSymbol() string
	GetPosition() Position
	SetPosition(Position)
	Injured(int)
}
