package prescription

type Medicine struct {
	Name string `json:"name"`
}

func NewMedicine(name string) Medicine {
	return Medicine{
		Name: name,
	}
}
