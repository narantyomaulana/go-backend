# Go Backend API

Proyek backend API dengan Go, menggunakan Gin Framework dan GORM.

## Persyaratan

- Go versi 1.23.4 atau lebih tinggi
- MySQL database
- Git

## Cara Clone dan Setup Proyek

### 1. Clone Repositori

```bash
git clone https://github.com/narantyomaulana/go-backend.git
cd go-backend
```

### 2. Install Dependencies

Install semua dependencies yang diperlukan:

```bash
go mod tidy
```

### 3. Konfigurasi Database

1. Buat database MySQL baru bernama `db_golang`

```sql
CREATE DATABASE db_golang;
```

2. Konfigurasi `.env` (file sudah ada dalam repositori):

```
APP_PORT=3000
DB_HOST=localhost
DB_PORT=3306
DB_USER=root
DB_PASS=YourPassword
DB_NAME=db_golang
JWT_SECRET=YourSecretKey
```

Sesuaikan konfigurasi database (`DB_USER`, `DB_PASS`) dengan pengaturan MySQL lokal Anda.

### 4. Menjalankan Aplikasi

Jalankan aplikasi dengan perintah:

```bash
go run main.go
```

Server akan berjalan pada `http://localhost:3000`

## Teknologi yang Digunakan

- [Gin Framework](https://github.com/gin-gonic/gin) - Web framework
- [GORM](https://gorm.io/) - ORM library
- [JWT](https://github.com/golang-jwt/jwt) - JSON Web Token untuk autentikasi