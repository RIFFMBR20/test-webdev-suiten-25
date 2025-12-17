package routes

import (
	"test-webdev-suiten-25/internal/controller"

	"github.com/gorilla/mux"
)

func RegisterEmployeeRoutes(
	router *mux.Router,
	controller controller.EmployeeController,
) {
	router.HandleFunc("/api/employee", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/employee", controller.Create).Methods("POST")

	router.HandleFunc("/api/employee/{id}", controller.GetByID).Methods("GET")
	router.HandleFunc("/api/employee/{id}", controller.Update).Methods("PUT")
	router.HandleFunc("/api/employee/{id}", controller.Delete).Methods("DELETE")
}
