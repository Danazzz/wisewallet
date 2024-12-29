**WiseWallet** adalah REST API berbasis Golang dengan framework Gin, menggunakan PostgreSQL sebagai database. API ini dirancang untuk melacak pengeluaran dan pemasukan pengguna, serta menambahkan elemen gamifikasi sederhana melalui pemberian badge.

**Fitur Utama**

**Authentication**

* **Register User** : Membuat akun pengguna baru.
* **Login User** : Login menggunakan username dan password.
* **Basic Auth** : Mengamankan sebagian besar endpoint API.

**Transaction Management**

* Menambahkan, membaca, memperbarui, dan menghapus transaksi pengguna.
* Mendukung tipe transaksi:
* income: Pemasukan
* expense: Pengeluaran

**Gamifikasi**

**Badge Management** :

Membuat dan mengelola badge berdasarkan kriteria tertentu.

**User-Badge Relationship** :

Mengaitkan badge ke pengguna berdasarkan pencapaian.

**Teknologi yang Digunakan**

* **Bahasa Pemrograman** : Golang
* **Framework** : Gin
* **Database** : PostgreSQL
* **Deployment** : Railway
* **Middleware** : Basic Auth

**Struktur Folder**

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

**API Endpoints**

**Authentication**

**Endpoint**			**Method**		**Description**		**Auth Required**

**/api/auth/register**		POST		**Register a user**	No

**/api/auth/login**		POST		**Login a user**		No

**Transactions**

**Endpoint**			**Method**		**Description**				**Auth Required**

**/api/transactions**		GET			**Get all transactions**		Yes

**/api/transactions**		POST		**Create a new transaction**	Yes

**/api/transactions/:id**	GET			**Get transaction by ID**		Yes

**/api/transactions/:id**	PUT			**Update transaction by ID**	Yes

**/api/transactions/:id**	DELETE		**Delete transaction by ID**	Yes

**Badges**

**Endpoint**		**Method**	**Description**			**Auth Required**

**/api/badges**		GET		**Get all badges**		Yes

**/api/badges**		POST	**Create a badge**		Yes

**/api/badges/:id**	GET		**Get badge by ID**		Yes

**/api/badges/:id**	PUT		**Update badge by ID**	Yes

**/api/badges/:id**	DELETE	**Delete badge by ID**	Yes

**User-Badge Relationships**

**Endpoint**			**Method**	**Description**						**Auth Required**

**/api/user_badges**		GET		**Get all badges for a user**			Yes

**/api/user_badges**		POST	**Assign a badge to a user**			Yes

**/api/user_badges/:id**	GET		**Get user-badge relationship by ID**	Yes

**/api/user_badges/:id**	DELETE	**Remove a badge from a user**		Yes

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
