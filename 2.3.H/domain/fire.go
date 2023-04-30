package domain

type Fire struct {
	Sprite
}

func NewFire() *Fire {
	return &Fire{
		Sprite: Sprite{
			spriteType: FIRE,
		},
	}
}
