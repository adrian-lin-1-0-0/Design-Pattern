package uno

import (
	"2.2.H/domain/game/base"
)

type UnoCard struct {
	Number base.CardDenomination
	Color  base.CardDenomination
}

func (card *UnoCard) Compare(other any) base.Ord {
	otherCard, _ := other.(*UnoCard)
	if card.Number.Compare(&otherCard.Number) == base.EQ {
		return base.EQ
	}
	return card.Color.Compare(&otherCard.Color)
}

var ColorNameList = [...]string{"BLUE", "RED", "YELLOW", "GREEN"}
var NumberNameList = [...]string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
