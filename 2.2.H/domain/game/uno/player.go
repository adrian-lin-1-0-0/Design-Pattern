package uno

import "2.2.H/domain/game/base"

type PlayerCore struct {
	topCard base.ICard
}

func (playerCore *PlayerCore) SetTopCard(card base.ICard) {
	playerCore.topCard = card
}
