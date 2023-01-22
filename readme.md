# **Simple RestAPI [EDU]**
jadi pada repository ini aku akan membagikan pengalamanku dalam membangun `RestAPI` dengan menggunakan bahasa `go`.

<br>

### *Depedency* :
1. [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
2. [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
3. [go-playground/validator](https://github.com/go-playground/validator)

<br>

### Table of content :
- [**Simple RestAPI \[EDU\]**](#simple-restapi-edu)
    - [*Depedency* :](#depedency-)
    - [Table of content :](#table-of-content-)
- [**Setup Project**](#setup-project)
  - [menginisialisasi project `go`.](#menginisialisasi-project-go)
  - [menambahkan beberapa depedency.](#menambahkan-beberapa-depedency)
- [**Membuat Database**](#membuat-database)
  - [membuat database `test_golang_api`](#membuat-database-test_golang_api)
  - [membuat tabel `category`](#membuat-tabel-category)
- [**Struktur Project**](#struktur-project)
- [**Penjelasan**](#penjelasan)

<br>

# **Setup Project**
---
## menginisialisasi project `go`.

```
go mod init <nama-project>
go mod init github.com/imniwa/go-rest-api-edu
```

## menambahkan beberapa depedency.
```
go get -u github.com/go-sql-driver/mysql
go get github.com/julienschmidt/httprouter
go get github.com/go-playground/validator
```
> opsi -u menginstruksikan perintah `get` untuk mengupdate modul yang menyediakan dependency yang disebutkan pada baris perintah untuk menggunakan rilis minor atau patch yang lebih baru jika tersedia.


# **Membuat Database**
---
## membuat database `test_golang_api`
```sql
CREATE DATABASE test_golang_api 
```

## membuat tabel `category`
```sql
CREATE TABLE category(
    id integer primary key auto_increment,
    name varchar(200) not null
)engine=InnoDB
```

# **Struktur Project**
struktur project kurang lebih sebagai berikut :
1. [`app/database.go`]  berisikan file untuk membuat koneksi dengan database.
2. [`controller/controller.go`] berisikan `interface` untuk setiap model.
3. [`controller/*`] berisikan file untuk controller yang mengatur setiap model,berdasarkan `interface` yang sudah dibuat.
4. [`exception/*`] berisikan exception untuk manghandle error.
5. [`helper/*`] berisikan kumpulan function - function tambahan yang perlu diimplementasikan di setiap file.
6. [`middleware/*`] berisikan middleware untuk mengatur authentication.
7. [`model/*`] berisikan 3 jenis model yaitu :
   - `entity`  digunakan untuk membuat model setiap tabel.
   - `request` digunakan untuk membuat `struct` request untuk `http`.
   - `response` digunakan untuk membuat `struct` response untuk `http`.
8. [`repositories/*`] berisikan file untuk melakukan SQL query ke database.
9. [`services/*`] berisikan file untuk memanggil repository.

# **Penjelasan**
1. `repositories` file ini digunakan untuk melakukan memanggilan Query ke database dengan cara mengeksekusi Context dari libary context.
2. `services` file ini digunakan untuk melakukan validasi terhadap request dan melakukan commit terhadap query yang dieksekusi pada `repositories` jika query berhasil dilakukan dan jika gagal maka akan dilakukan rollback.
3. `model entity` file ini digunakan untuk membuat aturan untuk penggunaan tabel di database sebagai model.