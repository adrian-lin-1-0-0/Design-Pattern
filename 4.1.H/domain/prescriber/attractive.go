package prescriber

import (
	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

type AttractiveDiagnosis struct{}

func (ad *AttractiveDiagnosis) Match(p patient.Patient, symptoms []prescription.Symptom) bool {
	hasSnore := false
	for _, symptom := range symptoms {
		if symptom == prescription.Snore {
			hasSnore = true
		}
	}
	return p.Gender == patient.Female && hasSnore
}

func (ad *AttractiveDiagnosis) Handle() (prescription.Prescription, error) {
	return prescription.Attractive, nil
}
