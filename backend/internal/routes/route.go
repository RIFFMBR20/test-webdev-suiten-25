package routes

import (
	"log"
	"os"
	"test-webdev-suiten-25/internal/controller"
	"test-webdev-suiten-25/internal/repository"
	"test-webdev-suiten-25/internal/service"

	_ "test-webdev-suiten-25/docs"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger/v2"
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
	employeeRepo := repository.ProvideEmployeeRepository(db)
	employeeService := service.ProvideEmployeeService(employeeRepo, divisionRepo, logger)
	employeeController := controller.ProvideEmployeeController(employeeService, logger)

	RegisterEmployeeRoutes(router, employeeController)

	//swagger
	router.PathPrefix("/swagger/").Handler(httpSwagger.WrapHandler)

	return router
}
