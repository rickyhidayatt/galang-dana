## _Proyek Golang Rest API_

Proyek ini merupakan aplikasi web funding yang menggunakan bahasa pemrograman Golang untuk membuat API.
Beberapa teknologi yang digunakan dalam proyek ini antara lain:

- Golang
- PostgreSQL
- Gin Framework
- Midtrans

Pada API ini, juga sudah diterapkan sistem authorization (otorisasi) untuk mengamankan akses ke data dan fitur-fitur pada aplikasi.

## Fitur

- User Registrasi, Login & Logout
- Transaksi
- Buat Campaign

## Instalasi

- Siapkan file .env dalam folder cmd

DB_USER=
DB_PASS=
DB_HOST=
DB_PORT=
DB_NAME=

ServerKey = server key dari api key midtrans
ClientKey = client key dari api key midtrans

## Run Program

```sh
go run main.go
```

## Dokumentasi API

https://documenter.getpostman.com/view/23608652/2s93ecwW4p
