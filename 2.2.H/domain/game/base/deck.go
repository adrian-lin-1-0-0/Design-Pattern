package base

import "math/rand"

type (
	IDeck interface {
		Draw() ICard
		Shuffle()
		Len() int
		Set([]ICard)
		ToSlice() []ICard
	}

	BaseDeck []ICard
)

func (deck *BaseDeck) ToSlice() []ICard {
	return *deck
}

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

func (deck *BaseDeck) Set(cards []ICard) {
	*deck = cards
}

func NewBaseDeck() *BaseDeck {
	return &BaseDeck{}
}
