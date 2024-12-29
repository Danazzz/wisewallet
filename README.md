**WiseWallet** adalah REST API berbasis Golang dengan framework Gin, menggunakan PostgreSQL sebagai database. API ini dirancang untuk melacak pengeluaran dan pemasukan pengguna, serta menambahkan elemen gamifikasi sederhana melalui pemberian badge.

---

#### **Fitur Utama**

---

**Authentication**

* **Register User** : Membuat akun pengguna baru.
* **Login User** : Login menggunakan username dan password.
* **Basic Auth** : Mengamankan sebagian besar endpoint API.

---

**Transaction Management**

1. Menambahkan, membaca, memperbarui, dan menghapus transaksi pengguna.
2. Mendukung tipe transaksi income (Pemasukan) dan expense (Pengeluaran)

---

**Gamifikasi**

**Badge Management** :

Membuat dan mengelola badge berdasarkan kriteria tertentu.

**User-Badge Relationship** :

Mengaitkan badge ke pengguna berdasarkan pencapaian dan menggambarkannya dengan JOIN.

---

**Migrations:**

**Automatic Migrations** : Migrasi table-table secara otomatis pada PostgreSQL.

**Teknologi yang Digunakan**

* **Bahasa Pemrograman** : Golang
* **Framework** : Gin
* **Database** : PostgreSQL
* **Deployment** : Railway
* **Middleware** : Basic Auth

---

#### **Struktur Folder**

---

wisewallet/

 │

├── controllers/			# Mengelola logic untuk setiap endpoint

├── database/			# Konfigurasi database dan migrasi

 │         ├── database.go

 │        └── migrations/	# File SQL untuk migrasi database

├── middlewares/		# Middleware untuk autentikasi dan keamanan

├── models/			# Model data (struct) untuk database

├── routes/				# Definisi routing API

├── main.go			# Entry point aplikasi

---

#### **API Endpoints**

---

**Authentication**


| Endpoint           | Method | Description            | Auth Required |
| ------------------ | ------ | ---------------------- | ------------- |
| /api/auth/register | POST   | Mendaftarkan pengguna  | No            |
| /api/auth/login    | POST   | Masuk sebagai Pengguna | No            |

**Transactions**

| Endpoint              | Method | Description                        | Auth Required |
| --------------------- | ------ | ---------------------------------- | ------------- |
| /api/transactions     | GET    | Dapatkan semua transaksi           | Yes           |
| /api/transactions     | POST   | Membuat transaksi baru             | Yes           |
| /api/transactions/:id | GET    | Dapatkan transaksi berdasarkan ID  | Yes           |
| /api/transactions/:id | PUT    | Perbarui transaksi berdasarkan ID  | Yes           |
| /api/transactions/:id | DELETE | Menghapus transaksi berdasarkan ID | Yes           |

**Badges**

| Endpoint        | Method | Description                | Auth Required |
| --------------- | ------ | -------------------------- | ------------- |
| /api/badges     | GET    | Dapatkan semua lencana     | Yes           |
| /api/badges     | POST   | Membuat lencana            | Yes           |
| /api/badges/:id | GET    | Dapatkan lencana dengan ID | Yes           |
| /api/badges/:id | PUT    | Perbarui lencana dengan ID | Yes           |
| /api/badges/:id | DELETE | Hapus lencana dengan ID    | Yes           |

**User-Badge Relationships**

| Endpoint             | Method | Description                                | Auth Required |
| -------------------- | ------ | ------------------------------------------ | ------------- |
| /api/user_badges     | GET    | Dapatkan semua pengguna dan lencana mereka | Yes           |
| /api/user_badges     | POST   | Membuat lencana untuk pengguna            | Yes           |
| /api/user_badges/:id | GET    | Dapatkan semua lencana dari pengguna       | Yes           |
| /api/user_badges/:id | DELETE | Menghapus lencana dari pengguna            | Yes           |

**Cara Menjalankan**

**1. Clone Repository**

git clone `<this-repository-url>`

cd wisewallet

**2. Install Dependencies**

go mod tidy

**3. Jalankan Server**

go run main.go

Database termigrasi otomatis

Server akan berjalan di port **8080**.
