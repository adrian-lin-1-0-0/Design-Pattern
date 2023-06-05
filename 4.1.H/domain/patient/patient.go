package patient

import "4.1.H/domain/prescription"

type Patient struct {
	Id     string
	Name   string
	Age    int
	Gender Gender
	Height float32
	Weight float32
	Cases  []prescription.Case
}

func (p *Patient) AddCase(newCase prescription.Case) {
	p.Cases = append(p.Cases, newCase)
}

func (p *Patient) GetBMI() float32 {
	return p.Weight / (p.Height * p.Height)
}
