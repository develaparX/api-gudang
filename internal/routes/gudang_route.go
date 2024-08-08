package routes

import (
	"api-gudang/internal/handler"

	"github.com/gin-gonic/gin"
)

func RegisterGudangRoutes(r *gin.Engine, gudangHandler *handler.GudangHandler) {
	gudangGroup := r.Group("/gudang")
	{
		gudangGroup.POST("/", gudangHandler.CreateGudang)
		gudangGroup.PUT("/", gudangHandler.UpdateGudang)
		gudangGroup.DELETE("/:kodeGudang", gudangHandler.DeleteGudang)
		gudangGroup.GET("/:kodeGudang", gudangHandler.GetGudangByKode)
		gudangGroup.GET("/", gudangHandler.GetAllGudang)
	}
}
