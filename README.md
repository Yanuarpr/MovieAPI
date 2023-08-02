# MovieAPI

# Movie API dengan Pola MVP dan MySQL Database

Aplikasi Movie API sederhana yang dibangun dengan bahasa pemrograman Golang menggunakan pola arsitektur Model-View-Presenter (MVP) dan database MySQL. API ini memungkinkan pengguna untuk melakukan operasi CRUD (Create, Read, Update, Delete) pada data film.

## Teknologi dan Pustaka yang Digunakan

- Golang (Go): Bahasa pemrograman yang digunakan untuk mengembangkan aplikasi ini.
- Echo: Framework web untuk membuat endpoint API HTTP.
- Gorm: Library ORM untuk berinteraksi dengan database MySQL.

## Instalasi dan Penggunaan

1. Pastikan Anda memiliki Go (versi minimal 1.16) terinstal di komputer Anda. Anda dapat mengunduh Go dari https://golang.org/dl/ dan mengikuti petunjuk instalasinya.

2. Pastikan MySQL telah terinstal di komputer Anda dan sudah berjalan. Buatlah sebuah database dengan nama yang Anda inginkan (misalnya, "movie_db") untuk menyimpan data film.

3. Clone repositori ini ke direktori lokal Anda:

```bash
git clone https://github.com/username/repo-name.git

cd repo-name

go get -u github.com/labstack/echo
go get -u gorm.io/gorm
go get -u gorm.io/driver/mysql
```

# Konfigurasi database:

Buka file main.go dan temukan bagian // Inisialisasi koneksi database MySQL.
Ganti nilai dsn dengan informasi koneksi database MySQL Anda. Contoh: username:password@tcp(localhost:3306)/movie_db?charset=utf8mb4&parseTime=True&loc=Local.
Jalankan aplikasi dengan perintah:

go run main.go

1. API Movie sekarang berjalan di http://localhost:8080. Anda dapat mengakses API ini dengan menggunakan klien HTTP seperti Postman atau mengintegrasikannya dengan aplikasi frontend Anda.

# Endpoint API
Berikut adalah daftar endpoint API yang tersedia:

GET /movies: Mengambil daftar semua film.

POST /movie: Menambahkan film baru.

PUT /movie/:id: Memperbarui data film berdasarkan ID.

DELETE /movie/:id: Menghapus film berdasarkan ID.

# Kontribusi
Kontribusi selalu dipersilakan! Jika Anda menemukan masalah, bug, atau memiliki perbaikan atau fitur baru untuk ditambahkan, silakan buat issue atau pull request.
