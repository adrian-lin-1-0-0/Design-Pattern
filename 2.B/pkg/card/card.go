package card

import (
	"fmt"
)

type Card struct {
	Suit Suit
	Rank Rank
}

type Suit int

const (
	Clubs Suit = iota
	Diamonds
	Hearts
	Spades
)

type Rank int

const (
	Three Rank = iota
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
	Two
)

// C : Clubs ♣
// D : Diamonds ♦
// H : Hearts ♥
// S : Spades ♠
var suitStrings = []string{"C", "D", "H", "S"}
var rankStrings = []string{"3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K", "A", "2"}

func (c *Card) compare(other *Card) int {

	switch {
	case c.Rank > other.Rank:
		return 1
	case c.Rank < other.Rank:
		return -1
	case c.Rank == other.Rank:
		if c.Suit > other.Suit {
			return 1
		}
		if c.Suit == other.Suit {
			return 0
		}
		return -1
	default:
		panic("unreachable")
	}
}

func (c *Card) LessThan(other *Card) bool {
	return c.compare(other) < 0
}

func (c *Card) GreaterThan(other *Card) bool {
	return c.compare(other) > 0
}

func (c *Card) EqualTo(other *Card) bool {
	return c.compare(other) == 0
}

func (c *Card) LessThanOrEqualTo(other *Card) bool {
	return c.compare(other) <= 0
}

func (c *Card) GreaterThanOrEqualTo(other *Card) bool {
	return c.compare(other) >= 0
}

func (c *Card) String() string {
	return fmt.Sprintf("%s[%s]", suitStrings[c.Suit], rankStrings[c.Rank])
}
