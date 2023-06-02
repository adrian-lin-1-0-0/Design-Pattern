package domain

import "testing"

func TestMonsterTouchPoison(t *testing.T) {
	_map := NewMap(10, 10)

	m := NewMonster(_map)
	_map.InsertObject(m, Position{X: 0, Y: 0})

	p := NewPoison(_map)
	m.Touch(p)
	m.TakeTurn()
	if m.IsDead() != true {
		t.Errorf("monster should be dead")
	}
}

func TestMonsterTouchAcceleratingPotion(t *testing.T) {
	_map := NewMap(10, 10)

	m := NewMonster(_map)
	_map.InsertObject(m, Position{X: 0, Y: 0})

	p := NewAcceleratingPotion(_map)
	m.Touch(p)
	m.TakeTurn()
	if m.ActionCount != 2 {
		t.Errorf("monster should be accelerated")
	}
	m.ResetState()
	if m.ActionCount != 1 {
		t.Errorf("monster should be reset")
	}
}
