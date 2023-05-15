package big2

type BigTwo struct {
	rounds []Round
}

type BigTwoOptions struct {
	PlayerCount int
}

func NewBigTwo(opts *BigTwoOptions) *BigTwo {
	players := make([]*Player, opts.PlayerCount)
	for i := 0; i < opts.PlayerCount; i++ {
		players[i] = NewPlayer(&PlayerOptions{
			Core: NewHumanPlayer(nil),
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
