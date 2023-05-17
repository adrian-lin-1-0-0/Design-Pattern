package component

import (
	"big2/pkg/player"
)

type BigTwo struct {
	Rounds  []Round
	Players []*player.Player
	Table   *Table
	Deck    *Deck
}

type BigTwoOptions struct {
	PlayerCount int
}

func NewBigTwo(opts *BigTwoOptions) *BigTwo {

	players := make([]*player.Player, opts.PlayerCount)
	for i := 0; i < opts.PlayerCount; i++ {
		players[i] = player.NewDefaultPlayer()
	}

	return (&BigTwo{
		Players: players,
		Deck:    NewDeck(nil),
		Table:   NewTable(),
	})
}

func (b *BigTwo) AddRound(r Round) *BigTwo {
	b.Rounds = append(b.Rounds, r)
	return b
}

func (b *BigTwo) Run() {
	for _, r := range b.Rounds {
		r(b)
	}
}
