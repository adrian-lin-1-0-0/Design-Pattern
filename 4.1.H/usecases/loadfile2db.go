package usecases

import (
	"4.1.H/repo/file"
	"4.1.H/repo/mem"
)

func LoadFile2DB(filePath string) error {
	patients, err := file.Json2Patients(filePath)
	if err != nil {
		return err
	}

	for _, patient := range patients {
		mem.DB.InsertPatient(&patient)
	}

	return nil
}
