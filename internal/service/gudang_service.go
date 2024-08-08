package service

import (
	"context"

	"api-gudang/internal/models"
	"api-gudang/internal/repository"
)

type GudangService interface {
	Create(ctx context.Context, gudang *models.Gudang) error
	Update(ctx context.Context, gudang *models.Gudang) error
	Delete(ctx context.Context, kodeGudang string) error
	GetByKode(ctx context.Context, kodeGudang string) (*models.Gudang, error)
	GetAll(ctx context.Context) ([]*models.Gudang, error)
}

type gudangService struct {
	gudangRepo repository.GudangRepository
}

func NewGudangService(gudangRepo repository.GudangRepository) GudangService {
	return &gudangService{
		gudangRepo: gudangRepo,
	}
}

func (s *gudangService) Create(ctx context.Context, gudang *models.Gudang) error {
	return s.gudangRepo.Create(ctx, gudang)
}

func (s *gudangService) Update(ctx context.Context, gudang *models.Gudang) error {
	return s.gudangRepo.Update(ctx, gudang)
}

func (s *gudangService) Delete(ctx context.Context, kodeGudang string) error {
	return s.gudangRepo.Delete(ctx, kodeGudang)
}

func (s *gudangService) GetByKode(ctx context.Context, kodeGudang string) (*models.Gudang, error) {
	return s.gudangRepo.GetByKode(ctx, kodeGudang)
}

func (s *gudangService) GetAll(ctx context.Context) ([]*models.Gudang, error) {
	return s.gudangRepo.GetAll(ctx)
}
