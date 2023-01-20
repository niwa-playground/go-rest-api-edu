# **RestAPI Educational**
&nbsp; jadi pada repository ini aku akan membagikan pengalamanku dalam membangun `RestAPI` dengan menggunakan bahasa `go`.

<br>

### *Depedency* :
1. [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
2. [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
3. [go-playground/validator](https://github.com/go-playground/validator)

<br>

### Table of content :
- [**RestAPI Educational**](#restapi-educational)
    - [*Depedency* :](#depedency-)
    - [Table of content :](#table-of-content-)
- [**Setup Project**](#setup-project)
  - [menginisialisasi project `go`.](#menginisialisasi-project-go)
  - [menambahkan beberapa depedency.](#menambahkan-beberapa-depedency)
- [**Membuat Database**](#membuat-database)
  - [membuat database `test_golang_api`](#membuat-database-test_golang_api)
  - [membuat tabel `category`](#membuat-tabel-category)
- [**Membuat CRUD**](#membuat-crud)
  - [Membuat Entity \[`model/entity`\]](#membuat-entity-modelentity)
  - [Membuat Repository \[`repository/*`\]](#membuat-repository-repository)
    - [Implementasi `func Save()`](#implementasi-func-save)
    - [Menambahkan Helper untuk menghandle `panic`](#menambahkan-helper-untuk-menghandle-panic)

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

# **Membuat CRUD**
---


## Membuat Entity [`model/entity`]
---
kelas entity ini akan kita masukan ke dalam directory `model\entity` didalam folder ini. untuk setiap tabel kita perlu representasikan ke dalam data struct. 
```go
type Category struct {
    Id      int
    Name    string
}
```

<br>

## Membuat Repository [`repository/*`]
---
sebelum membuat repository untuk `category` kita perlu membuat `interface`-nya terlebih dahulu.
```go
type BaseCategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Update(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category
	Delete(ctx context.Context, tx *sql.Tx, category entity.Category)
	FindById(ctx context.Context, tx *sql.Tx, categoryId int) entity.Category
	FindAll(ctx context.Context, tx *sql.Tx) []entity.Category
}
```

setelah itu baru kita implementasikan `interface` yang sudah dibuat ke dalam `<nama-table>_repository.go`.

### Implementasi `func Save()`
```go
func (repository *CategoryRepository) Save(ctx context.Context, tx *sql.Tx, category entity.Category) entity.Category {
	SQL := "INSERT INTO category(name) values (?)"
	
    result, err := tx.ExecContext(ctx, SQL, category.Name)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
    helper.PanicIfError(err)

	category.Id = int(id)
	return category
}
```

### Menambahkan Helper untuk menghandle `panic`
fungsi ini berguna untuk mengecek jika terjadi error dalam pemanggilan fungsi. 

untuk directory file akan aku simpan ke dalam `helper/error.go`.
```go
func PanicIfError(err error) {
	if err != nil {
		panic(err)
	}
}
```