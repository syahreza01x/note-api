# Note API - Aplikasi Catatan Sederhana dengan Go

Ini adalah aplikasi backend sederhana buat menyimpan dan mengelola catatan (notes) menggunakan bahasa Go.  
Kamu bisa tambah, lihat, ubah, dan hapus catatan lewat API yang mudah dipakai.

---

## Fitur

- Tambah, lihat, update, dan hapus catatan lewat REST API  
- Response JSON yang rapi dan mudah dibaca  
- Ada halaman web sederhana buat lihat semua catatan dalam bentuk list  
- Data disimpan di memori (sementara), jadi data hilang kalau server dimatikan  

---

## Cara Pakai

1. Jalankan Server
    - go run main.go
2. Buka browser atau pakai Postman/curl untuk akses API:
    - Lihat semua catatan (JSON): http://localhost:8000/notes
    - Lihat semua catatan (web page): http://localhost:8000/notes/view