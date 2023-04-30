package domain

type Hero struct {
	HP int
	Sprite
}

func (hero *Hero) IsDead() bool {
	return hero.HP <= 0
}

func (hero *Hero) GainHP(heal int) {
	hero.HP += heal
	if hero.IsDead() {
		hero.MarkedForDeletion()
	}
}

func NewHero() *Hero {
	return &Hero{
		HP: 30,
		Sprite: Sprite{
			spriteType: HERO,
		},
	}
}
