package handler

import (
	"net/http"

	"api-gudang/dto"
	"api-gudang/internal/models"
	"api-gudang/internal/service"

	"github.com/gin-gonic/gin"
)

type GudangHandler struct {
	gudangService service.GudangService
}

func NewGudangHandler(gudangService service.GudangService) *GudangHandler {
	return &GudangHandler{
		gudangService: gudangService,
	}
}

func (h *GudangHandler) CreateGudang(c *gin.Context) {
	var gudang models.Gudang
	if err := c.ShouldBindJSON(&gudang); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err := h.gudangService.Create(c.Request.Context(), &gudang)
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
		Data:    gudang,
	})
}

func (h *GudangHandler) UpdateGudang(c *gin.Context) {
	var gudang models.Gudang
	if err := c.ShouldBindJSON(&gudang); err != nil {
		c.JSON(http.StatusBadRequest, dto.Response{
			Success: false,
			Message: err.Error(),
		})
		return
	}

	err := h.gudangService.Update(c.Request.Context(), &gudang)
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
		Data:    gudang,
	})
}

func (h *GudangHandler) DeleteGudang(c *gin.Context) {
	kodeGudang := c.Param("kodeGudang")

	err := h.gudangService.Delete(c.Request.Context(), kodeGudang)
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

func (h *GudangHandler) GetGudangByKode(c *gin.Context) {
	kodeGudang := c.Param("kodeGudang")

	gudang, err := h.gudangService.GetByKode(c.Request.Context(), kodeGudang)
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
		Data:    gudang,
	})
}

func (h *GudangHandler) GetAllGudang(c *gin.Context) {
	gudangs, err := h.gudangService.GetAll(c.Request.Context())
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
		Data:    gudangs,
	})
}
