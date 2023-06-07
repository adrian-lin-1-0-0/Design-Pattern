package usecases

import (
	"4.1.H/domain/prescription"
	"4.1.H/repo/mem"
)

func PrescriptionDemand(id string, symptoms []string) (prescription.Prescription, error) {
	mem.DB.GetPatientById(id)
}
