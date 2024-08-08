
## **Dokumentasi Skema Database Gudang dan Barang**

### **1. Tabel `gudang`**

#### **Deskripsi:**
Tabel `gudang` menyimpan informasi mengenai gudang. Setiap gudang memiliki kode unik (`kode_gudang`) yang digunakan sebagai primary key.

#### **Skema:**
```sql
CREATE TABLE gudang (
    kode_gudang VARCHAR(5) UNIQUE PRIMARY KEY,  -- Primary key sebagai kode unik dengan panjang maksimal 5 karakter
    nama_gudang VARCHAR(255) NOT NULL  -- Nama gudang, tidak boleh null
);
```

#### **Kegunaan:**
- **`kode_gudang`**: Digunakan untuk mengidentifikasi gudang secara unik dalam sistem.
- **`nama_gudang`**: Nama dari gudang yang akan digunakan untuk keperluan identifikasi oleh pengguna.

---

### **2. Tabel `barang`**

#### **Deskripsi:**
Tabel `barang` menyimpan informasi mengenai barang yang disimpan di gudang. Setiap barang memiliki ID unik (`barang_id`) dan dihubungkan dengan gudang melalui foreign key `kode_gudang`.

#### **Skema:**
```sql
CREATE TABLE barang (
    barang_id UUID PRIMARY KEY DEFAULT gen_random_uuid(),  -- Primary key sebagai UUID unik
    kode_barang VARCHAR(50) UNIQUE NOT NULL,  -- Kode barang dengan panjang maksimal 50 karakter, harus unik
    nama_barang VARCHAR(255) NOT NULL,  -- Nama barang, tidak boleh null
    harga_barang INT NOT NULL,  
    jumlah_barang INT NOT NULL,  -- Jumlah barang, tidak boleh null
    expired_barang DATE NOT NULL,  -- Tanggal kadaluarsa barang, tidak boleh null
    kode_gudang VARCHAR(5) REFERENCES gudang(kode_gudang)  -- Foreign key yang merujuk ke tabel gudang
);
```

#### **Kegunaan:**
- **`barang_id`**: Identifikasi unik untuk setiap barang.
- **`kode_gudang`**: Menyimpan referensi ke gudang tempat barang disimpan, digunakan untuk menghubungkan barang dengan gudang tertentu.

---

### **3. Indeks**

#### **Deskripsi:**
Indeks digunakan untuk meningkatkan kinerja query, terutama ketika mencari data berdasarkan kolom yang sering digunakan dalam pencarian.

#### **Skema:**
```sql
-- Indeks untuk kolom kode_barang
CREATE INDEX idx_barang_kode_barang ON barang (kode_barang);

-- Indeks untuk kolom expired_barang
CREATE INDEX idx_barang_expired_barang ON barang (expired_barang);

-- Indeks untuk kolom kode_gudang
CREATE INDEX idx_barang_kode_gudang ON barang (kode_gudang);
```

#### **Kegunaan:**
- **`idx_barang_kode_barang`**: Meningkatkan kecepatan pencarian barang berdasarkan kode barang.
- **`idx_barang_expired_barang`**: Mempercepat pencarian barang yang kadaluarsa.
- **`idx_barang_kode_gudang`**: Mempercepat pencarian barang berdasarkan gudang.

---

### **4. Fungsi Trigger `check_expired_barang`**

#### **Deskripsi 1:**
Fungsi trigger ini memeriksa apakah barang yang akan dimasukkan atau diperbarui sudah kadaluarsa. Jika barang sudah kadaluarsa, maka fungsi ini akan memberikan peringatan.

#### **Skema:**
```sql
CREATE OR REPLACE FUNCTION check_expired_barang()
RETURNS TRIGGER AS $$
DECLARE
    nama_gudang VARCHAR(255);
BEGIN
    -- Ambil nama_gudang berdasarkan kode_gudang
    SELECT g.nama_gudang INTO nama_gudang 
    FROM gudang g 
    WHERE g.kode_gudang = NEW.kode_gudang;

    IF NEW.expired_barang < CURRENT_DATE THEN
        RAISE NOTICE 'Barang % dengan kode % di gudang % sudah kadaluarsa.', NEW.nama_barang, NEW.kode_barang, nama_gudang;
    END IF;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger untuk memanggil fungsi check_expired_barang
CREATE TRIGGER trg_check_expired
BEFORE INSERT OR UPDATE ON barang
FOR EACH ROW
EXECUTE FUNCTION check_expired_barang();
```
#### **Deskripsi 2:**
Fungsi trigger ini memeriksa apakah barang yang akan dimasukkan atau diperbarui sudah kadaluarsa. Jika barang sudah kadaluarsa, maka fungsi ini akan memberikan peringatan dan data tidak dapat dimasukkan.

