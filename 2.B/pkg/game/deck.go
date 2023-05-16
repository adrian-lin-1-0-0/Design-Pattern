package big2

import (
	"big2/pkg/card"
	"math/rand"
	"time"
)

type Deck struct {
	cards []card.Card
}

type DeckOptions struct {
	Shuffle bool
	Cards   []card.Card
}

func NewDeck(opts *DeckOptions) *Deck {

	var cards []card.Card

	if opts == nil {
		cards = make([]card.Card, 52)
		for i := 0; i < 52; i++ {
			cards[i] = card.Card{
				Suit: card.Suit(i / 13),
				Rank: card.Rank(i % 13)}
		}
		return (&Deck{cards: cards}).Shuffle()

	}

	cards = opts.Cards
	d := &Deck{cards: cards}
	if opts.Shuffle {
		d.Shuffle()
	}
	return d
}

func (d *Deck) Shuffle() *Deck {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(d.cards), func(i, j int) {
		d.cards[i], d.cards[j] = d.cards[j], d.cards[i]
	})
	return d
}

func (d *Deck) Deal() card.Card {
	card := d.cards[0]
	d.cards = d.cards[1:]
	return card
}

func (d *Deck) Len() int {
	return len(d.cards)
}
