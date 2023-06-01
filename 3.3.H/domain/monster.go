package domain

type Monster struct {
	*OrganismCore
}

func NewMonster() *Monster {
	o := &OrganismCore{
		MapObjectCore: &MapObjectCore{
			Symbol: "M",
		},
		AttackPower:  50,
		HP:           1,
		MaxHP:        1,
		DefaultState: NewMonsterNormal(),
		CurrentState: NewMonsterNormal(),
	}

	return &Monster{
		OrganismCore: o,
	}
}

func NewMonsterNormal() *MonsterNormal {
	return &MonsterNormal{
		Normal: NewNormal(),
	}
}

type MonsterNormal struct {
	*Normal
}

func (m *MonsterNormal) Effect(o *OrganismCore) {
	m.Normal.Effect(o)
	m.OrganismCore.AttackPower = 50
}

func (m *MonsterNormal) Action() {
	for _, p := range getCross(m.Position) {
		mapObject := m.MapObjectCore.Map.GetObjects(p)
		if mapObject.GetType() == _Character {
			err := m.Attack()
			if err != nil {
				break
			}
			return
		}
	}

	for _, direction := range []Dircetion{Up, Down, Left, Right} {
		if err := m.Move(direction); err == nil {
			return
		}
	}

	panic("cannot move")
}

func (m *MonsterNormal) Attack() error {
	return nil
}
