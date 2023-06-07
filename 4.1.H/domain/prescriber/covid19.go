package prescriber

import (
	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

type Covid19Diagnosis struct {
	DiagnosisHandler
}

func (cd *Covid19Diagnosis) Match(p patient.Patient, symptoms []prescription.Symptom) bool {
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
}

func (cd *Covid19Diagnosis) Handle() (prescription.Prescription, error) {
	return prescription.Covid19, nil
}
