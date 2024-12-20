package main

import "fmt"

type Employee struct {
	Name   string
	Skills []string
}

type Manager struct {
	Employee
	ReportingTeam []string
}

func (m Manager) printCompetencies() {
	fmt.Println("Manager Competencies:")
	for _, skill := range m.Skills {
		fmt.Println(skill)
	}
}

func main() {
	emp := Employee{
		Name:   "Alice",
		Skills: []string{"Leadership", "Management", "Communication"},
	}

	manager := Manager{
		Employee:      emp,
		ReportingTeam: []string{"Bob", "Charlie"},
	}

	manager.printCompetencies()
}
