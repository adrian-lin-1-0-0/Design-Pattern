package prescription

type Medicine struct {
	Name string
}

func NewMedicine(name string) Medicine {
	return Medicine{
		Name: name,
	}
}
