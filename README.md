# 📚 SIS-PUSTAKA

Sistem backend sederhana untuk perpustakaan digital menggunakan bahasa pemrograman **Go (Golang)**.

---

## 📌 Important

Untuk memulai anda hanya perlu:

- npm install
- cd ke sis-pustaka
- run cmd/server/main.go

## 📌 Deskripsi

Proyek ini adalah implementasi sederhana dari REST API untuk sistem perpustakaan dengan fitur:

- 🔀 **Routing** menggunakan [Chi v5](https://github.com/go-chi/chi)
- 💾 **Penyimpanan data in-memory** (tanpa database)
- ✅ **Validasi input** saat membuat buku baru (tidak boleh kosong)
- 📜 **Custom middleware logger** untuk mencatat request masuk
- 🧱 **Struktur modular**: handler, model, storage, middleware, utils, router, dll.

---

## 📌 Catatan

- ✅ Semua endpoint telah diuji menggunakan **Postman**
- ⚠️ Error handler dengan response format **JSON standar** telah diterapkan
- 🧠 **Singleton pattern** digunakan untuk memastikan storage hanya ada satu instance

---

## 🧪 Endpoint yang Tersedia

- `GET    /sis-pustaka/` → Ambil semua buku
- `POST   /sis-pustaka/` → Tambah buku baru
- `GET    /sis-pustaka/{id}` → Ambil detail buku
- `PUT    /sis-pustaka/{id}` → Update data buku
- `DELETE /sis-pustaka/{id}` → Hapus buku

---

## 🧾 Terima kasih! 🚀
