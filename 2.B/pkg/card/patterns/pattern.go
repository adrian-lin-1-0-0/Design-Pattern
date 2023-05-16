package patterns

import (
	"big2/pkg/card"
)

func CardPatternsFactory() *CardPatternsChain {
	return NewCardPatterns(NewSinglePattern()).
		Add(NewStraightPattern())

}

func NewCardPatterns(pattern CardPattern) *CardPatternsChain {
	return &CardPatternsChain{Pattern: pattern}
}

type CardPattern interface {
	GreaterThan(p CardPattern) bool
	Match(cards []card.Card) bool
	New(cards []card.Card) CardPattern
	GetCards() []card.Card
	GetName() string
}

type CardPatternsChain struct {
	Next    *CardPatternsChain
	Pattern CardPattern
}

func (c *CardPatternsChain) Add(pattern CardPattern) *CardPatternsChain {
	if c.Next == nil {
		c.Next = &CardPatternsChain{
			Pattern: pattern,
			Next:    nil,
		}
	} else {
		c.Next.Add(pattern)
	}
	return c
}

func (c *CardPatternsChain) ToPattern(cards []card.Card) (CardPattern, error) {
	if c.Pattern.Match(cards) {
		return c.Pattern.New(cards), nil
	}

	if c.Next != nil {
		return c.Next.ToPattern(cards)
	}

	return nil, ErrIllegalCardPattern
}
