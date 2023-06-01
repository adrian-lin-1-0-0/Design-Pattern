package domain

import "errors"

type Map struct {
	Width  int
	Height int
	Cells  [][]MapObject
}

func NewMap(width int, height int) *Map {
	cells := make([][]MapObject, height)
	for i := range cells {
		cells[i] = make([]MapObject, width)
	}
	return &Map{Width: width, Height: height, Cells: cells}
}

func (m *Map) GetCell(x int, y int) MapObject {
	return m.Cells[y][x]
}

func (m *Map) SetCell(x int, y int, obj MapObject) {
	m.Cells[y][x] = obj
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
			if obj != nil && obj.GetType() == _Character {
				return Position{X: x, Y: y}
			}
		}
	}
	return Position{X: -1, Y: -1}
}

func (m *Map) ShowMap() {
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
	if p.X < 0 || p.X >= m.Width || p.Y < 0 || p.Y >= m.Height {
		return nil
	}
	return m.Cells[p.Y][p.X]
}
