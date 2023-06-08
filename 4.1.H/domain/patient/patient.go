package patient

import "4.1.H/domain/prescription"

type Patient struct {
	Id     string              `json:"id"`
	Name   string              `json:"name"`
	Age    int                 `json:"age"`
	Gender Gender              `json:"gender"`
	Height float32             `json:"height"`
	Weight float32             `json:"weight"`
	Cases  []prescription.Case `json:"cases"`
}

func (p *Patient) AddCase(newCase prescription.Case) {
	p.Cases = append(p.Cases, newCase)
}

func (p *Patient) GetBMI() float32 {
	return p.Weight / (p.Height * p.Height)
}
