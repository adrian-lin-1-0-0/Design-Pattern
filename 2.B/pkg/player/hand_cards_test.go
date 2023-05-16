package player

import (
	"big2/pkg/card"
	"testing"
)

func TestHandCards_AddCard(t *testing.T) {
	handCards := NewHandCards()
	c3 := card.Card{Suit: card.Clubs, Rank: card.Three}
	c4 := card.Card{Suit: card.Clubs, Rank: card.Four}
	handCards.AddCard(c3)
	handCards.AddCard(c4)

	if handCards.GetCards()[0] != c3 {
		t.Error("AddCard() failed")
	}

	if handCards.GetCards()[1] != c4 {
		t.Error("AddCard() failed")
	}
}

func TestHandCards_Trans_Rollback(t *testing.T) {
	handCards := NewHandCards()
	c3 := card.Card{Suit: card.Clubs, Rank: card.Three}
	c4 := card.Card{Suit: card.Clubs, Rank: card.Four}
	handCards.AddCard(c3)
	handCards.AddCard(c4)

	handCards.Begin()

	c5 := card.Card{Suit: card.Clubs, Rank: card.Five}
	handCards.AddCard(c5)

	handCards.Rollback()

	if len(handCards.GetCards()) != 2 {
		t.Error("Rollback() failed")
	}

	handCards.AddCard(c5)

	if handCards.GetCards()[2] != c5 {
		t.Error("AddCard() failed")
	}
}

func TestHandCards_Trans_Commit(t *testing.T) {
	handCards := NewHandCards()
	c3 := card.Card{Suit: card.Clubs, Rank: card.Three}
	c4 := card.Card{Suit: card.Clubs, Rank: card.Four}
	handCards.AddCard(c3)
	handCards.AddCard(c4)

	handCards.Begin()

	c5 := card.Card{Suit: card.Clubs, Rank: card.Five}
	handCards.AddCard(c5)

	handCards.Commit()

	if handCards.GetCards()[2] != c5 {
		t.Error("Commit() failed")
	}

	c6 := card.Card{Suit: card.Clubs, Rank: card.Six}
	handCards.AddCard(c6)

	if handCards.GetCards()[3] != c6 {
		t.Error("AddCard() failed")
	}
}

func TestHandCards_Sort(t *testing.T) {

	tests := []struct {
		name      string
		handCards []card.Card
		want      []card.Card
	}{
		{
			name: "sort 3,6,4,5 to 3,4,5,6",
			handCards: []card.Card{
				{Rank: card.Three, Suit: card.Spades},
				{Rank: card.Six, Suit: card.Spades},
				{Rank: card.Four, Suit: card.Spades},
				{Rank: card.Five, Suit: card.Spades},
			},
			want: []card.Card{
				{Rank: card.Three, Suit: card.Spades},
				{Rank: card.Four, Suit: card.Spades},
				{Rank: card.Five, Suit: card.Spades},
				{Rank: card.Six, Suit: card.Spades},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := NewHandCards()
			h.SetCards(tt.handCards)
			h.Sort()
			for i := 0; i < len(h.GetCards()); i++ {
				if h.GetCards()[i] != tt.want[i] {
					t.Errorf("Sort() = %v, want %v", h.GetCards(), tt.want)
				}
			}
		})
	}
}
