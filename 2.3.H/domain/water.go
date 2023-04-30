package domain

type Water struct {
	Sprite
}

func NewWater() *Water {
	return &Water{
		Sprite: Sprite{
			spriteType: WATER,
		},
	}
}
