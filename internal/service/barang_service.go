package service

import (
	"context"

	"api-gudang/internal/models"
	"api-gudang/internal/repository"
)

type BarangService interface {
	Create(ctx context.Context, barang *models.Barang) error
	Update(ctx context.Context, barang *models.Barang) error
	Delete(ctx context.Context, barangID string) error
	GetByID(ctx context.Context, barangID string) (*models.Barang, error)
	GetAll(ctx context.Context, filters *models.BarangFilters) ([]*models.Barang, error)
	GetExpiredBarang(ctx context.Context) ([]*models.Barang, error)
}

type barangService struct {
	barangRepo repository.BarangRepository
	gudangRepo repository.GudangRepository
}

func NewBarangService(barangRepo repository.BarangRepository, gudangRepo repository.GudangRepository) BarangService {
	return &barangService{
		barangRepo: barangRepo,
		gudangRepo: gudangRepo,
	}
}

func (s *barangService) Create(ctx context.Context, barang *models.Barang) error {
	// Check if the gudang exists
	_, err := s.gudangRepo.GetByKode(ctx, barang.KodeGudang)
	if err != nil {
		return err
	}

	return s.barangRepo.Create(ctx, barang)
}

func (s *barangService) Update(ctx context.Context, barang *models.Barang) error {
	// Check if the gudang exists
	_, err := s.gudangRepo.GetByKode(ctx, barang.KodeGudang)
	if err != nil {
		return err
	}

	return s.barangRepo.Update(ctx, barang)
}

func (s *barangService) Delete(ctx context.Context, barangID string) error {
	return s.barangRepo.Delete(ctx, barangID)
}

func (s *barangService) GetByID(ctx context.Context, barangID string) (*models.Barang, error) {
	return s.barangRepo.GetByID(ctx, barangID)
}

func (s *barangService) GetAll(ctx context.Context, filters *models.BarangFilters) ([]*models.Barang, error) {
	return s.barangRepo.GetAll(ctx, filters)
}

func (s *barangService) GetExpiredBarang(ctx context.Context) ([]*models.Barang, error) {
	expiredBarangs, err := s.barangRepo.GetExpiredBarang(ctx)
	if err != nil {
		return nil, err
	}

	// Fetch the gudang names for the expired barangs
	for _, barang := range expiredBarangs {
		gudang, err := s.gudangRepo.GetByKode(ctx, barang.KodeGudang)
		if err != nil {
			return nil, err
		}
		barang.NamaGudang = gudang.NamaGudang
	}

	return expiredBarangs, nil
}
