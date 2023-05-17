package round

import "big2/pkg/game"

func NamePlayer(b *game.BigTwo) {
	for _, p := range b.Players {
		p.NamePlayer()
	}
}
