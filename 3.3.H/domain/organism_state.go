package domain

type Normal struct {
	*OrganismCore
	name string
}

func NewNormal() *Normal {
	return &Normal{
		name:         "Normal",
		OrganismCore: nil,
	}
}

func (n *Normal) Heal(health int) {
	n.HP += health
}

func (n *Normal) Effect(o *OrganismCore) {
	n.OrganismCore = o
	n.OrganismCore.ActionCount = 1
}

func (n *Normal) Injured(power int) {
	n.HP -= power
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
	if err := n.MoveTo(p); err != nil {
		return err
	}
	return nil
}

func (n *Normal) MoveTo(p Position) error {
	if obj := n.MapObjectCore.Map.GetObjects(p); obj != nil {
		n.Touch(obj)
		return nil
	}
	if err := n.MapObjectCore.Map.Move(n.Position, p); err != nil {
		return err
	}
	return nil
}

func (n *Normal) Touch(obj MapObject) {
	if state := obj.Touched(); state != nil {
		n.SetState(state)
	}
}

func (n *Normal) TakeTurn() {
	actionCountDown := n.OrganismCore.ActionCount

	for ; actionCountDown > 0; actionCountDown-- {
		n.OrganismCore.Action()
	}
}

func (n *Normal) IsDead() bool {
	return n.OrganismCore.HP <= 0
}

func (n *Normal) GetName() string {
	return n.name
}
