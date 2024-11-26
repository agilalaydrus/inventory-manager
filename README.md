
# Sistem Manajemen Inventaris

Proyek ini adalah sistem backend untuk mengelola aplikasi manajemen inventaris. Sistem ini dibangun menggunakan Golang dan framework Gin, serta berinteraksi dengan database relasional untuk mengelola produk, inventaris, dan pesanan.

## Tabel Isi

- [Deskripsi Proyek](#deskripsi-proyek)
- [Desain Database](#desain-database)
- [Query SQL](#query-sql)
- [Pengembangan RESTful API](#pengembangan-restful-api)
- [Instalasi](#instalasi)
- [Penggunaan](#penggunaan)
- [Contributing](#contributing)
- [Lisensi](#lisensi)

## Deskripsi Proyek

Proyek ini bertujuan untuk memberikan pengalaman langsung dalam pengembangan backend, integrasi API, dan penanganan penyimpanan file. Dengan menyelesaikan tugas ini, siswa akan menerapkan konsep pemrograman inti dalam skenario dunia nyata.

## Desain Database

Database ini memiliki tiga tabel utama:

1. **Produk**
    - ID (int, primary key)
    - Nama (string)
    - Deskripsi (string)
    - Harga (float)
    - Kategori (string)

2. **Inventaris**
    - ID Produk (int, foreign key)
    - Jumlah (int)
    - Lokasi (string)

3. **Pesanan**
    - ID Pesanan (int, primary key)
    - ID Produk (int, foreign key)
    - Jumlah (int)
    - Tanggal Pesanan (datetime)

## Query SQL

Berikut adalah contoh skrip SQL untuk membuat tabel dan memasukkan data sampel:

```sql
-- Membuat tabel Produk
CREATE TABLE Produk (
    id INT AUTO_INCREMENT PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    deskripsi TEXT,
    harga DECIMAL(10, 2) NOT NULL,
    kategori VARCHAR(100)
);

-- Membuat tabel Inventaris
CREATE TABLE Inventaris (
    id_produk INT,
    jumlah INT NOT NULL,
    lokasi VARCHAR(255),
    FOREIGN KEY (id_produk) REFERENCES Produk(id)
);

-- Membuat tabel Pesanan
CREATE TABLE Pesanan (
    id_pesanan INT AUTO_INCREMENT PRIMARY KEY,
    id_produk INT,
    jumlah INT NOT NULL,
    tanggal_pesanan DATETIME DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (id_produk) REFERENCES Produk(id)
);

-- Memasukkan data sampel
INSERT INTO Produk (nama, deskripsi, harga, kategori) VALUES 
('CocaCola', 'Minuman ringan berkarbonasi', 10.00, 'Minuman'),
('Pepsi', 'Minuman ringan berkarbonasi', 10.00, 'Minuman');

INSERT INTO Inventaris (id_produk, jumlah, lokasi) VALUES 
(1, 100, 'Gudang A'),
(2, 50, 'Gudang B');

INSERT INTO Pesanan (id_produk, jumlah) VALUES 
(1, 2),
(2, 3);
=======
# inventory-manager
>>>>>>> origin/main
