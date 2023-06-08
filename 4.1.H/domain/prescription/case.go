package prescription

type Case struct {
	Symptoms     []Symptom    `json:"symptoms"`
	Time         string       `json:"time"`
	Prescription Prescription `json:"prescription"`
}

func NewCase(symptoms []Symptom, time string, prescription Prescription) Case {
	return Case{
		Symptoms:     symptoms,
		Time:         time,
		Prescription: prescription,
	}
}
