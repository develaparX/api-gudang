package repository

import (
	"context"
	"database/sql"
	"fmt"

	"api-gudang/internal/models"
)

type GudangRepository interface {
	Create(ctx context.Context, gudang *models.Gudang) error
	Update(ctx context.Context, gudang *models.Gudang) error
	Delete(ctx context.Context, kodeGudang string) error
	GetByKode(ctx context.Context, kodeGudang string) (*models.Gudang, error)
	GetAll(ctx context.Context) ([]*models.Gudang, error)
}

type gudangRepository struct {
	db *sql.DB
}

func NewGudangRepository(db *sql.DB) GudangRepository {
	return &gudangRepository{db: db}
}

func (r *gudangRepository) Create(ctx context.Context, gudang *models.Gudang) error {
	query := `INSERT INTO gudang (kode_gudang, nama_gudang) VALUES ($1, $2)`
	_, err := r.db.ExecContext(ctx, query, gudang.KodeGudang, gudang.NamaGudang)
	if err != nil {
		return err
	}
	return nil
}

func (r *gudangRepository) Update(ctx context.Context, gudang *models.Gudang) error {
	query := `UPDATE gudang SET nama_gudang = $1 WHERE kode_gudang = $2`
	_, err := r.db.ExecContext(ctx, query, gudang.NamaGudang, gudang.KodeGudang)
	if err != nil {
		return err
	}
	return nil
}

func (r *gudangRepository) Delete(ctx context.Context, kodeGudang string) error {
	query := `DELETE FROM gudang WHERE kode_gudang = $1`
	_, err := r.db.ExecContext(ctx, query, kodeGudang)
	if err != nil {
		return err
	}
	return nil
}

func (r *gudangRepository) GetByKode(ctx context.Context, kodeGudang string) (*models.Gudang, error) {
	query := `SELECT  kode_gudang, nama_gudang FROM gudang WHERE kode_gudang = $1`
	row := r.db.QueryRowContext(ctx, query, kodeGudang)
	gudang := &models.Gudang{}
	err := row.Scan(&gudang.KodeGudang, &gudang.NamaGudang)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("gudang with code %s not found", kodeGudang)
		}
		return nil, err
	}
	return gudang, nil
}

func (r *gudangRepository) GetAll(ctx context.Context) ([]*models.Gudang, error) {
	query := `SELECT kode_gudang, nama_gudang FROM gudang`
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gudangs := make([]*models.Gudang, 0)
	for rows.Next() {
		gudang := &models.Gudang{}
		err := rows.Scan(&gudang.KodeGudang, &gudang.NamaGudang)
		if err != nil {
			return nil, err
		}
		gudangs = append(gudangs, gudang)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return gudangs, nil
}
