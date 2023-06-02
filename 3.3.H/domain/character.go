package domain

import (
	"fmt"
	"os"
)

type Character struct {
	*OrganismCore
	Dircetion Dircetion
}

func NewCharacter(m *Map) *Character {
	o := &OrganismCore{
		MapObjectCore: &MapObjectCore{
			Symbol: string(Up),
			Type:   _Organism,
			Map:    m,
		},
		OrganismType: _Character,
		Dircetion:    Up,
		AttackPower:  1,
		HP:           300,
		MaxHP:        300,
	}
	o.SetState(NewCharacterNormal())

	return &Character{
		OrganismCore: o,
	}
}

type CharacterNormal struct {
	*Normal
}

func NewCharacterNormal() *CharacterNormal {
	return &CharacterNormal{
		Normal: NewNormal(),
	}
}

func (c *CharacterNormal) Injured(damage int) {
	c.Normal.Injured(damage)
	c.SetState(NewInvincible())
}

func (c *CharacterNormal) Effect(o *OrganismCore) {
	c.Normal.Effect(o)
	c.OrganismCore.AttackPower = 1
}

func (c *CharacterNormal) Action() {

	for {
		var options string
		fmt.Print("w, a, s, d: 移動, q: 攻擊, e: 結束遊戲")
		fmt.Scanln(&options)
		var err error

		switch options {
		case "w":
			err = c.CurrentState.Move(Up)
		case "a":
			err = c.CurrentState.Move(Left)
		case "s":
			err = c.CurrentState.Move(Down)
		case "d":
			err = c.CurrentState.Move(Right)
		case "q":
			err = c.CurrentState.Attack()
		case "e":
			os.Exit(0)
		}
		if err == nil {
			break
		}
		fmt.Println(err.Error())

	}

}

func (c *CharacterNormal) Move(dircetion Dircetion) error {
	if err := c.Normal.Move(dircetion); err != nil {
		return err
	}
	c.Dircetion = dircetion
	c.MapObjectCore.Symbol = string(dircetion)
	return nil
}

func (c *CharacterNormal) Attack() error {
	c.AttackDircetion(c.Dircetion)
	return nil
}

func (c *CharacterNormal) AttackDircetion(dircetion Dircetion) {
	for _, organism := range c.FindOrganismByDircetion(c.Dircetion) {
		organism.Injured(c.AttackPower)
	}
}

func (c *CharacterNormal) FindOrganismByDircetion(dircetion Dircetion) []Organism {
	var organisms []Organism
	position := c.Position
	for {
		position = c.GetNextPosition(position, dircetion)
		if !c.MapObjectCore.Map.InMap(position) {
			break
		}
		if mapObj := c.MapObjectCore.Map.GetObjects(position); mapObj != nil {
			if mapObj.GetType() == _Obstacle {
				break
			}
			if mapObj.GetType() == _Organism {
				organisms = append(organisms, mapObj.(Organism))
			}
		}
	}
	return organisms
}

func (c *CharacterNormal) GetNextPosition(position Position, dircetion Dircetion) Position {
	switch dircetion {
	case Up:
		position.Y--
	case Down:
		position.Y++
	case Left:
		position.X--
	case Right:
		position.X++
	}
	return position
}
