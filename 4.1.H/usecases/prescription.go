package usecases

import (
	"4.1.H/domain/prescriber"
	"4.1.H/domain/prescription"
	"4.1.H/repo/mem"
)

func PrescriptionDemand(id string, symptoms []string) (prescription.Prescription, error) {
	person, exist := mem.DB.GetPatientById(id)
	if !exist {
		return prescription.Prescription{}, nil
	}
	p, err := prescriber.NewDefaultPrescriber().Handle(*person, str2Symptoms(symptoms))
	if err != nil {
		return prescription.Prescription{}, err
	}
	return p, nil
}

func str2Symptoms(symptoms []string) []prescription.Symptom {
	var result []prescription.Symptom
	for _, symptom := range symptoms {
		result = append(result, prescription.Symptom(symptom))
	}
	return result
}
