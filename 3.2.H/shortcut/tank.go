package shortcut

import "fmt"

type Tank struct{}

func (c *Tank) MoveForward() {
	fmt.Println("The tank has moved forward.")
}

func (c *Tank) MoveBackward() {
	fmt.Println("The tank has moved backward.")
}

type MoveTankBackward struct {
	Tank *Tank
}

func (c MoveTankBackward) Do() {
	c.Tank.MoveBackward()
}

func (c MoveTankBackward) Name() string {
	return "MoveTankBackward"
}

func (c MoveTankBackward) Undo() {
	c.Tank.MoveForward()
}

type MoveTankForward struct {
	Tank *Tank
}

func (c MoveTankForward) Do() {
	c.Tank.MoveForward()
}

func (c MoveTankForward) Undo() {
	c.Tank.MoveBackward()
}

func (c MoveTankForward) Name() string {
	return "MoveTankForward"
}
