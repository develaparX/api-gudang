package handler

import (
	"net/http"

	"api-gudang/dto"
	"api-gudang/internal/service"

	"github.com/gin-gonic/gin"
)

type BarangMonitoringHandler struct {
	barangService service.BarangService
}

func NewBarangMonitoringHandler(barangService service.BarangService) *BarangMonitoringHandler {
	return &BarangMonitoringHandler{
		barangService: barangService,
	}
}

func (h *BarangMonitoringHandler) GetExpiredBarang(c *gin.Context) {
	expiredBarangs, err := h.barangService.GetExpiredBarang(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully retreived",
		Data:    expiredBarangs,
	})
}
