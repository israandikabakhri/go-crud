# go-crud
CRUD menggunakan framework GIN dan Library GORM


## Dalam project ini melibatkan 3 file yaitu:

```
go-crud/app/Http/Controller/penduduk/penduduk.go
go-crud/app/Model/penduduk.go
go-crud/main.go
```

## Cara Menjalankan Project:

1. Buka **CMD**
2. Masuk ke Folder project melalui **CMD**
3. ketik **go run main.go** 


## Collection Postman

Untuk mempermudah terting data saya telah menyiapkan di folder ini:
```
go-crud/postman-collection/Go - CRUD.postman_collection.json
```
Silahkan import kedalam postman anda

Enjoyy..

Jika ada kendala jangan sungkan email ke andikaisra7@gmail.com


## Catatan Penggunaan Tiap Folder:

- app/Function -> Fungsinya sebagai penyimpanan query kompleks atau fungsi query yang memiliki relasi yang kompleks yang nanti sisa dipanggil oleh controller
- app/Helpers -> Fungsinya sebagai penyimpanan file helper seperti mengubah format tanggal, waktu, uang dan lain-lain
- app/Http/Controllers -> Fungsinya sebagai penyimpanan seluruh controller
- app/Http/Requests -> Fungsinya sebagai penyimpanan semua validator requests yang nanti akan dipanggil di controller untuk validasi input data seperti simpan dan update
- app/Model -> Fungsinya sebagai penyimpanan folder-folder Model, Connect Database serta migration

## Catatan Batasan Penggunaan Komponen:

- Model: Komponen yang berinteraksi dengan database dan pengubah format data seperti mengubah format tanggal yang dipanggil dari Helper
- Helper: Komponen fungsi yang dapat dipanggil dari mana saja
- Controller: Komponen yang mengelola pemanggilan data dari model, pemanggilan validasi requests, pemanggilan Functions Query Kompleks dan Pengelolaan Rest API
- Request: Komponen yang mengelola validasi input dari user termasuk keamanan SQL Injection dan XSS Script
- Main.go: Komponen yang mengelola Routing dan Pengelolaan hak akses serta authentikasi menggunakan JWT auth
