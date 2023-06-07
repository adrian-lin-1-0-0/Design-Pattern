package usecases

import (
	"4.1.H/domain/prescription"
	"4.1.H/repo/file"
)

func Prescription2File(filePath string, p prescription.Prescription) error {
	data, err := file.Prescription2Json(p)
	if err != nil {
		return err
	}
	return file.Json2File(filePath, data)
}
