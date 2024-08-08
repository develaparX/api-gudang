package routes

import (
	"api-gudang/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterBarangMonitoringRoutes(r *gin.Engine, barangMonitoringHandler *handler.BarangMonitoringHandler) {
	barangMonitoringGroup := r.Group("/barang/monitoring")
	{
		barangMonitoringGroup.GET("/expired", barangMonitoringHandler.GetExpiredBarang)
	}
}
