package domain

import (
	"testing"
)

func TestFireAndWaterInWorld(t *testing.T) {
	world := &World{
		Space:            [30]ISprite{},
		collisionHandler: CollisionHandlerFactor(),
	}
	world.Space[0] = NewFire()
	world.Space[1] = NewWater()
	world.MoveSprite(0, 1)

	if world.Space[0] != nil {
		t.Error("Fire should not be deleted")
	}

	if world.Space[1] != nil {
		t.Error("Water should not be deleted")
	}
}
func TestHeroAndFireInWorld(t *testing.T) {
	world := &World{
		Space:            [30]ISprite{},
		collisionHandler: CollisionHandlerFactor(),
	}
	world.Space[0] = NewHero()
	for i := 1; i <= 3; i++ {
		world.Space[i] = NewFire()
	}
	world.MoveSprite(0, 1)

	if world.Space[0] != nil {
		t.Error("Hero should be moved")
	}

	world.MoveSprite(1, 2)
	if world.Space[1] != nil {
		t.Error("Hero should be moved")
	}
	if world.Space[2] == nil {
		t.Error("Hero should be moved")
	}

	world.MoveSprite(2, 3)
	if world.Space[3] != nil {
		t.Error("Hero should not be deleted")
	}
}

func TestHeroAndWaterAndFireInWorld(t *testing.T) {
	world := &World{
		Space:            [30]ISprite{},
		collisionHandler: CollisionHandlerFactor(),
	}
	world.Space[0] = NewHero()
	world.Space[1] = NewWater()

	for i := 2; i <= 5; i++ {
		world.Space[i] = NewFire()
	}

	world.MoveSprite(0, 1)
	world.MoveSprite(1, 2)
	world.MoveSprite(2, 3)
	world.MoveSprite(3, 4)
	if world.Space[4] == nil {
		t.Error("Hero should not be deleted")
	}
	world.MoveSprite(4, 5)

}
