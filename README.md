# minimal-viable-inventory

## Blueprint Fase 1: Minimal Viable Inventory (MVI)

### 1. Business Requirements (Kebutuhan Bisnis)
Sebagai manager, saya ingin sistem ini bisa melakukan tiga hal dasar dengan integritas data yang tinggi:
* **Pencatatan Barang:** Setiap barang punya SKU unik, nama, harga, dan stok.
* **Validasi Stok:** Barang tidak boleh terjual jika stoknya kurang dari permintaan.
* **Perubahan Harga:** Harga barang bisa berubah kapan saja, tapi tidak boleh nol atau negatif.

### 2. Technical Architecture (Aturan Main)
Kamu wajib menggunakan **Clean Architecture** dan **DDD** sejak baris kode pertama:
* **Layer Domain:** Buat entitas `Product`. Di sini, logika bisnis seperti `ReduceStock` atau `UpdatePrice` harus berbentuk *Method* milik `struct`, bukan fungsi lepas.
* **Layer Repository (Interface):** Jangan buat koneksi database. Buatlah **Interface** yang mendefinisikan cara mencari barang berdasarkan SKU dan cara menyimpan kembali perubahan barang.
* **Layer Usecase:** Buat sebuah service bernama `InventoryManager`. Service ini yang akan mengoordinasi: "Ambil barang dari Repo -> Validasi/Ubah data di Domain -> Simpan balik ke Repo".

### 3. Folder Structure (Standard Go Layout)
Saya minta kamu siapkan struktur folder ini agar tim lain (atau kamu di masa depan) tidak bingung:
```text
internal/
  domain/
    entity/      <-- Objek bisnis & Logika inti
    repository/  <-- Kontrak/Interface data
  usecase/       <-- Alur kerja (Orchestrator)
```

---

## Tugas Eksekusi Kamu Hari Ini:

Jangan tanya "Gimana kodenya?". Kamu harus cari tahu sendiri (atau tanya ke saya secara spesifik) tentang hal-hal berikut:

1.  **Gimana cara bikin Unit Test di Go?** Saya ingin kamu membuktikan logika `ReduceStock` kamu benar lewat file `_test.go`, bukan lewat `fmt.Println` di `main.go`.
2.  **Gimana cara implementasi Interface di memori?** Karena kita belum pakai Database, kamu harus bikin "In-Memory Repository" (pakai `Map` di dalam Go) yang mengimplementasikan interface repository yang sudah kamu buat.
3.  **Gimana cara menangani Error yang elegan?** Jangan cuma `return err`. Coba riset tentang `errors.Is` atau `errors.As` di Go.



---

**Laporan pertama yang saya tunggu dari kamu:**
"Manager, saya sudah buat folder `internal/domain/entity`. Saya punya kendala saat menentukan apakah `UpdatePrice` itu harus mengecek *history* harga atau langsung timpa saja di level Entity. Menurut Anda bagaimana?"

**Silakan eksekusi. Jangan takut salah, yang penting Go-mu jalan dan Test-mu Pass.** Ada pertanyaan terkait *Business Logic* dari POS ini sebelum kamu mulai ngetik?