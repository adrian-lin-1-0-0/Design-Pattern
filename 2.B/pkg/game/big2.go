package game

import (
	"big2/pkg/game/component"
	"big2/pkg/game/round"
)

func NewDefaultBigTwo() *component.BigTwo {
	return component.NewBigTwo(&component.BigTwoOptions{PlayerCount: 4}).
		AddRound(round.Deal).
		AddRound(round.NamePlayer).
		AddRound(round.DefaultPlay)
}
