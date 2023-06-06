package repo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"4.1.H/domain/patient"
)

func Pares() {
	jsonFile, err := os.Open("patients.json")
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var patients []patient.Patient

	err = json.Unmarshal(byteValue, &patients)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(patients[0].Cases[0].Prescription.Medicines)
}
