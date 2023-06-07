package main

import (
	"fmt"

	"4.1.H/repo/mem"
	"4.1.H/usecases"
)

func main() {
	filePath := "patients.json"

	err := usecases.LoadFile2DB(filePath)
	if err != nil {
		panic(err)
	}

	p, _ := mem.DB.GetPatientById("A123456789")
	fmt.Printf("%#v\n", p)

}
