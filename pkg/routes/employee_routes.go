package routes

import (
	"github.com/IshanSaha05/microservice/pkg/controller"
	"github.com/gorilla/mux"
)

var EmployeeRoutes = func(router *mux.Router) {
	router.HandleFunc("/api/v1/employees", controller.CreateEmployee).Methods("POST")
	router.HandleFunc("/api/v1/employees", controller.GetAllEmployees).Methods("GET")
	router.HandleFunc("/api/v1/employees/{employee_id}", controller.GetEmployeeByID).Methods("GET")
	router.HandleFunc("/api/v1/employees/{employee_id}", controller.UpdateEmployeeByID).Methods("PUT")
	router.HandleFunc("/api/v1/employees/{employee_id}", controller.DeleteEmployeeByID).Methods("DELETE")
}
