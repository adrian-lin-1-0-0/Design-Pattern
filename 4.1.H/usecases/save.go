package usecases

import (
	"strings"

	"4.1.H/domain/prescriber"
	"4.1.H/domain/prescription"
	"4.1.H/repo/file"
	"4.1.H/repo/mem"
)

func Prescription2File(filePath string, p prescription.Prescription) error {

	var data []byte
	var err error
	if strings.Contains(filePath, ".json") {
		data, err = file.Prescription2Json(p)
		if err != nil {
			return err
		}
	}
	if strings.Contains(filePath, ".csv") {
		data, err = file.Prescription2Csv(p)
		if err != nil {
			return err
		}
	}

	return file.Byte2File(filePath, data)
}
func Prescriber2Mem(p *prescriber.Prescriber) {
	mem.Prescriber = p
}
