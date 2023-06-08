package usecases

import (
	"strings"

	"4.1.H/domain/prescriber"
	"4.1.H/domain/prescription"
	"4.1.H/repo/file"
	"4.1.H/repo/mem"
)

func Prescription2File(p prescription.Prescription, filePath string) error {

	var data []byte
	var err error

	switch {
	case strings.Contains(filePath, ".json"):
		data, err = file.Prescription2Json(p)
		if err != nil {
			return err
		}
	case strings.Contains(filePath, ".csv"):
		data, err = file.Prescription2Csv(p)
		if err != nil {
			return err
		}
	default:
		return nil
	}

	/**
	Prescription2Json 跟 Prescription2Csv　都是

	func(prescription.Prescription)([]byte,error)

	但是沒有特邊開一個 func(prescription.Prescription)([]byte,error)　的 interface
	Q1:
	這樣還算是 strategy pattern 嗎？
	Q2:
	這層本來就在做雜事, 這樣寫有違反 OCP 嗎？
	*/

	return file.Byte2File(filePath, data)
}
func Prescriber2Mem(p *prescriber.Prescriber) {
	mem.Prescriber = p
}
