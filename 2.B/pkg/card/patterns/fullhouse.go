package patterns

import (
	"big2/pkg/card"
)

type fullHousePattern struct {
	cards      []card.Card
	compareIdx int
}

func NewFullHousePattern() *fullHousePattern {
	return &fullHousePattern{}
}

func (p *fullHousePattern) GreaterThan(pattern CardPattern) bool {
	cards := pattern.GetCards()

	return p.cards[p.compareIdx].GreaterThan(&cards[findFullHouseCompareIdx(cards)])
}

func findFullHouseCompareIdx(cards []card.Card) int {
	countFirstRank := 0
	countSecondRank := 0
	for _, c := range cards {
		if c.Rank == cards[0].Rank {
			countFirstRank++
		} else {
			countSecondRank++
		}
	}

	compareIdx := -1
	if countFirstRank == 3 {
		compareIdx = 2
	} else {
		compareIdx = 4
	}
	return compareIdx
}

func (p *fullHousePattern) GetCards() []card.Card {
	return p.cards
}

func (p *fullHousePattern) Match(cards []card.Card) bool {
	if len(cards) != 5 {
		return false
	}

	rankCount := make(map[card.Rank]int)
	for _, card := range cards {
		rankCount[card.Rank]++
	}

	var hasThreeOfAKind, hasPair bool
	for _, count := range rankCount {
		if count == 3 {
			hasThreeOfAKind = true
		} else if count == 2 {
			hasPair = true
		}
	}

	return hasThreeOfAKind && hasPair
}

func (p *fullHousePattern) GetName() string {
	return "葫蘆"
}

func (p *fullHousePattern) New(cards []card.Card) CardPattern {
	return &fullHousePattern{
		cards:      cards,
		compareIdx: findFullHouseCompareIdx(cards),
	}
}
