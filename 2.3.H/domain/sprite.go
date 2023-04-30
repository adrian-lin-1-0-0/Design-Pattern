package domain

type ISprite interface {
	MarkedForDeletion()
	IsMarkedForDeletion() bool
	GetType() SpriteType
}
type SpriteType int

const (
	NONE SpriteType = 1 << iota
	HERO
	WATER
	FIRE
)

type Sprite struct {
	markedForDeletion bool
	spriteType        SpriteType
}

func (s *Sprite) MarkedForDeletion() {
	s.markedForDeletion = true
}

func (s *Sprite) IsMarkedForDeletion() bool {
	return s.markedForDeletion
}

func (s *Sprite) GetType() SpriteType {
	return s.spriteType
}
