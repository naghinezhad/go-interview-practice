package main

import "fmt"

type Employee struct {
	ID     int
	Name   string
	Age    int
	Salary float64
}

type Manager struct {
	Employees []Employee
}

// AddEmployee adds a new employee to the manager's list.
func (m *Manager) AddEmployee(e Employee) {
    
    if m.Employees == nil {
        m.Employees = []Employee{}
    }
    
	m.Employees = append(m.Employees, e)
}

// RemoveEmployee removes an employee by ID from the manager's list.
func (m *Manager) RemoveEmployee(id int) {
    
    if m.Employees == nil || len(m.Employees) == 0 {
	    return
	}
    
	for i, e := range m.Employees {
	    if e.ID == id {
	        m.Employees = append(m.Employees[:i], m.Employees[i+1:]...)
	        return
	    }
	}
}

// GetAverageSalary calculates the average salary of all employees.
func (m *Manager) GetAverageSalary() float64 {
	nbrSalary, totalSalary := 0.0, 0.0
	
	if m.Employees == nil || len(m.Employees) == 0 {
	    return 0.0
	}
	
	for _, e := range m.Employees {
	    totalSalary += e.Salary
	    nbrSalary += 1
	}
	
	return (totalSalary/nbrSalary)
}

// FindEmployeeByID finds and returns an employee by their ID.
func (m *Manager) FindEmployeeByID(id int) *Employee {
    
    if m.Employees == nil || len(m.Employees) == 0 {
	    return nil
	}
	
	for _, e := range m.Employees {
	    if e.ID == id {
	        return &e
	    }
	}
	return nil
}

func main() {
	manager := Manager{}
	manager.AddEmployee(Employee{ID: 1, Name: "Alice", Age: 30, Salary: 70000})
	manager.AddEmployee(Employee{ID: 2, Name: "Bob", Age: 25, Salary: 65000})
	manager.RemoveEmployee(1)
	averageSalary := manager.GetAverageSalary()
	employee := manager.FindEmployeeByID(2)

	fmt.Printf("Average Salary: %f\n", averageSalary)
	if employee != nil {
		fmt.Printf("Employee found: %+v\n", *employee)
	}
}
