package domain

type CollisionHandler interface {
	Collision(ISprite, ISprite) (ISprite, ISprite)
}

type CollisionHandlerTemplate struct {
	Next           CollisionHandler
	collisionLogic func(ISprite, ISprite) (ISprite, ISprite)
	comboType      SpriteType
}

func CollisionHandlerFactor() CollisionHandler {
	waterFireCollisionHandler := NewWaterFireCollisionHandler()
	heroWaterCollisionHandler := NewHeroWaterCollisionHandler()
	heroFireCollisionHandler := NewHeroFireCollisionHandler()

	waterFireCollisionHandler.Next = heroWaterCollisionHandler
	heroWaterCollisionHandler.Next = heroFireCollisionHandler

	return waterFireCollisionHandler
}

func (collisionHandler *CollisionHandlerTemplate) Collision(sprite1 ISprite, sprite2 ISprite) (ISprite, ISprite) {
	if sprite1.GetType() == sprite2.GetType() {
		return sprite1, sprite2
	}

	if sprite1.GetType()|sprite2.GetType() == collisionHandler.comboType {

		sprite1, sprite2 = collisionHandler.collisionLogic(sprite1, sprite2)
		if sprite1.IsMarkedForDeletion() {
			sprite1 = nil
		}
		if sprite2.IsMarkedForDeletion() {
			sprite2 = nil
		}

		return sprite1, sprite2
	} else {

		if collisionHandler.Next == nil {
			return sprite1, sprite2
		}

		return collisionHandler.Next.Collision(sprite1, sprite2)
	}
}

// Hero collides Fire or Fire collides Hero
func HeroFireCollision(sprite1 ISprite, sprite2 ISprite) (ISprite, ISprite) {
	var hero ISprite
	var fire ISprite
	if sprite1.GetType() == FIRE {
		hero, fire = sprite2, sprite1
	} else {
		hero, fire = sprite1, sprite2
	}

	fire.MarkedForDeletion()
	heroGainHP(hero, -10)

	return sprite2, sprite1
}

func NewHeroFireCollisionHandler() *CollisionHandlerTemplate {
	return &CollisionHandlerTemplate{
		collisionLogic: HeroFireCollision,
		comboType:      HERO | FIRE,
	}
}

// Water collides Fire or Fire collides Water
func WaterFireCollision(sprite1 ISprite, sprite2 ISprite) (ISprite, ISprite) {
	sprite1.MarkedForDeletion()
	sprite2.MarkedForDeletion()
	return sprite1, sprite2
}

func NewWaterFireCollisionHandler() *CollisionHandlerTemplate {
	return &CollisionHandlerTemplate{
		collisionLogic: WaterFireCollision,
		comboType:      WATER | FIRE,
	}
}

// Hero collides Water or Water collides Hero
func HeroWaterCollision(sprite1 ISprite, sprite2 ISprite) (ISprite, ISprite) {
	var hero ISprite
	var water ISprite
	if sprite1.GetType() == WATER {
		hero, water = sprite2, sprite1
	} else {
		hero, water = sprite1, sprite2
	}

	water.MarkedForDeletion()
	heroGainHP(hero, 10)

	return sprite2, sprite1
}

func NewHeroWaterCollisionHandler() *CollisionHandlerTemplate {
	return &CollisionHandlerTemplate{
		collisionLogic: HeroWaterCollision,
		comboType:      HERO | WATER,
	}
}

func heroGainHP(sprite ISprite, health int) {
	if hero, ok := sprite.(*Hero); ok {
		hero.GainHP(health)
	}
}
