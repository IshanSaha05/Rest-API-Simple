package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/IshanSaha05/microservice/pkg/models"
	"github.com/gorilla/mux"
)

var employees []models.Employees

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	byteEmployee, err := io.ReadAll(r.Body)

	if err != nil {
		fmt.Println("Error while converting body to byte slice.")
		os.Exit(1)
	}

	var employee models.Employees

	err = json.Unmarshal(byteEmployee, &employee)

	if err != nil {
		fmt.Println("Error while Unmarshling.")
		os.Exit(1)
	}

	fmt.Printf("\nID: %d\nName: %s\nDesignation: %s\n", employee.ID, employee.Name, employee.Name)

	employees = append(employees, employee)

	fmt.Println("\nMessage: Creation Completed.")

	byteEmployee, err = json.Marshal(employee)

	if err != nil {
		fmt.Println("Error while Marshling.")
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteEmployee)
}

func GetAllEmployees(w http.ResponseWriter, r *http.Request) {
	fmt.Println("\nPrinting All the Employees.")
	fmt.Printf("---------------------------\n")

	for _, employee := range employees {
		fmt.Printf("ID: %d\tName: %s\tDesignation: %s\n", employee.ID, employee.Name, employee.Designation)
	}

	byteEmployee, err := json.Marshal(employees)

	if err != nil {
		fmt.Println("Error while Marshling.")
		os.Exit(1)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(byteEmployee)
}

func GetEmployeeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["employee_id"], 0, 0)

	if err != nil {
		fmt.Println("Error while converting from string to integer.")
		os.Exit(1)
	}

	for _, employee := range employees {
		if employee.ID == int(id) {
			fmt.Printf("\nEmployee with ID: %d found.\n", id)
			fmt.Printf("Employee Details of ID: %d\n--------------------------\n", id)
			fmt.Printf("ID: %d\tName: %s\tDesignation: %s\n", employee.ID, employee.Name, employee.Designation)

			var byteEmployee []byte

			byteEmployee, err = json.Marshal(employee)

			if err != nil {
				fmt.Println("Error while Marshling.")
				os.Exit(1)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(byteEmployee)

			return
		}
	}

	fmt.Printf("\nMessage: No employee with ID: %d found.\n", id)
}

func UpdateEmployeeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["employee_id"], 0, 0)

	if err != nil {
		fmt.Println("Error while converting string to integer.")
		os.Exit(1)
	}

	for index, employee := range employees {
		if employee.ID == int(id) {
			fmt.Printf("\nEmployee with ID: %d found.\n", id)
			fmt.Printf("Employee details of ID: %d before any updates.\n", id)
			fmt.Printf("ID: %d\tName: %s\tDesignation: %s\n", employee.ID, employee.Name, employee.Designation)

			byteEmployee, err := io.ReadAll(r.Body)

			if err != nil {
				fmt.Println("Error while converting body to byte slice.")
				os.Exit(1)
			}

			var newEmployee models.Employees
			newEmployee.ID = employee.ID

			err = json.Unmarshal(byteEmployee, &newEmployee)

			if err != nil {
				fmt.Println("Error while Unmarshling.")
				os.Exit(1)
			}

			/*
				// Not working
				employee.Name = newEmployee.Name
				employee.Designation = newEmployee.Designation

					fmt.Printf("\nEmployee details of ID: %d after any updates.\n", employee.ID)
				fmt.Printf("ID: %d\tName: %s\tDesignation: %s\n", employee.ID, employee.Name, employee.Designation)

				byteEmployee, err = json.Marshal(employee)

				if err != nil {
					fmt.Println("Error while Marshling.")
					os.Exit(1)
				}

				w.Header().Set("Content-Type", "application/json")
				w.WriteHeader(http.StatusOK)
				w.Write(byteEmployee)

				return
			*/

			employees = append(employees[:index], employees[index+1:]...)
			employees = append(employees, newEmployee)

			fmt.Printf("\nEmployee details of ID: %d after any updates.\n", newEmployee.ID)
			fmt.Printf("ID: %d\tName: %s\tDesignation: %s\n", newEmployee.ID, newEmployee.Name, newEmployee.Designation)

			byteEmployee, err = json.Marshal(newEmployee)

			if err != nil {
				fmt.Println("Error while Marshling.")
				os.Exit(1)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(byteEmployee)

			return
		}
	}

	fmt.Printf("\nMessage: No employee with ID: %d found.\n", id)
}

func DeleteEmployeeByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	id, err := strconv.ParseInt(params["employee_id"], 0, 0)

	if err != nil {
		fmt.Println("Error while converting from string to int.")
		os.Exit(1)
	}

	for index, employee := range employees {
		if employee.ID == int(id) {
			fmt.Printf("\nEmployee with ID: %d found.\n", id)
			fmt.Printf("Employee details of ID: %d which will be deleted.\n", id)
			fmt.Printf("ID: %d\tName: %s\tDesignation: %s\n", employee.ID, employee.Name, employee.Designation)

			employees = append(employees[:index], employees[index+1:]...)

			byteEmployee, err := json.Marshal(employee)

			if err != nil {
				fmt.Println("Error while Marshling.")
				os.Exit(1)
			}

			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusOK)
			w.Write(byteEmployee)

			return
		}
	}

	fmt.Printf("\nMessage: No employee with ID: %d found.\n", id)
}
