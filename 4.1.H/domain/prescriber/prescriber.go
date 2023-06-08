package prescriber

import (
	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

type Prescriber struct {
	DiagnosisHandlers []DiagnosisHandler
}

func (p *Prescriber) Handle(patient patient.Patient, symptoms []prescription.Symptom) prescription.Prescription {
	for _, handler := range p.DiagnosisHandlers {
		if handler.Match(patient, symptoms) {
			return handler.Prescribe()
		}
	}
	return prescription.None
}

func (p *Prescriber) Add(handler DiagnosisHandler) {
	p.DiagnosisHandlers = append(p.DiagnosisHandlers, handler)
}

func NewDefaultPrescriber() *Prescriber {
	p := NewPrescriber()
	p.Add(Covid19Diagnosis)
	p.Add(AttractiveDiagnosis)
	p.Add(SleepApneaSyndromeDiagnosis)
	return p
}

func NewPrescriber() *Prescriber {
	return &Prescriber{}
}
