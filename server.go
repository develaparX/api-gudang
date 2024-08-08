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
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}

	db, err := config.GetDB()
	if err != nil {
		panic(err)
	}

	gudangRepo := repository.NewGudangRepository(db)
	barangRepo := repository.NewBarangRepository(db)

	gudangService := service.NewGudangService(gudangRepo)
	barangService := service.NewBarangService(barangRepo, gudangRepo)

	gudangHandler := handler.NewGudangHandler(gudangService)
	barangHandler := handler.NewBarangHandler(barangService, gudangService)

	r := gin.Default()

	routes.RegisterGudangRoutes(r, gudangHandler)
	routes.RegisterBarangRoutes(r, barangHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "4000"
	}
	r.Run(":" + port)
}
