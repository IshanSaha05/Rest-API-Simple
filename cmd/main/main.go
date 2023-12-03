package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/IshanSaha05/microservice/pkg/routes"
	"github.com/gorilla/mux"
)

func main() {
	fmt.Println("Welcome to Employee REST-API Code.")
	fmt.Println("----------------------------------")

	router := mux.NewRouter()

	routes.EmployeeRoutes(router)
	//http.Handle("/", router) //would it be require?

	fmt.Println("Starting server at localhost:8080")

	err := http.ListenAndServe(":8080", router)

	if err != nil {
		fmt.Println("Error while starting the server.")
		os.Exit(1)
	}
}
