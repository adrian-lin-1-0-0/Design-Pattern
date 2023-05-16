package big2

import (
	"big2/pkg/card"
	"fmt"
	"strings"
)

//0    1    2    3    4    5     6    7    8    9    10   11   12
//C[3] C[4] S[7] S[8] H[9] D[10] S[J] D[Q] H[Q] D[A] S[A] D[2] H[2]

func CardsToString(cards []card.Card) string {
	idxLine := ""
	cardLine := ""
	for idx, c := range cards {
		cardStr := c.String() + " "
		idxLine += padSpace(fmt.Sprintf("%d", idx), len(cardStr))
		cardLine += cardStr
	}
	return idxLine + "\n" + cardLine
}

func padSpace(str string, length int) string {
	if len(str) >= length {
		return str
	}
	return str + strings.Repeat(" ", length-len(str))
}
