package uno

import "2.2.H/domain/game/base"

func NewDeck() *base.BaseDeck {
	deck := &base.BaseDeck{}
	genUno(deck)
	return deck
}

func genUno(d *base.BaseDeck) {
	for colorValue, colorName := range ColorNameList {
		for numberValue, numberName := range NumberNameList {
			d.AddCard(&UnoCard{
				Number: base.CardDenomination{
					Name:  numberName,
					Value: numberValue,
				},
				Color: base.CardDenomination{
					Name:  colorName,
					Value: colorValue,
				},
			})
		}
	}
}
