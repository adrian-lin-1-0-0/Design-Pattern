package db

type RealEmployee struct {
	*BaseEmployee
}

func NewRealEmployee(id int, name string, age int, subordinatesIds []int) *RealEmployee {
	return &RealEmployee{
		BaseEmployee: NewBaseEmployee(id, name, age, subordinatesIds),
	}
}

func (e *RealEmployee) GetSubordinates() []Employee {
	//假裝已經從資料庫取得資料
	return nil
}
