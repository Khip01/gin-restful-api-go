package product_controller

import (
	"encoding/json"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khip01/gin-restfulapi-go/models"
	"gorm.io/gorm"
)

// GET atau menampilkan semua data
func Index(context *gin.Context) {
	// Membuat Slice dari struct Product di folder models
	var products []models.Product

	// Memanggilnya di database
	models.DB.Find(&products)

	// Response JSON (Status dan Map Products)
	context.JSON(http.StatusOK, gin.H{"products": products})
}

// GET atau menampilkan detail sebuah data
func Show(context *gin.Context) {
	// Membuat Struct product baru dari models struct Product
	var product models.Product

	// Membuat variabel untuk menampung parameter dari context url
	id := context.Param("id")

	// Membuat variable handler jika error
	checkError := models.DB.First(&product, id).Error
	// Kondisi jika terdapat error
	if checkError != nil {
		// switch case tipe error, apakah dia NOT FOUND / SERVER ERROR
		switch checkError {
		// Jika error karena data tidak ditemukan/Not Found 404
		case gorm.ErrRecordNotFound:
			// Response JSON (Status dan Map Message nya)
			context.AbortWithStatusJSON(http.StatusNotFound, gin.H{"message": "Data Not Found 404"})
			return
		// Jika error karena Internal Server Error
		default:
			// Response JSON (Status dan Map Message nya)
			context.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"Message": checkError.Error()})
			return
		}
	}

	// Jika tidak error, mengembalikan data yang ada di database
	context.JSON(http.StatusOK, gin.H{"product": product})

}

// POST atau membuat sebuah data baru
func Create(context *gin.Context) {
	// Membuat Struct product baru dari models struct Product
	var product models.Product

	// Membuat variabel handler jika error, jika tidak error maka akan menyimpan ke struct product diatas
	checkError := context.ShouldBindJSON(&product)
	// Kondisi jika error
	if checkError != nil {
		// Response JSON (Status dan message nya)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": checkError.Error()})
		return
	}

	// Jika tidak ada error maka simpan data input ke database
	models.DB.Create(&product)
	// Response JSON (status dan Map product yang telah diinput)
	context.JSON(http.StatusOK, gin.H{"product": product})
}

// PUT atau update data yang dipilih
func Update(context *gin.Context) {
	// Membuat Struct product baru dari models struct Product
	var product models.Product

	// Membuat var id untuk parameter yang direquest
	id := context.Param("id")

	// Membuat error handler jika terdapat error, jika tidak maka akan menyimpan di struct product
	errorCheck := context.ShouldBindJSON(&product)
	// Kondisi jika terdapat error pada saat menyimpan
	if errorCheck != nil {
		// Response jika error (Status Bad Request dan Message Error)
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": errorCheck.Error()})
		return
	}

	// Membuat variable error handler jika error, sekalian mengupdate data ke database
	checkRowAffected := models.DB.Model(&product).Where("id = ?", id).Updates(&product).RowsAffected
	if checkRowAffected == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	// Jika data berhasil ditambahkan maka memberikan Response JSON (Status dan message)
	context.JSON(http.StatusOK, gin.H{"message": "Sukses, data berhasil diupdate!"})
}

// DELETE atau menghapus data yang dipilih
func Delete(context *gin.Context) {
	// Membuat variabel struct product baru dari models struct Product
	var product models.Product

	// Cara TRADISIONAL
	// Membuat Map untuk tempat menyimpan Input ID dari data yang akan dihapus
	// input := map[string]string{"id": "0"}

	// Cara AMAN
	// Mebuat Struct untuk tempat menyimpan input ID dari data yang akan dihapus
	var input struct {
		Id json.Number
	}

	// Check error pada saat menyimpan id ke input
	checkError := context.ShouldBindJSON(&input)
	if checkError != nil {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": checkError.Error()})
		return
	}

	// Cara TRADISIONAL menggunakan parsing jika data string
	// Sebelum menghapus data menggunakan input ID tertentu, maka input ID harus dikonversi ke Int terlebih dahulu
	/* Parseint berisi (
	string = input["id"]
	base   = 10 (int itu base 10)
	byte   = 64 (mengikuti tipedata int di models struct Product))
	*/
	// id, _ := strconv.ParseInt(input["id"], 10, 64)

	// Cara AMAN menggunakan json number
	id, _ := input.Id.Int64()

	// Membuat variabel error handler jika proses hapus error
	checkErrorDelete := models.DB.Delete(&product, id).RowsAffected
	// Kondisi jika tidak ada data yang diupdate karena data tidak tersedia
	if checkErrorDelete == 0 {
		context.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Data tidak ditemukan"})
		return
	}

	// Jika data berhasil dihapus
	context.JSON(http.StatusOK, gin.H{"message": "Sukses, data berhasil dihapus!"})
}
