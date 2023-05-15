package patterns

import "big2"

type CardPattern interface {
	GreaterThan(p *big2.CardPattens) bool
	Identify(cards []big2.Card) bool
	New(cards []big2.Card) CardPattern
}

type CardPatterns struct {
	Next    *CardPatterns
	Pattern CardPattern
}

func (c *CardPatterns) ToPattern(cards []big2.Card) CardPattern {
	if c.Pattern.Identify(cards) {
		return c.Pattern.New(cards)
	}

	if c.Next != nil {
		return c.Next.ToPattern(cards)
	}

	return nil
}
