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

func LoadFile2DBPanic(filePath string) {
	err := LoadFile2DB(filePath)
	if err != nil {
		panic(err)
	}
}
