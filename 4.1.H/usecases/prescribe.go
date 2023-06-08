package usecases

import (
	"log"
	"sync"
	"time"

	"4.1.H/domain/prescription"
	"4.1.H/repo/mem"
)

type PrescribeSystem struct {
	requests chan PrescriptionDemand
	wg       sync.WaitGroup
}

type PrescriptionDemand struct {
	PatientID      string
	Symptoms       []string
	AfterPrescribe func(prescription.Prescription) error
}

type PrescribeSystemOptions struct {
	DiagnosisSupportFile string
	PatientJson          string
}

func NewPrescriptionDemand(patientID string, symptoms []string, afterPrescribe func(prescription.Prescription) error) PrescriptionDemand {
	return PrescriptionDemand{
		PatientID:      patientID,
		Symptoms:       symptoms,
		AfterPrescribe: afterPrescribe,
	}
}

func NewPrescribeSystem(opts *PrescribeSystemOptions) *PrescribeSystem {
	Text2Prescriber(opts.DiagnosisSupportFile)
	LoadFile2DBPanic(opts.PatientJson)

	return &PrescribeSystem{
		requests: make(chan PrescriptionDemand, 100),
	}
}

func (p *PrescribeSystem) Request(demand PrescriptionDemand) {
	p.wg.Add(1)
	p.requests <- demand
}

func (p *PrescribeSystem) Wait() {
	p.wg.Wait()
}

func (p *PrescribeSystem) Start() {

	go func() {
		for demand := range p.requests {
			log.Println("[ PrescribeSystem ] [ Prescribe ] Start prescribe: ", demand.PatientID)
			_prescription, _ := prescribe(demand.PatientID, demand.Symptoms)
			time.Sleep(5 * time.Second)
			demand.AfterPrescribe(_prescription)
			log.Println("[ PrescribeSystem ] [ Prescribe ] End prescribe: ", demand.PatientID)
			p.wg.Done()
		}
	}()
}

func prescribe(id string, strSymptoms []string) (prescription.Prescription, error) {
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
