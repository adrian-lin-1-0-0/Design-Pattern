package patterns

import (
	"big2/pkg/card"
	"strings"
)

type StraightPattern struct {
	cards []card.Card
}

var (
	matchStraight = [...]string{
		"345678910JQKA2",
		"34KA2", //KA234
		"345A2", //A2345
		"3QKA2", //QKA23
	}
)

func MatchStraight(patternStr string) bool {
	for _, s := range matchStraight {
		if strings.Contains(s, patternStr) {
			return true
		}
	}
	return false
}

func cardToPatternStr(cards []card.Card) string {
	patternStr := ""
	for _, c := range cards {
		patternStr += c.Rank.String()
	}
	return patternStr
}

func (s *StraightPattern) GreaterThan(p CardPattern) bool {
	return s.cards[4].Rank > p.GetCards()[4].Rank
}

func (s *StraightPattern) Match(cards []card.Card) bool {
	return (len(cards) == 5) && MatchStraight(cardToPatternStr(cards))
}
