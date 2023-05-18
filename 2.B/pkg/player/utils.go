package player

import (
	"big2/pkg/card"
	"bufio"
	"io"
	"strings"
)

func removeCardsByIdx(cards []card.Card, idx []int) []card.Card {
	emptyCard := card.Card{}

	for _, handIdx := range idx {
		cards[handIdx] = emptyCard
	}

	newCards := []card.Card{}
	for _, card := range cards {
		if card != emptyCard {
			newCards = append(newCards, card)
		}
	}

	return newCards
}

func getCardsByIdx(cards []card.Card, idx []int) []card.Card {
	cardsByIdx := []card.Card{}

	for _, handIdx := range idx {
		cardsByIdx = append(cardsByIdx, cards[handIdx])
	}

	return cardsByIdx
}

// read line including space
func readLine(ioReader io.Reader) string {
	reader := bufio.NewReader(ioReader)
	line, _ := reader.ReadString('\n')
	return strings.TrimRight(line, "\n")
}
