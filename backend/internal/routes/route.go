package routes

import (
	"log"
	"os"
	"test-webdev-suiten-25/internal/controller"
	"test-webdev-suiten-25/internal/repository"
	"test-webdev-suiten-25/internal/service"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func NewRouter(db *gorm.DB) *mux.Router {
	router := mux.NewRouter()

	logger := log.New(os.Stdout, "[APP] ", log.LstdFlags)

	// division
	divisionRepo := repository.ProvideDivisionRepository(db)
	divisionService := service.ProvideDivisionService(divisionRepo, logger)
	divisionController := controller.ProvideDivisionController(divisionService, logger)

	RegisterDivisionRoutes(router, divisionController)

	// employee
	employeeRepo := repository.ProvideDivisionRepository(db)
	employeeService := service.ProvideDivisionService(employeeRepo, logger)
	employeeController := controller.ProvideDivisionController(employeeService, logger)

	RegisterEmployeeRoutes(router, employeeController)

	return router
}
