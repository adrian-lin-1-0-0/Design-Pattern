package base

import "math/rand"

type (
	IDeck interface {
		Draw() ICard
		Shuffle()
		Len() int
	}

	BaseDeck []ICard
)

func (deck *BaseDeck) Draw() ICard {
	card := (*deck)[0]
	copy((*deck)[0:], (*deck)[1:])
	*deck = (*deck)[:len(*deck)-1]
	return card
}

func (deck *BaseDeck) Shuffle() {
	for i := range *deck {
		j := i + rand.Intn(len(*deck)-i)
		(*deck)[i], (*deck)[j] = (*deck)[j], (*deck)[i]
	}
}

func (deck *BaseDeck) Len() int {
	return len(*deck)
}

func (deck *BaseDeck) AddCard(card ICard) {
	*deck = append(*deck, card)
}

func NewBaseDeck() *BaseDeck {
	return &BaseDeck{}
}
