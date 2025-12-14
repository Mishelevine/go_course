package main

import "fmt"

type Employee interface {
	GetSalary() int
}

type Director struct {
	name   string
	salary int
}

type Worker struct {
	name       string
	timeWorked int
}

func (d Director) GetSalary() int {
	return d.salary
}

func (w Worker) GetSalary() int {
	return w.timeWorked * 200
}

func AvgSalary(staff []Employee) int {
	var result int
	
	for _, employee := range staff {
		salary := employee.GetSalary()
		result += salary
	}

	return result / len(staff)
}

func task2_6() {
	team := make([]Employee, 0, 3)
	team = append(team, Director{"Иванов", 50000},
		Worker{"Петров", 40},
		Worker{"Сидоров", 300})
	fmt.Printf("Средняя зарплата в компании равна %d\n", AvgSalary(team))
}
