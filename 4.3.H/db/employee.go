package db

type Employee interface {
	Id() int
	Name() string
	Age() int
	GetSubordinates() []Employee
}

func NewBaseEmployee(id int, name string, age int, subordinatesIds []int) *BaseEmployee {
	return &BaseEmployee{
		id:              id,
		name:            name,
		age:             age,
		subordinatesIds: subordinatesIds,
	}
}

type BaseEmployee struct {
	id              int
	name            string
	age             int
	subordinatesIds []int
}

func (e *BaseEmployee) Id() int {
	return e.id
}

func (e *BaseEmployee) Name() string {
	return e.name
}

func (e *BaseEmployee) Age() int {
	return e.age
}

func (e *BaseEmployee) GetSubordinates() []Employee {
	return nil
}
