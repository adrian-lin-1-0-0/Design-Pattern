package game

import (
	"big2/pkg/card/patterns"
	"big2/pkg/player"
)

type BigTwo struct {
	rounds []Round
}

type BigTwoOptions struct {
	PlayerCount int
}

func NewBigTwo(opts *BigTwoOptions) *BigTwo {
	players := make([]*player.Player, opts.PlayerCount)
	for i := 0; i < opts.PlayerCount; i++ {
		players[i] = player.NewDefaultPlayer()
	}

	firstRound := NewFirstRound(&FirstRoundOptions{
		Players: players,
		Deck:    NewDeck(nil),
	})

	playround := NewPlayRound(&PlayRoundOptions{
		Players:           players,
		Table:             NewTable(),
		CardPatternsChain: patterns.CardPatternsFactory(),
	})

	return &BigTwo{
		rounds: []Round{firstRound, playround},
	}
}

func (b *BigTwo) AddRound(r Round) {
	b.rounds = append(b.rounds, r)
}

func (b *BigTwo) Run() {
	for _, r := range b.rounds {
		r.Run()
	}
}
