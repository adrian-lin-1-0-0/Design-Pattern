package message

import (
	"big2/pkg/card"
	"testing"
)

func TestCardsWithIdxToString(t *testing.T) {
	//345678910JQKA2
	cards := []card.Card{
		{Rank: card.Three, Suit: card.Spades},
		{Rank: card.Four, Suit: card.Spades},
		{Rank: card.Five, Suit: card.Spades},
		{Rank: card.Six, Suit: card.Spades},
		{Rank: card.Seven, Suit: card.Spades},
		{Rank: card.Eight, Suit: card.Spades},
		{Rank: card.Nine, Suit: card.Spades},
		{Rank: card.Ten, Suit: card.Spades},
		{Rank: card.Jack, Suit: card.Spades},
		{Rank: card.Queen, Suit: card.Spades},
		{Rank: card.King, Suit: card.Spades},
		{Rank: card.Ace, Suit: card.Spades},
		{Rank: card.Two, Suit: card.Spades},
	}

	got := CardsWithIdxToString(cards)
	want := "0    1    2    3    4    5    6    7     8    9    10   11   12\nS[3] S[4] S[5] S[6] S[7] S[8] S[9] S[10] S[J] S[Q] S[K] S[A] S[2]"

	if got != want {
		t.Errorf("CardsToString() = %v, want %v", got, want)
	}
}
