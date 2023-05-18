package message

import (
	"big2/pkg/card"
	"fmt"
	"strings"
)

//0    1    2    3    4    5     6    7    8    9    10   11   12
//C[3] C[4] S[7] S[8] H[9] D[10] S[J] D[Q] H[Q] D[A] S[A] D[2] H[2]

func CardsWithIdxToString(cards []card.Card) string {
	if len(cards) == 0 {
		return ""
	}
	idxLine := ""
	cardLine := ""
	for idx, c := range cards {
		if idx != len(cards)-1 {
			cardStr := c.String() + " "
			idxLine += padSpace(fmt.Sprintf("%d", idx), len(cardStr))
			cardLine += cardStr
		}
	}
	idxLine += fmt.Sprintf("%d", len(cards)-1)
	cardLine += cards[len(cards)-1].String()
	return idxLine + "\n" + cardLine
}

func padSpace(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(" ", length-len(str))
}

func CardsToString(cards []card.Card) string {
	cardStings := make([]string, len(cards))
	for idx, c := range cards {
		cardStings[idx] = c.String()
	}
	return strings.Join(cardStings, " ")
}
