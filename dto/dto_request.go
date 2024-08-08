package dto

import "time"

type Barang struct {
	BarangID      string    `json:"barang_id"`
	KodeBarang    string    `json:"kode_barang"`
	NamaBarang    string    `json:"nama_barang"`
	HargaBarang   int       `json:"harga_barang"`
	JumlahBarang  int       `json:"jumlah_barang"`
	ExpiredBarang time.Time `json:"expired_barang"`
	KodeGudang    string    `json:"kode_gudang"`
}
