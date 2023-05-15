package patterns

import "big2"

type SinglePattern struct {
	cards []big2.Card
}

func (s *SinglePattern) GreaterThan(p *big2.CardPattens) bool {
	return s.cards[0].Rank > p.Cards[0].Rank
}

func (s *SinglePattern) Identify(cards []big2.Card) bool {
	return len(cards) == 1
}

func (s *SinglePattern) New(cards []big2.Card) CardPattern {
	return &SinglePattern{
		cards: cards,
	}
}
