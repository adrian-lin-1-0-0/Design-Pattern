package round

import "big2/pkg/game"

func Deal(b *game.BigTwo) {
	for b.Deck.Len() > 0 {
		for _, player := range b.Players {
			player.DealtHandCards(b.Deck.Deal())
		}
	}
}
