package base

type (
	ICard interface {
		Compare(any) Ord
	}
)
