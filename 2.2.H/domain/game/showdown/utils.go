package showdown

import (
	"2.2.H/domain/game/base"
)

var minCard = base.BaseCard{
	Rank: base.CardDenomination{Value: -1},
	Suit: base.CardDenomination{Value: -1},
}

func FinxMaxCardIdx(cards []base.ICard) (idx int) {
	if len(cards) == 0 {
		return -1
	}

	maxIdx := 0
	maxCard := minCard
	for idx, card := range cards {
		if base.LT == maxCard.Compare(card) {
			maxIdx = idx
			tmp, _ := card.(*base.BaseCard)
			maxCard = *tmp
		}
	}
	return maxIdx
}
