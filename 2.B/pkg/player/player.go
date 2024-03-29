package player

import (
	"big2/pkg/card"
	"big2/pkg/notify/message"
	"errors"
	"fmt"
	"io"
	"os"
)

type PlayerCore interface {
	NamePlayer(*Player)
	Play(*Player) []int
}

type Player struct {
	Name      string
	handCards *HandCards
	core      PlayerCore
	Reader    io.Reader
	Writer    io.Writer
}

type PlayerOptions struct {
	Core   PlayerCore
	Reader io.Reader
	Writer io.Writer
}

func NewPlayer(opts *PlayerOptions) *Player {
	return &Player{
		core:      opts.Core,
		Reader:    opts.Reader,
		Writer:    opts.Writer,
		handCards: NewHandCards(),
	}
}

func NewDefaultPlayer() *Player {
	return NewPlayer(&PlayerOptions{
		Core:   NewHumanPlayer(),
		Reader: os.Stdin,
		Writer: os.Stdout,
	})
}

func (p *Player) Play() []card.Card {
	fmt.Fprintf(p.Writer, message.HandCards, message.CardsWithIdxToString(p.handCards.GetCards()))
	play := p.core.Play(p)
	if len(play) == 0 {
		panic(errors.New("play is empty"))
	}

	if len(play) == 1 && play[0] == -1 {
		return nil
	}
	playCards := getCardsByIdx(p.handCards.GetCards(), play)
	p.handCards.SetCards(removeCardsByIdx(p.handCards.GetCards(), play))

	return playCards
}

func (p *Player) NamePlayer() {
	p.core.NamePlayer(p)
}

func (p *Player) DealtHandCards(c card.Card) {
	p.handCards.AddCard(c)
}

func (p *Player) SortHandCards() {
	p.handCards.Sort()
}

func (p *Player) Begin() {
	p.handCards.Begin()
}

func (p *Player) Commit() {
	p.handCards.Commit()
}

func (p *Player) Rollback() {
	p.handCards.Rollback()
}

func (p *Player) HandCards() []card.Card {
	return p.handCards.GetCards()
}
