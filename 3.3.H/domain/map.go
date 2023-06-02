package domain

import (
	"errors"
	"fmt"
)

type Map struct {
	Width  int
	Height int
	Cells  [][]MapObject
}

func DefaultMap() *Map {
	m := NewMap(10, 10)

	m.InsertObject(NewCharacter(m), Position{X: 0, Y: 0})
	m.InsertObject(NewMonster(m), Position{X: 5, Y: 5})
	m.InsertObject(NewKingRock(m), Position{X: 3, Y: 3})
	return m
}

func NewMap(width int, height int) *Map {
	cells := make([][]MapObject, height)
	for i := range cells {
		cells[i] = make([]MapObject, width)
	}
	return &Map{Width: width, Height: height, Cells: cells}
}

func (m *Map) InsertObject(obj MapObject, p Position) error {
	if p.X < 0 || p.X >= m.Width || p.Y < 0 || p.Y >= m.Height {
		return errors.New("out of range")
	}
	m.Cells[p.Y][p.X] = obj
	obj.SetPosition(p)
	return nil
}

func (m *Map) GetCell(x int, y int) MapObject {
	return m.Cells[y][x]
}

func (m *Map) SetCell(x int, y int, obj MapObject) {
	m.Cells[y][x] = obj
	if obj != nil {
		obj.SetPosition(Position{X: x, Y: y})
	}
}

func (m *Map) Move(from Position, to Position) error {
	if to.X < 0 || to.X >= m.Width || to.Y < 0 || to.Y >= m.Height {
		return errors.New("out of range")
	}
	obj := m.GetCell(from.X, from.Y)
	m.SetCell(from.X, from.Y, nil)
	m.SetCell(to.X, to.Y, obj)
	return nil
}

func (m *Map) FindCharater() Position {
	for y, row := range m.Cells {
		for x, obj := range row {
			if IsCharacter(obj) {
				return Position{X: x, Y: y}
			}
		}
	}
	return Position{X: -1, Y: -1}
}

func (m *Map) ShowMap() {
	character := m.GetFirstCharacters()
	fmt.Printf("HP: %d, State : %s\n", character.HP, character.CurrentState.GetName())
	for _, row := range m.Cells {
		for _, obj := range row {
			if obj != nil {
				print(obj.GetSymbol())
			} else {
				print(" ")
			}
		}
		println()
	}
}

func (m *Map) GetObjects(p Position) MapObject {
	if !m.InMap(p) {
		return nil
	}
	return m.Cells[p.Y][p.X]
}

func (m *Map) GetAllObjects() []MapObject {
	var objects []MapObject
	for _, row := range m.Cells {
		for _, obj := range row {
			if obj != nil {
				objects = append(objects, obj)
			}
		}
	}
	return objects
}

func (m *Map) GetEmpty() []Position {
	var empty []Position
	for y, row := range m.Cells {
		for x, obj := range row {
			if obj == nil {
				empty = append(empty, Position{X: x, Y: y})
			}
		}
	}
	return empty
}

func (m *Map) GetFirstCharacters() *Character {
	for _, obj := range m.GetAllObjects() {
		if IsCharacter(obj) {
			return obj.(*Character)
		}
	}
	return nil
}

func (m *Map) TakeTurn() {

	for _, obj := range m.GetAllObjects() {
		if IsOrganism(obj) {
			if obj.(Organism).IsDead() {
				m.Remove(obj.GetPosition())
			}
		}
	}

	character := m.GetFirstCharacters()
	if character == nil {
		panic("game over")
	}

	character.TakeTurn()

	for _, obj := range m.GetAllObjects() {
		if IsMonster(obj) {
			obj.(*Monster).TakeTurn()
		}
	}
}

func (m *Map) InMap(p Position) bool {
	return p.X >= 0 && p.X < m.Width && p.Y >= 0 && p.Y < m.Height
}

func (m *Map) Remove(p Position) {
	m.Cells[p.Y][p.X] = nil
}
