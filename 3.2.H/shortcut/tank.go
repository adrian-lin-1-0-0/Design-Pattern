package shortcut

import "fmt"

type Tank struct{}

func (c *Tank) MoveForward() {
	fmt.Println("The tank has moved forward.")
}

func (c *Tank) MoveBackward() {
	fmt.Println("The tank has moved backward.")
}

type TankMoveBackward struct {
	Tank *Tank
}

func (c *TankMoveBackward) Do() {
	c.Tank.MoveBackward()
}

func (c *TankMoveBackward) Undo() {
	c.Tank.MoveForward()
}

type TankMoveForward struct {
	Tank *Tank
}

func (c *TankMoveForward) Do() {
	c.Tank.MoveForward()
}

func (c *TankMoveForward) Undo() {
	c.Tank.MoveBackward()
}
