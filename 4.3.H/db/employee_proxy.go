package db

type VirtualRealEmployeeProxy struct {
	*RealEmployee
	database Database
}

func NewEmployeeProry(employee Employee) *VirtualRealEmployeeProxy {
	return &VirtualRealEmployeeProxy{
		RealEmployee: employee.(*RealEmployee),
	}
}

func (e *VirtualRealEmployeeProxy) GetSubordinates() []Employee {
	if e.database == nil {
		e.database = NewRealDatabase()
	}
	subordinates := []Employee{}
	for _, subordinateId := range e.subordinatesIds {
		subordinates = append(subordinates, e.database.GetEmployeeById(subordinateId))
	}
	return subordinates
}
