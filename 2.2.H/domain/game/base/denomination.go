package base

type CardDenomination struct {
	Name  string
	Value int
}

func (cd *CardDenomination) Compare(other *CardDenomination) Ord {
	switch {
	case cd.Value < other.Value:
		return LT
	case cd.Value > other.Value:
		return GT
	default:
		return EQ
	}
}
