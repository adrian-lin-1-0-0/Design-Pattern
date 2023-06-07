package prescriber

import (
	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

type DiagnosisHandler struct {
	next Diagnosis
}

type Diagnosis interface {
	Handle() (prescription.Prescription, error)
	Match(patient patient.Patient, symptoms []prescription.Symptom) bool
	Add(diagnosis Diagnosis)
	Next() Diagnosis
}

func (dh *DiagnosisHandler) Handle(p patient.Patient, symptoms []prescription.Symptom) (prescription.Prescription, error) {
	if dh.next == nil {
		return prescription.Prescription{}, nil
	}
	if dh.Next().Match(p, symptoms) {
		return dh.Next().Handle()
	}
	return dh.Next().Handle()
}

func (dh *DiagnosisHandler) Add(diagnosis Diagnosis) {
	if dh.next == nil {
		dh.next = diagnosis
		return
	}
	dh.Next().Add(diagnosis)
}

func (dh *DiagnosisHandler) Next() Diagnosis {
	return dh.next
}

func t() {
	diagnosis := &DiagnosisHandler{}
	diagnosis.Add(&Covid19Diagnosis{})
	diagnosis.Add(&AttractiveDiagnosis{})
	diagnosis.Add(&SleepApneaSyndromeDiagnosis{})
}
