package main

import (
	"4.1.H/usecases"
)

func main() {
	usecases.Text2Prescriber("supports.txt")
	usecases.LoadFile2DBPanic("patients.json")
	p, _ := usecases.PrescriptionDemand("A123456789", []string{"Cough", "Sneeze", "Headache"})
	usecases.Prescription2File("prescription.csv", p)
}
