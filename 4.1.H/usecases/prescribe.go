package usecases

import (
	"time"

	"4.1.H/domain/prescription"
	"4.1.H/repo/mem"
)

func PrescriptionDemand(id string, strSymptoms []string) (prescription.Prescription, error) {
	person, exist := mem.DB.GetPatientById(id)
	if !exist {
		return prescription.Prescription{}, nil
	}
	symptoms := str2Symptoms(strSymptoms)
	p := mem.Prescriber.Handle(*person, symptoms)
	person.AddCase(prescription.NewCase(
		symptoms,
		time.Now().Format("2006-01-02 15:04:05"),
		p,
	))

	return p, nil
}

func str2Symptoms(symptoms []string) []prescription.Symptom {
	var result []prescription.Symptom
	for _, symptom := range symptoms {
		result = append(result, prescription.Symptom(symptom))
	}
	return result
}
