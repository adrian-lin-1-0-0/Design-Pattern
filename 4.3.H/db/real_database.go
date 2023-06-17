package db

import (
	"fmt"
	"os/exec"
	"strconv"
	"strings"
)

type RealDatabase struct {
}

func NewRealDatabase() *RealDatabase {
	return &RealDatabase{}
}

func (db RealDatabase) GetEmployeeById(id int) Employee {
	cmd := exec.Command("sed", "-n", fmt.Sprintf("%d{p;q;}", id+1), "employees")

	output, err := cmd.Output()
	if err != nil {
		panic(err)
	}

	return textToRealEmployee(string(output))
}

func textToRealEmployee(text string) Employee {
	text = strings.TrimSuffix(text, "\n")
	row := strings.Split(text, " ")
	id, err := strconv.Atoi(row[0])
	if err != nil {
		return nil
	}
	name := row[1]
	age, err := strconv.Atoi(row[2])
	if err != nil {
		return nil
	}
	subordinatesIds := []int{}
	for _, subordinatesId := range strings.Split(row[3], ",") {
		subordinateId, err := strconv.Atoi(subordinatesId)
		if err != nil {
			break
		}
		subordinatesIds = append(subordinatesIds, subordinateId)
	}
	return NewRealEmployee(id, name, age, subordinatesIds)
}
