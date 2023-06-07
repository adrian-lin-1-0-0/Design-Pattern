package usecases

import (
	"4.1.H/domain/prescriber"
	"4.1.H/repo/file"
)

func CreatePrescriberWithfilePath(filePath string) (*prescriber.Prescriber, error) {
	diagnosisHandlers, err := file.Txt2DiagnosisHandlers(filePath)
	if err != nil {
		return nil, err
	}
	p := prescriber.NewPrescriber()

	for _, diagnosisHandler := range diagnosisHandlers {
		p.Add(diagnosisHandler)
	}

	return p, nil
}

func Text2Prescriber(filePath string) {
	p, err := CreatePrescriberWithfilePath(filePath)
	if err != nil {
		panic(err)
	}
	Prescriber2Mem(p)
}
