package patterns

import (
	"big2/pkg/card"
)

type SinglePattern struct {
	cards []card.Card
}

func NewSinglePattern() *SinglePattern {
	return &SinglePattern{}
}

func (s *SinglePattern) GreaterThan(p CardPattern) bool {

	return s.cards[0].GreaterThan(&p.GetCards()[0])
}

func (s *SinglePattern) GetCards() []card.Card {
	return s.cards
}

func (s *SinglePattern) Match(cards []card.Card) bool {
	return len(cards) == 1
}

func (s *SinglePattern) GetName() string {
	return "單張"
}

func (s *SinglePattern) New(cards []card.Card) CardPattern {
	return &SinglePattern{
		cards: cards,
	}
}
