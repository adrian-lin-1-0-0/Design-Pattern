package prescription

type Prescription struct {
	Name              string     `json:"name"`
	PotentialDiseases string     `json:"potentialDiseases"`
	Medicines         []Medicine `json:"medicines"`
	Usage             string     `json:"usage"`
}

func NewPrescription(name string, potentialDiseases string, medicines []Medicine, usage string) Prescription {
	return Prescription{
		Name:              name,
		PotentialDiseases: potentialDiseases,
		Medicines:         medicines,
		Usage:             usage,
	}
}
