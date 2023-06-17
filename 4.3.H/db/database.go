package db

type Database interface {
	GetEmployeeById(id int) Employee
}
