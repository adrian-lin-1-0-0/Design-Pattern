package showdown

import "2.2.H/domain/game/base"

func NewDeck() *base.BaseDeck {
	deck := &base.BaseDeck{}
	genPoker(deck)
	return deck
}

func genPoker(d *base.BaseDeck) {
	for suitValue, suitName := range SuitNameList {
		for rankValue, rankName := range RankNameList {
			d.AddCard(&PokerCard{
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
