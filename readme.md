# **RestAPI Educational**
&nbsp; jadi pada repository ini aku akan membagikan pengalamanku dalam membangun `RestAPI` dengan menggunakan bahasa `go`.

<br>

### *Depedency* :
1. [go-sql-driver/mysql](https://github.com/go-sql-driver/mysql)
2. [julienschmidt/httprouter](https://github.com/julienschmidt/httprouter)
3. [go-playground/validator](https://github.com/go-playground/validator)

<br>
<br>

### *Outline project* :
 1. [Setup **Project**](#setup-project)
 2. Membuat _Database_
 3. Membuat _CRUD_
 4. Membuat _Router Server_

<br>
<br>

## **Setup Project**
---
menginisialisasi project `go`.

```
go mod init <nama-project>

go mod init github.com/imniwa/go-rest-api-edu
```

menambahkan beberapa depedency.
```
go get -u https://github.com/go-sql-driver/mysql

go get https://github.com/julienschmidt/httprouter

go get https://github.com/go-playground/validator
```
> opsi -u menginstruksikan perintah `get` untuk mengupdate modul yang menyediakan dependency yang disebutkan pada baris perintah untuk menggunakan rilis minor atau patch yang lebih baru jika tersedia.


