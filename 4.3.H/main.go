package main

import (
	"fmt"

	"4.3.H/db"
)

func main() {
	database := db.NewRealDatabase()
	employee := database.GetEmployeeById(2)
	fmt.Println(employee.Id(), employee.Name(), employee.Age())
	//2 fixiabis 15

	databaseProxy := db.NewProtectionRealDatabaseProxy()
	employeeProxy := databaseProxy.GetEmployeeById(2)
	/**
	You need to export PASSWORD=1qaz2wsx first,
	otherwise it will panic("Incorrect password") .
	*/

	subordinates := employeeProxy.GetSubordinates()
	for _, subordinate := range subordinates {
		fmt.Println(subordinate.Id(), subordinate.Name())
	}
	//1 waterball
	//2 fong
}
