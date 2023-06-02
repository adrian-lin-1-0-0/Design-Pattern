package domain

import "math/rand"

type Treasure struct {
	*MapObjectCore
	State State
	name  string
}

func (t Treasure) Touched() State {
	t.MapObjectCore.Map.Remove(t.GetPosition())
	return t.State
}

func (t Treasure) GetName() string {
	return t.name
}

func NewTreasureRandom(m *Map) *Treasure {
	r := rand.Intn(100)
	switch {
	case r <= 10:
		return NewDokodemoDoor(m)
	case r <= 20:
		return NewSuperStar(m)
	case r <= 45:
		return NewPoison(m)
	case r <= 65:
		return NewAcceleratingPotion(m)
	case r <= 80:
		return NewHealingPotion(m)
	case r <= 90:
		return NewDevilFruit(m)
	case r <= 100:
		return NewKingRock(m)
	}
	return nil
}

func NewTreasureMapObjectCore(m *Map) *MapObjectCore {
	return &MapObjectCore{
		Symbol: "x",
		Type:   _Treasure,
		Map:    m,
	}
}

func NewDokodemoDoor(m *Map) *Treasure {
	return &Treasure{
		MapObjectCore: NewTreasureMapObjectCore(m),
		State:         NewTeleport(),
		name:          "Dokodemo Door",
	}
}

func NewSuperStar(m *Map) *Treasure {
	return &Treasure{
		MapObjectCore: NewTreasureMapObjectCore(m),
		State:         NewInvincible(),
		name:          "Super Star",
	}
}

func NewPoison(m *Map) *Treasure {
	return &Treasure{
		MapObjectCore: NewTreasureMapObjectCore(m),
		State:         NewPoisoned(),
		name:          "Poison",
	}
}

func NewAcceleratingPotion(m *Map) *Treasure {
	return &Treasure{
		MapObjectCore: NewTreasureMapObjectCore(m),
		State:         NewAccelerated(),
		name:          "Accelerating Potion",
	}
}

func NewHealingPotion(m *Map) *Treasure {
	return &Treasure{
		MapObjectCore: NewTreasureMapObjectCore(m),
		State:         NewHealing(),
		name:          "Healing Potion",
	}
}

func NewDevilFruit(m *Map) *Treasure {
	return &Treasure{
		MapObjectCore: NewTreasureMapObjectCore(m),
		State:         NewOrderless(),
		name:          "Devil Fruit",
	}
}

func NewKingRock(m *Map) *Treasure {
	return &Treasure{
		MapObjectCore: NewTreasureMapObjectCore(m),
		State:         NewStockpile(),
		name:          "King's Rock",
	}
}
