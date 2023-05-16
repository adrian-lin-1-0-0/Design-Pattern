package big2

import "big2/pkg/player"

type BigTwo struct {
	rounds []Round
}

type BigTwoOptions struct {
	PlayerCount int
}

func NewBigTwo(opts *BigTwoOptions) *BigTwo {
	players := make([]*player.Player, opts.PlayerCount)
	for i := 0; i < opts.PlayerCount; i++ {
		players[i] = player.NewPlayer(&player.PlayerOptions{
			Core: player.NewHumanPlayer(nil),
		})
	}

	firstRound := NewFirstRound(&FirstRoundOptions{
		Players: players,
		Deck:    NewDeck(nil),
	})

	return &BigTwo{
		rounds: []Round{firstRound},
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
