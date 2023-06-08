package mem

import (
	"sync"

	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

type database struct {
	Patients *patient.PatientDatabase
	mutex    sync.RWMutex
}

var DB *database

func (db *database) InsertPatient(patient *patient.Patient) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Patients.InsertPatient(patient)
}

func (db *database) InsertCaseById(id string, newCase prescription.Case) {
	db.mutex.Lock()
	defer db.mutex.Unlock()
	db.Patients.InsertCaseById(id, newCase)
}

func (db *database) GetPatientById(id string) (*patient.Patient, bool) {
	db.mutex.RLock()
	defer db.mutex.RUnlock()
	return db.Patients.GetPatientById(id)
}

func init() {
	DB = &database{
		Patients: patient.NewPatientDatabase(),
	}
}
