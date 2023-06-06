package prescription

type Prescription struct {
	Name              string
	PotentialDiseases string
	Medicines         []Medicine
	Usage             string
}

func NewPrescription(name string, potentialDiseases string, medicines []Medicine, usage string) Prescription {
	return Prescription{
		Name:              name,
		PotentialDiseases: potentialDiseases,
		Medicines:         medicines,
		Usage:             usage,
	}
}
