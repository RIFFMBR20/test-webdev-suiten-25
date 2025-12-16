package routes

import (
	"test-webdev-suiten-25/internal/controller"

	"github.com/gorilla/mux"
)

func RegisterEmployeeRoutes(
	router *mux.Router,
	controller controller.EmployeeController,
) {
	router.HandleFunc("/api/divisions", controller.GetAll).Methods("GET")
	router.HandleFunc("/api/divisions", controller.Create).Methods("POST")

	router.HandleFunc("/api/divisions/{id}", controller.GetByID).Methods("GET")
	router.HandleFunc("/api/divisions/{id}", controller.Update).Methods("PUT")
	router.HandleFunc("/api/divisions/{id}", controller.Delete).Methods("DELETE")
}
