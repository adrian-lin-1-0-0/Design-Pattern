package prescriber

import (
	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

type DiagnosisHandler struct {
	Match     func(p patient.Patient, symptoms []prescription.Symptom) bool
	Prescribe func() prescription.Prescription
}

func NewDiagnosisHandler(match func(patient.Patient, []prescription.Symptom) bool, p prescription.Prescription) DiagnosisHandler {
	return DiagnosisHandler{
		Match:     match,
		Prescribe: func() prescription.Prescription { return p },
	}
}
