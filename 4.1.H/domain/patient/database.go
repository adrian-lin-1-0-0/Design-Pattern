package patient

import "4.1.H/domain/prescription"

type PatientDatabase struct {
	Patients map[string]*Patient
}

func NewPatientDatabase() *PatientDatabase {
	return &PatientDatabase{
		Patients: make(map[string]*Patient),
	}
}

func (db *PatientDatabase) InsertPatient(patient *Patient) {
	db.Patients[patient.Id] = patient
}

func (db *PatientDatabase) InsertCaseById(id string, newCase prescription.Case) {
	patient, ok := db.Patients[id]
	if !ok {
		return
	}
	patient.AddCase(newCase)
}

func (db *PatientDatabase) GetPatientById(id string) (*Patient, bool) {
	patient, ok := db.Patients[id]
	return patient, ok
}
