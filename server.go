package main

import (
	"os"

	"api-gudang/config"
	"api-gudang/internal/handler"
	"api-gudang/internal/repository"
	"api-gudang/internal/routes"
	"api-gudang/internal/service"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	// Get database connection
	db, err := config.GetDB()
	if err != nil {
		panic(err)
	}

	// Initialize the repositories
	gudangRepo := repository.NewGudangRepository(db)
	barangRepo := repository.NewBarangRepository(db)

	// Initialize the services
	gudangService := service.NewGudangService(gudangRepo)
	barangService := service.NewBarangService(barangRepo, gudangRepo)

	// Initialize the handlers
	gudangHandler := handler.NewGudangHandler(gudangService)
	barangHandler := handler.NewBarangHandler(barangService, gudangService)
	barangMonitoringHandler := handler.NewBarangMonitoringHandler(barangService)

	// Set up the router
	r := gin.Default()

	// Register the routes
	routes.RegisterGudangRoutes(r, gudangHandler)
	routes.RegisterBarangRoutes(r, barangHandler)
	routes.RegisterBarangMonitoringRoutes(r, barangMonitoringHandler)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	r.Run(":" + port)
}
