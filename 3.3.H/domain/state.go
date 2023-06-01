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
}

type Normal struct {
	*OrganismCore
}

func NewNormal() *Normal {
	return &Normal{
		OrganismCore: nil,
	}
}

func (n *Normal) Heal(health int) {
	n.HP += health
}

func (n *Normal) Effect(o *OrganismCore) {
	n.OrganismCore = o
	n.ActionCount = 1
}

func (n *Normal) Injured(power int) {
	n.HP -= power
	n.SetState(NewInvincible())
}

func (n *Normal) Action() {
	panic("not implemented")
}

func (n *Normal) Attack() error {
	panic("not implemented")
}

func (n *Normal) Move(direction Dircetion) error {
	p := n.Position
	switch direction {
	case Up:
		p.Y--
	case Down:
		p.Y++
	case Left:
		p.X--
	case Right:
		p.X++
	}
	return n.MoveTo(p)
}

func (n *Normal) MoveTo(p Position) error {
	if obj := n.MapObjectCore.Map.GetObjects(p); obj != nil {
		n.Touch(obj)
	}
	if err := n.MapObjectCore.Map.Move(n.Position, p); err != nil {
		return err
	}
	n.Position = p
	return nil
}

func (n *Normal) Touch(obj MapObject) {
	if state := obj.Touched(); state != nil {
		n.SetState(state)
	}
}

func (n *Normal) TakeTurn() {
	actionCountDown := n.ActionCount

	for ; actionCountDown > 0; actionCountDown-- {
		n.Action()
	}
}

func (n *Normal) IsDead() bool {
	return n.OrganismCore.HP <= 0
}
