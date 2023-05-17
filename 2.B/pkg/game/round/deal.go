package round

import "big2/pkg/game/component"

func Deal(b *component.BigTwo) {
	for b.Deck.Len() > 0 {
		for _, player := range b.Players {
			player.DealtHandCards(b.Deck.Deal())
		}
	}

	for _, player := range b.Players {
		player.SortHandCards()
	}
}
