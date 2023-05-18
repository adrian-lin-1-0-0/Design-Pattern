package patterns

import (
	"big2/pkg/card"
	"testing"
)

func TestStraightPattern_Match(t *testing.T) {

	tests := []struct {
		name  string
		cards []card.Card

		want bool
	}{
		{
			name: "345A2(A2345)",
			cards: []card.Card{
				{Rank: card.Three, Suit: card.Spades},
				{Rank: card.Four, Suit: card.Spades},
				{Rank: card.Five, Suit: card.Spades},
				{Rank: card.Ace, Suit: card.Spades},
				{Rank: card.Two, Suit: card.Spades},
			},
			want: true,
		},
		{
			name: "AA345",
			cards: []card.Card{
				{Rank: card.Ace, Suit: card.Spades},
				{Rank: card.Ace, Suit: card.Hearts},
				{Rank: card.Three, Suit: card.Spades},
				{Rank: card.Four, Suit: card.Spades},
				{Rank: card.Five, Suit: card.Spades},
			},
			want: false,
		},
		{
			name: "34KA2(KA234)",
			cards: []card.Card{
				{Rank: card.Three, Suit: card.Spades},
				{Rank: card.Four, Suit: card.Spades},
				{Rank: card.King, Suit: card.Spades},
				{Rank: card.Ace, Suit: card.Spades},
				{Rank: card.Two, Suit: card.Spades},
			},
			want: true,
		},
		{
			name: "8910JQ",
			cards: []card.Card{
				{Rank: card.Eight, Suit: card.Spades},
				{Rank: card.Nine, Suit: card.Spades},
				{Rank: card.Ten, Suit: card.Spades},
				{Rank: card.Jack, Suit: card.Spades},
				{Rank: card.Queen, Suit: card.Spades},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StraightPattern{}
			if got := s.Match(tt.cards); got != tt.want {
				t.Errorf("StraightPattern.Match() = %v, want %v", got, tt.want)
			}
		})
	}
}
