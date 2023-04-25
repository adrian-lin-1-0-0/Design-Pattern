package showdown

import "2.2.H/domain/game/base"

type PokerCard struct {
	Rank base.CardDenomination
	Suit base.CardDenomination
}

func (card *PokerCard) Compare(other any) base.Ord {
	otherCard, _ := other.(*PokerCard)
	if card.Rank.Compare(&otherCard.Rank) == base.EQ {
		return card.Suit.Compare(&otherCard.Suit)
	}
	return card.Rank.Compare(&otherCard.Rank)
}

var SuitNameList = [...]string{"♣", "♦", "♥", "♠"}
var RankNameList = [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}
