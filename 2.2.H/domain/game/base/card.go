package base

type (
	ICard interface {
		Compare(any) Ord
	}

	BaseCard struct {
		Rank CardDenomination
		Suit CardDenomination
	}
)

func (card *BaseCard) Compare(other any) Ord {
	otherCard, _ := other.(*BaseCard)
	if card.Rank.Compare(&otherCard.Rank) == EQ {
		return card.Suit.Compare(&otherCard.Suit)
	}
	return card.Rank.Compare(&otherCard.Rank)
}
