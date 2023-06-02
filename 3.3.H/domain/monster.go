package domain

import (
	"errors"
	"fmt"
)

type Monster struct {
	*OrganismCore
}

func NewMonster(m *Map) *Monster {
	o := &OrganismCore{
		MapObjectCore: &MapObjectCore{
			Symbol: "M",
			Type:   _Organism,
			Map:    m,
		},
		OrganismType: _Monster,
		AttackPower:  50,
		HP:           1,
		MaxHP:        1,
	}

	o.SetState(NewMonsterNormal())

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
	if err := m.CurrentState.Attack(); err == nil {
		fmt.Println("Monster attack")
		return
	}

	for _, direction := range []Dircetion{Up, Down, Left, Right} {
		if err := m.Normal.Move(direction); err == nil {
			fmt.Println("Monster move to", direction)
			return
		}
	}

	panic("cannot move")
}

func (m *MonsterNormal) Attack() error {
	for _, p := range GetCross(m.Position) {
		mapObject := m.MapObjectCore.Map.GetObjects(p)
		if IsCharacter(mapObject) {
			mapObject.(Organism).Injured(m.AttackPower)
			return nil
		}
	}
	return errors.New("cannot attack")
}
