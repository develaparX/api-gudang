package handler

import (
	"net/http"

	"api-gudang/dto"
	"api-gudang/internal/models"
	"api-gudang/internal/service"

	"github.com/gin-gonic/gin"
)

type BarangHandler struct {
	barangService service.BarangService
	gudangService service.GudangService
}

func NewBarangHandler(barangService service.BarangService, gudangService service.GudangService) *BarangHandler {
	return &BarangHandler{
		barangService: barangService,
		gudangService: gudangService,
	}
}

func (h *BarangHandler) CreateBarang(c *gin.Context) {
	var barang models.Barang
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err := h.barangService.Create(c.Request.Context(), &barang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusCreated, dto.Response{
		Success: true,
		Message: "success created",
		Data:    barang,
	})
}

func (h *BarangHandler) UpdateBarang(c *gin.Context) {
	var barang models.Barang
	if err := c.ShouldBindJSON(&barang); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err := h.barangService.Update(c.Request.Context(), &barang)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "success updated",
		Data:    barang,
	})
}

func (h *BarangHandler) DeleteBarang(c *gin.Context) {
	barangID := c.Param("barangID")

	err := h.barangService.Delete(c.Request.Context(), barangID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "success delete",
	})
}

func (h *BarangHandler) GetBarangByID(c *gin.Context) {
	barangID := c.Param("barangID")

	barang, err := h.barangService.GetByID(c.Request.Context(), barangID)
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
		Data:    barang,
	})
}

func (h *BarangHandler) GetAllBarang(c *gin.Context) {
	kodeGudang := c.Query("kode_gudang")
	expiredBarang := c.Query("expired_barang")

	filters := &models.BarangFilters{
		KodeGudang:    kodeGudang,
		ExpiredBarang: expiredBarang,
	}

	barangs, err := h.barangService.GetAll(c.Request.Context(), filters)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, dto.Response{
		Success: true,
		Message: "successfully retrieved",
		Data:    barangs,
	})
}
