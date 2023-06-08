package usecases

import "4.1.H/domain/prescription"

func str2Symptoms(symptoms []string) []prescription.Symptom {
	var result []prescription.Symptom
	for _, symptom := range symptoms {
		result = append(result, prescription.Symptom(symptom))
	}
	return result
}
