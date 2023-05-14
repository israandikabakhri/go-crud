package penduduk

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	Model "main.core/go-crud/app/Model/penduduk"
)

func Data(c *gin.Context) {

	var penduduk []Model.Penduduk

	Model.DB.Table("penduduk").Find(&penduduk)

	c.JSON(http.StatusOK, gin.H{"data": penduduk})

}

func Detail(c *gin.Context) {

	var penduduk Model.Penduduk
	nik := c.Param("nik")

	if err := Model.DB.Table("penduduk").First(&penduduk, nik).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Tidak Ditemukan"})
			return
		default:
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		}
	}

	c.JSON(http.StatusOK, gin.H{"data": penduduk})

}

func Search(c *gin.Context) {

	var penduduk []Model.Penduduk

	var input struct {
		Key string `json:"key"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	key := input.Key

	Model.DB.Table("penduduk").Where("nama LIKE ?", "%"+key+"%").Find(&penduduk)

	c.JSON(http.StatusOK, gin.H{"data": penduduk})

}

func Create(c *gin.Context) {
	var penduduk Model.Penduduk

	if err := c.ShouldBindJSON(&penduduk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := Model.DB.Table("penduduk").Create(&penduduk).Error; err != nil {
		if IsDuplicateError(err) {
			c.AbortWithStatusJSON(http.StatusConflict, gin.H{"message": "Data dengan NIK yang sama sudah ada"})
			return
		}

		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"msg": "Berhasil Menyimpan Data"})
}

func IsDuplicateError(err error) bool {
	var mysqlError *mysql.MySQLError
	if errors.As(err, &mysqlError) && mysqlError.Number == 1062 {
		return true
	}

	return false
}

func Update(c *gin.Context) {

	var penduduk Model.Penduduk
	nik := c.Param("nik")

	if err := c.ShouldBindJSON(&penduduk); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if Model.DB.Table("penduduk").Model(&penduduk).Where("nik = ?", nik).Updates(&penduduk).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Tidak dapat mengubah data penduduk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Data Berhasil Diperbaharui"})

}

func Delete(c *gin.Context) {

	var penduduk Model.Penduduk

	var input struct {
		Nik string `json:"nik"`
	}

	if err := c.ShouldBindJSON(&input); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	nik := input.Nik
	if Model.DB.Table("penduduk").Delete(&penduduk, nik).RowsAffected == 0 {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Tidak dapat menghapus data penduduk"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": "Data Berhasil Dihapus"})

}