#### **Skema:**
```sql
CREATE OR REPLACE FUNCTION check_expired_barang()
RETURNS TRIGGER AS $$
DECLARE
    nama_gudang VARCHAR(255);
BEGIN
    -- Cek apakah barang sudah kadaluarsa
    IF NEW.expired_barang < CURRENT_DATE THEN
        RAISE EXCEPTION 'Item % dengan kode % sudah kadaluarsa, Tidak Dapat Di Input', NEW.nama_barang, NEW.kode_barang;
    END IF;

    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Trigger untuk memanggil fungsi check_expired_barang
CREATE TRIGGER trg_check_expired
BEFORE INSERT OR UPDATE ON barang
FOR EACH ROW
EXECUTE FUNCTION check_expired_barang();
```

#### **Kegunaan:**
- Mencegah penyimpanan atau pembaruan data barang yang sudah kadaluarsa dengan memberikan peringatan kepada pengguna.

---

### **5. Fungsi Trigger `check_barang_before_delete`**

#### **Deskripsi:**
Fungsi trigger ini memeriksa apakah ada barang yang terkait dengan gudang sebelum gudang dihapus. Jika masih ada barang yang terkait, maka penghapusan gudang akan dibatalkan.

#### **Skema:**
```sql
CREATE OR REPLACE FUNCTION check_barang_before_delete()
RETURNS TRIGGER AS $$
BEGIN
    -- Cek apakah masih ada barang yang terkait dengan gudang
    IF EXISTS (SELECT 1 FROM barang WHERE kode_gudang = OLD.kode_gudang) THEN
        RAISE EXCEPTION 'Gagal menghapus %, gudang tersebut masih berisi barang', OLD.kode_gudang;
    END IF;
    RETURN OLD;
END;
$$ LANGUAGE plpgsql;

-- Trigger untuk memanggil fungsi check_barang_before_delete
CREATE TRIGGER trigger_before_delete_gudang
BEFORE DELETE ON gudang
FOR EACH ROW
EXECUTE FUNCTION check_barang_before_delete();
```

#### **Kegunaan:**
- Melindungi integritas data dengan memastikan bahwa gudang tidak dapat dihapus jika masih ada barang yang terhubung ke gudang tersebut.

---

### **6. Stored Procedure `get_barang_list`**

#### **Deskripsi:**
Stored procedure ini digunakan untuk mengambil daftar barang dengan parameter paging (batas dan offset) dan filter berdasarkan gudang dan tanggal kadaluarsa.

#### **Skema:**
```sql
CREATE OR REPLACE FUNCTION get_barang_list(
    p_limit INT,
    p_offset INT,
    p_kode_gudang VARCHAR(5) DEFAULT NULL,
    p_expired_barang DATE DEFAULT NULL
)
RETURNS TABLE(
    barang_id UUID,
    kode_barang VARCHAR,
    nama_barang VARCHAR,
    harga_barang INT,
    jumlah_barang INT,
    expired_barang DATE,
    kode_gudang VARCHAR,
    nama_gudang VARCHAR
) AS $$
BEGIN
    RETURN QUERY
    SELECT b.barang_id, b.kode_barang, b.nama_barang, b.harga_barang, b.jumlah_barang, b.expired_barang, g.kode_gudang, g.nama_gudang
    FROM barang b
    JOIN gudang g ON b.kode_gudang = g.kode_gudang
    WHERE (p_kode_gudang IS NULL OR g.kode_gudang = p_kode_gudang)
    AND (p_expired_barang IS NULL OR b.expired_barang <= p_expired_barang)
    ORDER BY b.expired_barang
    LIMIT p_limit OFFSET p_offset;
END;
$$ LANGUAGE plpgsql;
```

#### **Kegunaan:**
- Mengambil data barang berdasarkan kriteria yang ditentukan, seperti gudang tertentu dan batas waktu kadaluarsa, dengan kemampuan paging untuk mengatur jumlah hasil yang dikembalikan.

---