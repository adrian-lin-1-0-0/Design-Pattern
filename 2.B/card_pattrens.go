package big2

//Single
//Pair
//Straight
//Full house

type cardPattenName string

const (
	Single    cardPattenName = "Single"
	Pair      cardPattenName = "Pair"
	Straight  cardPattenName = "Straight"
	FullHouse cardPattenName = "FullHouse"
)

type CardPattens struct {
	Name        cardPattenName
	Cards       []Card
	GreaterThan func(*CardPattens) bool
}

type IdentifyCardPatten struct {
	identify func([]Card) *CardPattens
	Next     *IdentifyCardPatten
}

func (i *IdentifyCardPatten) Identify(cards []Card) *CardPattens {

	cardPattern := i.identify(cards)
	if cardPattern != nil {
		return cardPattern
	}

	if i.Next != nil {
		return i.Next.Identify(cards)
	}

	return nil
}

func IdentifySingle(cards []Card) *CardPattens {
	if len(cards) != 1 {
		return nil
	}
	return &CardPattens{
		Name:  Single,
		Cards: cards,
	}
}

func IdentifyPair(cards []Card) *CardPattens {
	if len(cards) != 2 {
		return nil
	}

	if cards[0].Rank != cards[1].Rank {
		return nil
	}

	return &CardPattens{
		Name:  Pair,
		Cards: cards,
	}
}

type SinglePattern struct {
	cards []Card
}

func (s *SinglePattern) GreaterThan(p *CardPattens) bool {
	return s.cards[0].Rank > p.Cards[0].Rank
}

func (s *SinglePattern) Identify(cards []Card) bool {
	return len(cards) == 1
}
