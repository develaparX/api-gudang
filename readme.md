### Dokumentasi API Gudang-API 

#### Cara Menjalankan Project

1. **Pastikan GoLang v1.22 Terinstal**
   - Sebelum memulai, pastikan bahwa GoLang versi 1.22 sudah terinstal di sistem Anda.

2. **Jalankan Query di `ddl.sql`**
   - Sebelum menjalankan project, jalankan terlebih dahulu query-query yang ada di dalam file `ddl.sql`.

3. **Konfigurasi `.env`**
   - Salin file `.env.example` dan beri nama `.env`.
   - Atur nilai-nilai di dalam file `.env` sesuai dengan pengaturan database dan konfigurasi lainnya.

4. **Instalasi Dependencies**
   - Masuk ke direktori project dan jalankan perintah berikut untuk menginstall semua dependency yang dibutuhkan:
     ```bash
     go mod tidy
     ```

5. **Menjalankan Project**
   - Setelah semua konfigurasi siap, jalankan project dengan perintah:
     ```bash
     go run .
     ```

6. **Referensi File `readme.postgres.md dam readme.algoritmaBerfikir.md`**
   - Untuk penjelasan lebih lanjut mengenai query-query yang ada di file `ddl.sql`, silakan lihat dokumentasi di dalam file `readme.postgres.md`.
   - penjelasan mengenai algoritma berifikir ada di file `readme.algoritmaBerfikir.md`

---

### Dokumentasi API

Berikut adalah penjelasan singkat dan endpoint dari API Gudang-API:

#### 1. **Barang**

- **Create Barang**
  - **Endpoint:** `POST /barang`
  - **URL:** `localhost:3500/barang`
  - **Body:**
    ```json
    {
      "kode_barang": "AABB1",
      "nama_barang": "permen",
      "harga_barang": 10000,
      "jumlah_barang": 10,
      "expired_barang": "2025-01-02T15:04:05Z",
      "kode_gudang": "G001"
    }
    ```
  - Menambahkan barang baru ke dalam sistem.

- **Get All Barang**
  - **Endpoint:** `GET /barang`
  - **URL:** `localhost:3500/barang`
  - **Query Parameters:**
    - `limit` (opsional): Membatasi jumlah hasil yang dikembalikan. Default: 10.
    - `offset` (opsional): Mengatur offset untuk hasil yang dikembalikan. Default: 0.
    - `kode_gudang` (opsional): Filter berdasarkan kode gudang.
    - `expired_barang` (opsional): Filter berdasarkan tanggal expired barang (format: `YYYY-MM-DD`).
  - Mengambil semua data barang yang ada di dalam sistem dengan dukungan filter dan pagination.

- **Get Barang by ID**
  - **Endpoint:** `GET /barang/{id}`
  - **URL:** `localhost:3500/barang/{id}`
  - Mengambil data barang berdasarkan ID.

- **Update Barang**
  - **Endpoint:** `PUT /barang`
  - **URL:** `localhost:3500/barang`
  - **Body:**
    ```json
    {
      "kode_barang": "AABB1",
      "nama_barang": "Topi",
      "harga_barang": 10000,
      "jumlah_barang": 10,
      "expired_barang": "2025-01-02T15:04:05Z",
      "kode_gudang": "G001"
    }
    ```
  - Memperbarui data barang yang ada.

- **Delete Barang**
  - **Endpoint:** `DELETE /barang/{id}`
  - **URL:** `localhost:3500/barang/{id}`
  - Menghapus barang berdasarkan ID.

- **Get Expired Barang**
  - **Endpoint:** `GET /barang/expired`
  - **URL:** `localhost:3500/barang/expired`
  - Mengambil semua barang yang sudah expired.

#### 2. **Gudang**

- **Create Gudang**
  - **Endpoint:** `POST /gudang`
  - **URL:** `localhost:3500/gudang`
  - **Body:**
    ```json
    {
      "kode_gudang": "G005",
      "nama_gudang": "Gudang Banjarnegara"
    }
    ```
  - Menambahkan gudang baru ke dalam sistem.

- **Get All Gudang**
  - **Endpoint:** `GET /gudang`
  - **URL:** `localhost:3500/gudang`
  - Mengambil semua data gudang yang ada di dalam sistem.

- **Get Gudang by Kode Gudang**
  - **Endpoint:** `GET /gudang/{kode_gudang}`
  - **URL:** `localhost:3500/gudang/{kode_gudang}`
  - Mengambil data gudang berdasarkan kode gudang.

- **Update Gudang**
  - **Endpoint:** `PUT /gudang`
  - **URL:** `localhost:3500/gudang`
  - **Body:**
    ```json
    {
      "kode_gudang": "G001",
      "nama_gudang": "Gudang Banjarnegara"
    }
    ```
  - Memperbarui data gudang yang ada.

- **Delete Gudang**
  - **Endpoint:** `DELETE /gudang/{kode_gudang}`
  - **URL:** `localhost:3500/gudang/{kode_gudang}`
  - Menghapus gudang berdasarkan kode gudang.

---

Pastikan untuk mengikuti langkah-langkah tersebut dengan urutan yang benar agar API dapat berjalan dengan baik.