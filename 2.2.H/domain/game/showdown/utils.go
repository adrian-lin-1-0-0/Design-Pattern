package showdown

import (
	"2.2.H/domain/game/base"
)

func FindMaxCardIdx(cards []base.ICard) (idx int) {
	if len(cards) == 0 {
		return -1
	}

	maxIdx := 0
	maxCard := cards[0]

	for i := 1; i < len(cards); i++ {
		if base.LT == maxCard.Compare(cards[i]) {
			maxIdx = idx
			maxCard, _ = cards[i].(*base.BaseCard)
		}
	}

	return maxIdx
}
