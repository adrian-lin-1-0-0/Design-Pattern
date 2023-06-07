package file

import (
	"encoding/json"
	"io/ioutil"
	"os"

	"4.1.H/domain/patient"
	"4.1.H/domain/prescription"
)

func Json2Patients(filePath string) ([]patient.Patient, error) {
	jsonFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	defer jsonFile.Close()
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var patients []patient.Patient

	err = json.Unmarshal(byteValue, &patients)
	if err != nil {
		return nil, err
	}
	return patients, nil
}

func Prescription2Json(p prescription.Prescription) ([]byte, error) {
	return json.Marshal(p)
}
