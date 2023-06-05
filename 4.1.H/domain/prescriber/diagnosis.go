package prescriber

import (
	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

type DiagnosisHandler struct {
	Next Diagnosis
}

type Diagnosis interface {
	Handle() (prescription.Prescription, error)
	Match(patient patient.Patient, symptoms []prescription.Symptom) bool
	Add(diagnosis Diagnosis)
	Next() Diagnosis
}

func (dh *DiagnosisHandler) Handle(p patient.Patient, symptoms []prescription.Symptom) (prescription.Prescription, error) {
	if dh.Next == nil {
		return prescription.Prescription{}, nil
	}
	if dh.Next.Match(p, symptoms) {
		return dh.Next.Handle()
	}
	return dh.Next.Handle()
}

func (dh *DiagnosisHandler) Add(diagnosis Diagnosis) {
	if dh.Next == nil {
		dh.Next = diagnosis
		return
	}
	dh.Next.Add(diagnosis)
}
