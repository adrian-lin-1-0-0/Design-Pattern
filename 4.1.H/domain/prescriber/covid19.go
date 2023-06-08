package prescriber

import (
	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

var Covid19Diagnosis = NewDiagnosisHandler(
	func(p patient.Patient, symptoms []prescription.Symptom) bool {
		hasCough := false
		hasSneeze := false
		hasHeadache := false
		for _, symptom := range symptoms {
			if symptom == prescription.Cough {
				hasCough = true
			}
			if symptom == prescription.Sneeze {
				hasSneeze = true
			}
			if symptom == prescription.Headache {
				hasHeadache = true
			}
		}
		return hasCough && hasSneeze && hasHeadache
	},
	prescription.Covid19,
)
