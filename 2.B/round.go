package big2

type Round interface {
	Run()
}

type firstRound struct {
	players []*Player
	deck    *Deck
}

type FirstRoundOptions struct {
	Players []*Player
	Deck    *Deck
}

func (r *firstRound) Run() {
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

func (r *playRound) Run() {
	for {
		for _, p := range r.players {
			for {
				p.Begin()
				topPlay := p.Play()
				if topPlay == nil {
					//pass
					goto Commit
				}

				if IslegitimatePlay(topPlay) {
					r.table.TopPlay = topPlay
					r.table.TopPlayer = p
					goto Commit
				}
				p.Rollback()
				//TODO
				//此牌型不合法，請再嘗試一次。
				continue
			Commit:
				p.Commit()
				break

			}

		}
	}
}

func IslegitimatePlay(play []Card) bool {
	return true
}
