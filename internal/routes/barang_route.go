package routes

import (
	"api-gudang/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBarangRoutes(r *gin.Engine, barangHandler *handler.BarangHandler) {
	barangGroup := r.Group("/barang")
	{
		barangGroup.POST("/", barangHandler.CreateBarang)
		barangGroup.PUT("/", barangHandler.UpdateBarang)
		barangGroup.DELETE("/:barangID", barangHandler.DeleteBarang)
		barangGroup.GET("/:barangID", barangHandler.GetBarangByID)
		barangGroup.GET("/", barangHandler.GetAllBarang)
	}
}
