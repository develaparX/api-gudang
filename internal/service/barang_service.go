package service

import (
	"context"
	"time"

	"api-gudang/dto"
	"api-gudang/internal/models"
	"api-gudang/internal/repository"
)

type BarangService interface {
	Create(ctx context.Context, barang *dto.Barang) error
	Update(ctx context.Context, barang *dto.Barang) error
	Delete(ctx context.Context, barangID string) error
	GetByID(ctx context.Context, barangID string) (*models.Barang, error)
	GetAll(ctx context.Context, limit, offset int, kodeGudang *string, expiredBarang *time.Time) ([]*models.Barang, error)
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

func (s *barangService) Create(ctx context.Context, barang *dto.Barang) error {
	gudang, err := s.gudangRepo.GetByKode(ctx, barang.KodeGudang)
	if err != nil {
		return err
	}
	barang.KodeGudang = gudang.KodeGudang

	return s.barangRepo.Create(ctx, barang)
}

func (s *barangService) Update(ctx context.Context, barang *dto.Barang) error {
	gudang, err := s.gudangRepo.GetByKode(ctx, barang.KodeGudang)
	if err != nil {
		return err
	}
	barang.KodeGudang = gudang.KodeGudang

	return s.barangRepo.Update(ctx, barang)
}

func (s *barangService) Delete(ctx context.Context, barangID string) error {
	return s.barangRepo.Delete(ctx, barangID)
}

func (s *barangService) GetByID(ctx context.Context, barangID string) (*models.Barang, error) {
	barang, err := s.barangRepo.GetByID(ctx, barangID)
	if err != nil {
		return nil, err
	}

	gudang, err := s.gudangRepo.GetByKode(ctx, barang.Gudang[0].KodeGudang)
	if err != nil {
		return nil, err
	}
	barang.Gudang[0].NamaGudang = gudang.NamaGudang

	return barang, nil
}

func (s *barangService) GetAll(ctx context.Context, limit, offset int, kodeGudang *string, expiredBarang *time.Time) ([]*models.Barang, error) {
	barangs, err := s.barangRepo.GetAll(ctx, limit, offset, kodeGudang, expiredBarang)
	if err != nil {
		return nil, err
	}

	for _, barang := range barangs {
		gudang, err := s.gudangRepo.GetByKode(ctx, barang.Gudang[0].KodeGudang)
		if err != nil {
			return nil, err
		}
		barang.Gudang[0].NamaGudang = gudang.NamaGudang
	}
	return barangs, nil
}

func (s *barangService) GetExpiredBarang(ctx context.Context) ([]*models.Barang, error) {
	expiredBarangs, err := s.barangRepo.GetExpiredBarang(ctx)
	if err != nil {
		return nil, err
	}

	for _, barang := range expiredBarangs {
		gudang, err := s.gudangRepo.GetByKode(ctx, barang.Gudang[0].KodeGudang)
		if err != nil {
			return nil, err
		}
		barang.Gudang[0].NamaGudang = gudang.NamaGudang
	}
	return expiredBarangs, nil
}
