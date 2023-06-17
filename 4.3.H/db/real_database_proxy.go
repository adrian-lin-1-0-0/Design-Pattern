package db

import "os"

const PASSWORD string = "1qaz2wsx"

type ProtectionRealDatabaseProxy struct {
	*RealDatabase
}

func NewProtectionRealDatabaseProxy() *ProtectionRealDatabaseProxy {
	return &ProtectionRealDatabaseProxy{NewRealDatabase()}
}

func (db ProtectionRealDatabaseProxy) GetEmployeeById(id int) Employee {
	if os.Getenv("PASSWORD") != PASSWORD {
		panic("Incorrect password")
	}

	employee := db.RealDatabase.GetEmployeeById(id)
	if employee == nil {
		return nil
	}
	return NewEmployeeProry(employee)
}
