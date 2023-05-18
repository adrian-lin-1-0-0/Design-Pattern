package card

import (
	"errors"
	"regexp"
	"strings"
)

//S[A]
func StringToCard(s string) Card {
	return Card{
		Rank: stringToCardRank(s),
		Suit: stringToCardSuit(s[0]),
	}
}

func stringToCardSuit(s byte) Suit {
	switch s {
	case 'C':
		return Clubs
	case 'D':
		return Diamonds
	case 'H':
		return Hearts
	case 'S':
		return Spades
	}
	return Clubs
}

func stringToCardRank(s string) Rank {
	rank, err := findStringRank(s)
	if err != nil {
		panic(err)
	}

	switch rank {
	case "3":
		return Three
	case "4":
		return Four
	case "5":
		return Five
	case "6":
		return Six
	case "7":
		return Seven
	case "8":
		return Eight
	case "9":
		return Nine
	case "10":
		return Ten
	case "J":
		return Jack
	case "Q":
		return Queen
	case "K":
		return King
	case "A":
		return Ace
	case "2":
		return Two
	default:
		panic("unreachable")
	}
}

func findStringRank(s string) (string, error) {
	regex := regexp.MustCompile(`\[(.*?)\]`)
	matches := regex.FindStringSubmatch(s)
	if len(matches) > 0 {
		return matches[1], nil
	}
	return "", errors.New("no match")
}

func StringToCards(s string) []Card {
	var cards []Card
	cardStrings := strings.Split(s, " ")

	for _, cardString := range cardStrings {
		cards = append(cards, StringToCard(cardString))
	}
	//reverse the cards
	for i, j := 0, len(cards)-1; i < j; i, j = i+1, j-1 {
		cards[i], cards[j] = cards[j], cards[i]
	}
	return cards
}
