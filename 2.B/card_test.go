package big2

import (
	"testing"
)

func TestCard_String(t *testing.T) {

	tests := []struct {
		name string
		card Card
		want string
	}{
		{"C3", Card{Clubs, Three}, "C[3]"},
		{"D4", Card{Diamonds, Four}, "D[4]"},
		{"H5", Card{Hearts, Five}, "H[5]"},
		{"S6", Card{Spades, Six}, "S[6]"},
		{"C7", Card{Clubs, Seven}, "C[7]"},
		{"D8", Card{Diamonds, Eight}, "D[8]"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.String(); got != tt.want {
				t.Errorf("Card.String() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCard_EqualTo(t *testing.T) {

	tests := []struct {
		name  string
		card  Card
		other Card
		want  bool
	}{
		{"C3 == C3", Card{Clubs, Three}, Card{Clubs, Three}, true},
		{"C3 != D3", Card{Clubs, Three}, Card{Diamonds, Three}, false},
		{"C3 != C4", Card{Clubs, Three}, Card{Clubs, Four}, false},
		{"C3 != D4", Card{Clubs, Three}, Card{Diamonds, Four}, false},
		{"C3 != H5", Card{Clubs, Three}, Card{Hearts, Five}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.card.EqualTo(&tt.other); got != tt.want {
				t.Errorf("Card.EqualTo() = %v, want %v", got, tt.want)
			}
		})
	}
}
