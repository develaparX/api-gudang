package repository

import (
	"api-gudang/dto"
	"api-gudang/internal/models"
	"context"
	"database/sql"
	"fmt"
	"time"
)

type BarangRepository interface {
	Create(ctx context.Context, barang *dto.Barang) error
	Update(ctx context.Context, barang *dto.Barang) error
	Delete(ctx context.Context, barangID string) error
	GetByID(ctx context.Context, barangID string) (*models.Barang, error)
	GetAll(ctx context.Context, limit, offset int, kodeGudang *string, expiredBarang *time.Time) ([]*models.Barang, error)
	GetExpiredBarang(ctx context.Context) ([]*models.Barang, error)
}

type barangRepository struct {
	db *sql.DB
}

func NewBarangRepository(db *sql.DB) BarangRepository {
	return &barangRepository{db: db}
}

func (r *barangRepository) Create(ctx context.Context, barang *dto.Barang) error {
	query := `INSERT INTO barang (kode_barang, nama_barang, harga_barang, jumlah_barang, expired_barang, kode_gudang) 
              VALUES ($1, $2, $3, $4, $5, $6)`

	// Asumsikan Gudang sudah ada di model Barang
	_, err := r.db.ExecContext(ctx, query, barang.KodeBarang, barang.NamaBarang, barang.HargaBarang,
		barang.JumlahBarang, barang.ExpiredBarang, barang.KodeGudang)
	if err != nil {
		return err
	}
	return nil
}

func (r *barangRepository) Update(ctx context.Context, barang *dto.Barang) error {
	query := `UPDATE barang SET kode_barang = $1, nama_barang = $2, harga_barang = $3, jumlah_barang = $4, expired_barang = $5, 
              kode_gudang = $6 WHERE barang_id = $7`

	// Asumsikan Gudang sudah ada di model Barang
	_, err := r.db.ExecContext(ctx, query, barang.KodeBarang, barang.NamaBarang, barang.HargaBarang, barang.JumlahBarang,
		barang.ExpiredBarang, barang.KodeGudang, barang.BarangID)
	if err != nil {
		return err
	}
	return nil
}

func (r *barangRepository) Delete(ctx context.Context, barangID string) error {
	query := `DELETE FROM barang WHERE barang_id = $1`
	_, err := r.db.ExecContext(ctx, query, barangID)
	if err != nil {
		return err
	}
	return nil
}

func (r *barangRepository) GetByID(ctx context.Context, barangID string) (*models.Barang, error) {
	query := `SELECT barang_id, kode_barang, nama_barang, harga_barang, jumlah_barang, expired_barang, kode_gudang
              FROM barang WHERE barang_id = $1`

	row := r.db.QueryRowContext(ctx, query, barangID)

	var barang models.Barang
	var gudang models.Gudang
	err := row.Scan(&barang.BarangID, &barang.KodeBarang, &barang.NamaBarang, &barang.HargaBarang, &barang.JumlahBarang,
		&barang.ExpiredBarang, &gudang.KodeGudang)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("barang with ID %s not found", barangID)
		}
		return nil, err
	}

	barang.Gudang = append(barang.Gudang, gudang)
	return &barang, nil
}

func (r *barangRepository) GetAll(ctx context.Context, limit, offset int, kodeGudang *string, expiredBarang *time.Time) ([]*models.Barang, error) {
	query := `SELECT barang_id, kode_barang, nama_barang, harga_barang, jumlah_barang, expired_barang, kode_gudang, nama_gudang 
              FROM get_barang_list($1, $2, $3, $4)`

	rows, err := r.db.QueryContext(ctx, query, limit, offset, kodeGudang, expiredBarang)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var barangs []*models.Barang

	for rows.Next() {
		var barang models.Barang
		var gudang models.Gudang
		err := rows.Scan(&barang.BarangID, &barang.KodeBarang, &barang.NamaBarang, &barang.HargaBarang, &barang.JumlahBarang,
			&barang.ExpiredBarang, &gudang.KodeGudang, &gudang.NamaGudang)
		if err != nil {
			return nil, err
		}
		barang.Gudang = append(barang.Gudang, gudang)
		barangs = append(barangs, &barang)
	}
	return barangs, nil
}

func (r *barangRepository) GetExpiredBarang(ctx context.Context) ([]*models.Barang, error) {
	query := `SELECT barang_id, kode_barang, nama_barang, harga_barang, jumlah_barang, expired_barang, kode_gudang, nama_gudang 
              FROM get_barang_list(100, 0, NULL, CURRENT_DATE)`

	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var barangs []*models.Barang

	for rows.Next() {
		var barang models.Barang
		var gudang models.Gudang
		err := rows.Scan(&barang.BarangID, &barang.KodeBarang, &barang.NamaBarang, &barang.HargaBarang, &barang.JumlahBarang,
			&barang.ExpiredBarang, &gudang.KodeGudang, &gudang.NamaGudang)
		if err != nil {
			return nil, err
		}
		barang.Gudang = append(barang.Gudang, gudang)
		barangs = append(barangs, &barang)
	}
	return barangs, nil
}
