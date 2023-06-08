package prescriber

import (
	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

var AttractiveDiagnosis = NewDiagnosisHandler(
	func(p patient.Patient, symptoms []prescription.Symptom) bool {
		hasSnore := false
		for _, symptom := range symptoms {
			if symptom == prescription.Snore {
				hasSnore = true
			}
		}
		return p.Gender == patient.Female && hasSnore && p.Age == 18
	},
	prescription.Attractive,
)
