package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"test-webdev-suiten-25/config"
	"test-webdev-suiten-25/database"
	"test-webdev-suiten-25/docs"
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

	envPort := strings.TrimSpace(os.Getenv("PORT"))
	cfgPort := strings.TrimSpace(cfg.AppPort)

	port := envPort
	portSource := "PORT(env)"
	if port == "" {
		port = cfgPort
		portSource = "cfg.AppPort"
	}
	if port == "" {
		port = "8080"
		portSource = "default(8080)"
	}
	addr := ":" + port

	envSwaggerHost := strings.TrimSpace(os.Getenv("SWAGGER_HOST"))
	envRailwayDomain := strings.TrimSpace(os.Getenv("RAILWAY_PUBLIC_DOMAIN"))

	host := envSwaggerHost
	hostSource := "SWAGGER_HOST(env)"
	if host == "" {
		host = envRailwayDomain
		hostSource = "RAILWAY_PUBLIC_DOMAIN(env)"
	}

	if host != "" {
		docs.SwaggerInfo.Host = host
		docs.SwaggerInfo.Schemes = []string{"https"}
	} else {
		docs.SwaggerInfo.Schemes = []string{"http"}
	}

	mux := routes.NewRouter(database.DB)

	log.Printf("Server listen addr=%s (source=%s, env.PORT=%q, cfg.AppPort=%q)", addr, portSource, envPort, cfgPort)
	log.Printf("Swagger config host=%q (source=%s, SWAGGER_HOST=%q, RAILWAY_PUBLIC_DOMAIN=%q) schemes=%v basePath=%q",
		docs.SwaggerInfo.Host, hostSource, envSwaggerHost, envRailwayDomain, docs.SwaggerInfo.Schemes, docs.SwaggerInfo.BasePath,
	)
	log.Printf("Swagger UI should be available at: %s://%s/swagger/index.html",
		firstSchemeOrDefault(docs.SwaggerInfo.Schemes, "http"),
		func() string {
			if docs.SwaggerInfo.Host != "" {
				return docs.SwaggerInfo.Host
			}
			return "localhost" + addr
		}(),
	)

	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatalf("failed to start server: %v", err)
	}
}

func firstSchemeOrDefault(schemes []string, def string) string {
	if len(schemes) == 0 {
		return def
	}
	if schemes[0] == "" {
		return def
	}
	return schemes[0]
}
