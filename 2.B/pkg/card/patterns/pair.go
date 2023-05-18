package patterns

import "big2/pkg/card"

type PairPattern struct {
	cards []card.Card
}

func NewPairPattern() *PairPattern {
	return &PairPattern{}
}

func (p *PairPattern) GreaterThan(pattern CardPattern) bool {
	if len(pattern.GetCards()) != 2 {
		return false
	}
	return p.cards[1].GreaterThan(&pattern.GetCards()[1])
}

func (p *PairPattern) GetCards() []card.Card {
	return p.cards
}

func (p *PairPattern) Match(cards []card.Card) bool {
	if len(cards) != 2 {
		return false
	}
	return cards[0].Rank == cards[1].Rank
}

func (p *PairPattern) GetName() string {
	return "對子"
}

func (p *PairPattern) New(cards []card.Card) CardPattern {
	return &PairPattern{
		cards: cards,
	}
}
