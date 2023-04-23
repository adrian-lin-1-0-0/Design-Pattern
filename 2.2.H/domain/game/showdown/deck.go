package showdown

import "2.2.H/domain/game/base"

var SuitNameList = [...]string{"♣", "♦", "♥", "♠"}
var RankNameList = [...]string{"2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A"}

func NewDeck() *base.BaseDeck {
	deck := &base.BaseDeck{}
	genPoker(deck)
	return deck
}

func genPoker(d *base.BaseDeck) {
	for suitValue, suitName := range SuitNameList {
		for rankValue, rankName := range RankNameList {
			d.AddCard(&base.BaseCard{
				Rank: base.CardDenomination{
					Name:  rankName,
					Value: rankValue,
				},
				Suit: base.CardDenomination{
					Name:  suitName,
					Value: suitValue,
				},
			})
		}
	}
}
