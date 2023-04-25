package base

type (
	IHand interface {
		AddCard(ICard)
		PlayCard(int) ICard
		Len() int
		ToSlice() []ICard
	}

	BaseHand []ICard
)

func (hand *BaseHand) ToSlice() []ICard {
	return *hand
}

func (hand *BaseHand) AddCard(card ICard) {
	*hand = append(*hand, card)
}

func (hand *BaseHand) PlayCard(idx int) ICard {
	card := (*hand)[idx]
	copy((*hand)[idx:], (*hand)[idx+1:])
	*hand = (*hand)[:len(*hand)-1]
	return card
}

func (hand *BaseHand) Len() int {
	return len(*hand)
}
