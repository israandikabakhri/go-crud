# go-crud
CRUD menggunakan framework GIN dan Library GORM


Dalam project ini melibatkan 3 file yaitu:
```
app/Http/Controller/penduduk/penduduk.go
app/Model/penduduk.go
main.go
```





Catatan Penggunaan Tiap Folder:
- app/Function -> Fungsinya sebagai penyimpanan query kompleks atau fungsi query yang memiliki relasi yang kompleks yang nanti sisa dipanggil oleh controller
- app/Helpers -> Fungsinya sebagai penyimpanan file helper seperti mengubah format tanggal, waktu, uang dan lain-lain
- app/Http/Controllers -> Fungsinya sebagai penyimpanan seluruh controller
- app/Http/Requests -> Fungsinya sebagai penyimpanan semua validator requests yang nanti akan dipanggil di controller untuk validasi input data seperti simpan dan update
- app/Model -> Fungsinya sebagai penyimpanan folder-folder Model, Connect Database serta migration

Catatan Batasan Penggunaan Komponen:
- Model: Komponen yang berinteraksi dengan database dan pengubah format data seperti mengubah format tanggal yang dipanggil dari Helper
- Helper: Komponen fungsi yang dapat dipanggil dari mana saja
- Controller: Komponen yang mengelola pemanggilan data dari model, pemanggilan validasi requests, pemanggilan Functions Query Kompleks dan Pengelolaan Rest API
- Request: Komponen yang mengelola validasi input dari user termasuk keamanan SQL Injection dan XSS Script
- Main.go: Komponen yang mengelola Routing dan Pengelolaan hak akses serta authentikasi menggunakan JWT auth

Catatan Setting data: