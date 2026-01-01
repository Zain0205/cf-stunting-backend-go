
# CF Stunting Backend (Go + Fiber)

Backend API untuk sistem **Certainty Factor (CF) Stunting** berbasis **rule-based**, dibangun menggunakan **Go (Fiber)** dan **MySQL**, serta sudah **fully dockerized**.

Project ini dirancang sebagai **API-only backend** untuk:

* Autentikasi JWT
* Role berbasis kategori pengguna
* Manajemen pertanyaan & rule CF
* Proses diagnosis stunting berbasis Certainty Factor

---

## 🚀 Tech Stack

* **Go** `1.25.5`
* **Fiber** (Web Framework)
* **GORM** (ORM)
* **MySQL 8**
* **JWT Authentication**
* **Docker & Docker Compose**

---

## 👤 Kategori / Role Pengguna

Setiap user **hanya memiliki satu kategori**, dan kategori ini menentukan:

* daftar pertanyaan
* rule certainty factor

Kategori:

1. **Perempuan Prakonsepsi**
2. **Perempuan Pernah Melahirkan**
3. **Remaja 19 Tahun**

Role akan disimpan di JWT dan digunakan untuk membatasi akses pertanyaan.

---

## 📂 Struktur Folder

```
backend-go/
├── cmd/
│   └── main.go            # Entry point aplikasi
├── internal/
│   ├── auth/              # JWT, login, register
│   ├── cf/                # Engine Certainty Factor
│   ├── config/            # Loader env
│   ├── database/          # MySQL connection
│   ├── diagnosis/         # Logic diagnosis
│   ├── middleware/        # Auth & role middleware
│   ├── models/            # GORM models
│   └── seed/              # Seeder data awal
├── .env.example
├── Dockerfile
├── docker-compose.yml
├── go.mod
└── go.sum
```

---

## ⚙️ Environment Variable

Copy file `.env.example` menjadi `.env`:

```bash
cp .env.example .env
```

Isi minimal:

```env
APP_PORT=8080
APP_ENV=development

DB_USER=root
DB_PASSWORD=secret
DB_HOST=mysql
DB_PORT=3306
DB_NAME=cf_stunting

JWT_SECRET=supersecret
```

---

## 🐳 Menjalankan dengan Docker

### 1️⃣ Build image

```bash
docker compose build --no-cache
```

### 2️⃣ Jalankan container

```bash
docker compose up
```

### 3️⃣ Test API

```bash
curl http://localhost:8080
```

Response:

```json
{
  "status": "CF Stunting API running"
}
```

---

## 🗄️ Database

* Database akan **auto-created** saat container MySQL pertama kali jalan
* Tabel akan dibuat otomatis oleh **GORM AutoMigrate**
* Seeder digunakan untuk data awal (kategori, domain, pertanyaan)

---

## 🔐 Authentication

* Menggunakan **JWT**
* Token dikirim via header:

```http
Authorization: Bearer <token>
```

* Role/kategori user disimpan di token

---

## 📌 Roadmap Implementasi

* [x] Docker & MySQL setup
* [x] Auto migrate database
* [ ] Auth JWT (register & login)
* [ ] Seeder pertanyaan & rule CF
* [ ] CF engine
* [ ] Endpoint diagnosis
* [ ] Endpoint history diagnosis

---

## 🧠 Catatan Penting

* `depends_on` **tidak menunggu MySQL siap**, oleh karena itu backend menggunakan retry DB
* Jangan gunakan `localhost` untuk DB host di Docker, gunakan `mysql`
* Project ini **API only**, belum ada UI admin

---

## ✍️ Author

Dikembangkan untuk penelitian sistem diagnosis stunting berbasis **Certainty Factor**.

---

