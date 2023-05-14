package big2

type Round interface {
	Play()
}

type firstRound struct {
	players []*Player
	deck    *Deck
}

type FirstRoundOptions struct {
	Players []*Player
	Deck    *Deck
}

func (r *firstRound) Play() {
	for _, p := range r.players {
		p.NamePlayer()
	}

	for r.deck.Len() > 0 {
		for _, p := range r.players {
			p.DealtHandCards(r.deck.Deal())
		}
	}
}

func NewFirstRound(opts *FirstRoundOptions) *firstRound {
	return &firstRound{
		players: opts.Players,
		deck:    opts.Deck,
	}
}

type playRound struct {
	players []*Player
	table   *Table
}

func (r *playRound) Play() {
	// TODO
}
