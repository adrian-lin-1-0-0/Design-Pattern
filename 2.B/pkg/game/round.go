package game

import (
	"big2/pkg/card/patterns"
	"big2/pkg/notify/message"
	"big2/pkg/player"
	"fmt"
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

	for _, p := range r.players {
		p.SortHandCards()
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
			fmt.Fprintf(p.Writer, message.YourTurn, p.Name)
			for {
				p.Begin()
				topPlay := p.Play()
				if topPlay == nil && r.table.TopPlay != nil {
					if r.table.TopPlayer == p {
						fmt.Fprint(p.Writer, message.CantPassInNewRound)
						goto Rollback
					}
					fmt.Fprintf(p.Writer, message.PlayerPass, p.Name)
					goto Commit
				}

				cardPattern, err = r.cardPatternsChain.ToPattern(topPlay)
				if err != nil {
					goto IllegalPlay
				}

				if r.table.TopPlay == nil {
					r.table.TopPlay = cardPattern
					r.table.TopPlayer = p
					goto Commit
				}

				if cardPattern.GetName() != r.table.TopPlay.GetName() {
					goto IllegalPlay
				}

				if !cardPattern.GreaterThan(r.table.TopPlay) {
					goto IllegalPlay
				}

				r.table.TopPlay = cardPattern
				r.table.TopPlayer = p
			Commit:
				fmt.Fprintf(p.Writer, message.PlayerPlay, p.Name, cardPattern.GetName(), message.CardsToString(topPlay))
				p.Commit()
				if len(p.HandCards()) == 0 {
					fmt.Fprintf(p.Writer, message.GameOver, p.Name)
					return
				}
				break
			IllegalPlay:
				fmt.Fprint(p.Writer, message.IllegalPlay)

			Rollback:
				p.Rollback()
			}
		}
		r.table.TopPlay = nil

	}
}
