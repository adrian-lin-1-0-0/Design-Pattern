package round

import "big2/pkg/game/component"

func NamePlayer(b *component.BigTwo) {
	for _, p := range b.Players {
		p.NamePlayer()
	}
}
