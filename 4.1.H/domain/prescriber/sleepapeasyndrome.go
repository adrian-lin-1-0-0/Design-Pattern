package prescriber

import (
	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

type SleepApneaSyndromeDiagnosis struct{}

func (sas *SleepApneaSyndromeDiagnosis) Match(p patient.Patient, symptoms []prescription.Symptom) bool {
	hasSnore := false
	for _, symptom := range symptoms {
		if symptom == prescription.Snore {
			hasSnore = true
		}
	}
	return p.GetBMI() > 26 && hasSnore
}

func (sas *SleepApneaSyndromeDiagnosis) Handle() (prescription.Prescription, error) {
	return prescription.SleepApneaSyndrome, nil
}
