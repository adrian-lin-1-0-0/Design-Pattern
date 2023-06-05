package prescription

type Case struct {
	Symptoms     []Symptom
	Time         string
	Prescription Prescription
}

func NewCase(symptoms []Symptom, time string, prescription Prescription) Case {
	return Case{
		Symptoms:     symptoms,
		Time:         time,
		Prescription: prescription,
	}
}
