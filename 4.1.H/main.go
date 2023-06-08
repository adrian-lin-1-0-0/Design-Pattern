package main

import (
	"fmt"
	"log"

	"4.1.H/domain/prescription"
	"4.1.H/usecases"
)

func PrintPrescription() func(prescription.Prescription) error {
	return func(p prescription.Prescription) error {
		fmt.Println(p)
		return nil
	}
}

func SavePrescription(filePath string) func(prescription.Prescription) error {
	return func(p prescription.Prescription) error {
		usecases.Prescription2File(p, filePath)
		return nil
	}
}

func main() {
	ps := usecases.NewPrescribeSystem(&usecases.PrescribeSystemOptions{
		DiagnosisSupportFile: "supports.txt",
		PatientJson:          "patients.json",
	})
	ps.Start()

	log.Println("[ main ] [ Request ] Start request: 1")
	ps.Request(usecases.NewPrescriptionDemand(
		"A123456789",
		[]string{"Cough", "Sneeze", "Headache"},
		PrintPrescription(),
	))
	log.Println("[ main ] [ Request ] Start request: 2")

	ps.Request(usecases.NewPrescriptionDemand(
		"A123456789",
		[]string{"Cough", "Sneeze", "Headache"},
		SavePrescription("prescription.json"),
	))

	log.Println("[ main ] [ Request ] Start request: 3")

	ps.Request(usecases.NewPrescriptionDemand(
		"A123456789",
		[]string{"Cough", "Sneeze", "Headache"},
		SavePrescription("prescription.csv"),
	))

	ps.Wait()
}
