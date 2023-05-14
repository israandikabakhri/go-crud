package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"main.core/go-crud/app/Http/Controllers/penduduk"
	"main.core/go-crud/config/database"
	"main.core/go-crud/migration"
)

func init() {
	database.ConnectDatabase()
	migration.Migration()
}

func main() {

	r := gin.Default()

	// Menambahkan middleware CORS
	r.Use(cors.Default())

	r.GET("/api/penduduk", penduduk.Data)
	r.GET("/api/penduduk/:nik", penduduk.Detail)
	r.POST("/api/penduduk/search", penduduk.Search)
	r.POST("/api/penduduk", penduduk.Create)
	r.PUT("/api/penduduk/:nik", penduduk.Update)
	r.DELETE("/api/penduduk", penduduk.Delete)

	r.Run()

}
