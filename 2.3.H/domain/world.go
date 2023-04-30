package domain

type World struct {
	Space            [30]ISprite
	collisionHandler CollisionHandler
}

func (w *World) MoveSprite(fromPosition int, toPosition int) {
	if w.Space[fromPosition] == nil {
		return
	}
	if w.Space[toPosition] == nil {
		w.Space[toPosition], w.Space[fromPosition] = w.Space[fromPosition], nil
		return
	}

	w.Space[fromPosition], w.Space[toPosition] = w.collisionHandler.Collision(w.Space[fromPosition], w.Space[toPosition])
}

func NewWorld() *World {
	return &World{
		Space:            [30]ISprite{},
		collisionHandler: CollisionHandlerFactor(),
	}
}
