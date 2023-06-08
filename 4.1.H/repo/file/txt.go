package file

import (
	"bufio"
	"os"

	"4.1.H/domain/prescriber"
)

func Txt2DiagnosisHandlers(filePath string) ([]prescriber.DiagnosisHandler, error) {
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	diagnosisHandlers := []prescriber.DiagnosisHandler{}

	for _, disease := range lines {
		if diagnosisHandler := string2DiagnosisHandler(disease); diagnosisHandler != nil {
			diagnosisHandlers = append(diagnosisHandlers, *diagnosisHandler)
		}
	}

	return diagnosisHandlers, nil
}

func string2DiagnosisHandler(diagnosis string) *prescriber.DiagnosisHandler {
	switch diagnosis {
	case "COVID-19":
		return &prescriber.Covid19Diagnosis
	case "Attractive":
		return &prescriber.AttractiveDiagnosis
	case "SleepApneaSyndrome":
		return &prescriber.SleepApneaSyndromeDiagnosis
	}
	return nil
}
