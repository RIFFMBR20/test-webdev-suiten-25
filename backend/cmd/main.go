package main

import (
	"log"
	"net/http"
	"test-webdev-suiten-25/config"
	"test-webdev-suiten-25/database"
	"test-webdev-suiten-25/internal/models/dao"
	"test-webdev-suiten-25/internal/routes"

	"github.com/joho/godotenv"
)

// Package main API.
//
// @title           Test Webdev Suiten API
// @version         1.0
// @description     API documentation
//
// @host            localhost:8080
// @BasePath        /
//
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	cfg := config.LoadConfig()

	if err := database.Connect(cfg); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	log.Println("Starting database migration...")

	err := database.DB.AutoMigrate(
		&dao.Division{},
		&dao.Employee{},
		&dao.Attendance{},
	)

	if err != nil {
		log.Fatalf("Failed to perform auto migration: %v", err)
	}
	log.Println("Database migration completed successfully.")

	port := ":" + cfg.AppPort

	mux := routes.NewRouter(database.DB)

	log.Printf("Server starting on port %s", port)
	log.Println("Swagger UI available at http://localhost:8080/swagger/index.html")
	if err := http.ListenAndServe(port, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}

}
