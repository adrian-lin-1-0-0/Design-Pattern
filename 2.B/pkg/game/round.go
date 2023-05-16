package big2

import (
	"big2/pkg/card"
	"big2/pkg/card/patterns"
	"big2/pkg/player"
)

type Round interface {
	Run()
}

type firstRound struct {
	players []*player.Player
	deck    *Deck
}

type FirstRoundOptions struct {
	Players []*player.Player
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
	players           []*player.Player
	table             *Table
	cardPatternsChain *patterns.CardPatternsChain
}

type PlayRoundOptions struct {
	Players           []*player.Player
	Table             *Table
	CardPatternsChain *patterns.CardPatternsChain
}

func NewPlayRound(opts *PlayRoundOptions) *playRound {
	return &playRound{
		players:           opts.Players,
		table:             opts.Table,
		cardPatternsChain: opts.CardPatternsChain,
	}
}

func (r *playRound) Run() {
	var err error
	var cardPattern patterns.CardPattern

	for {
		for _, p := range r.players {
			for {
				p.Begin()
				topPlay := p.Play()
				if topPlay == nil {
					//pass
					goto Commit
				}

				cardPattern, err = r.cardPatternsChain.ToPattern(topPlay)
				if err != nil {
					//TODO
					// notify IllegalPlay
					goto Rollback
				}

				if cardPattern.GetName() != r.table.TopPlay.GetName() {
					goto Rollback
				}

				if !cardPattern.GreaterThan(r.table.TopPlay) {
					goto Rollback
				}

				r.table.TopPlay = cardPattern
				r.table.TopPlayer = p

			Commit:
				{
					p.Commit()
					if len(p.HandCards()) == 0 {
						//TODO
						// notify Win
						return
					}
					break
				}
			Rollback:
				{
					p.Rollback()
				}
			}

		}
	}
}

func IslegitimatePlay(play []card.Card) bool {
	return true
}
