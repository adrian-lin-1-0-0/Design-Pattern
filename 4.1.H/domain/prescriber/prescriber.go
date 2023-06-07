package prescriber

type Prescriber struct {
	DiagnosisHandler
}

func NewPrescriber() *Prescriber {
	diagnosis := &DiagnosisHandler{}
	return &Prescriber{*diagnosis}
}

func NewDefaultPrescriber() *Prescriber {
	p := NewPrescriber()
	p.Add(&Covid19Diagnosis{})
	p.Add(&AttractiveDiagnosis{})
	p.Add(&SleepApneaSyndromeDiagnosis{})

	return p
}
