package controllers

import (
	"go-jwt/database"
	"go-jwt/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

//Profile adalah fungsi untuk mengambil profil pengguna
//Mengembalikan nilai 404 jika pengguna tidak ditemukan
//Mengembalikan nilai 500 jika terjadi error ketika mencari pengguna di basis data

func Profile(c *gin.Context) {
	//Menginisialisasi model pengguna
	var user models.User

	//Mengambil email dari authorization middleware
	email, _ := c.Get("email")

	//Mencari pengguna dari basis data berdasarkan data email
	result := database.GlobalDB.Where("email = ?", email.(string)).First(&user)

	//jika pengguna tidak ditemukan maka fungsi akan mengembalikan nilai kode status 404
	if result.Error == gorm.ErrRecordNotFound {
		c.JSON(404, gin.H{
			"Error": "Pengguna tidak ditemukan",
		})
		c.Abort()
		return
	}

	// jika terjadi error ketika mencari data pengguna di dalam basis data
	if result.Error != nil {
		c.JSON(500, gin.H{
			"Error": "tidak dapat mengambil data",
		})
		c.Abort()
		return
	}

	//menetapkan password pengguna menjadi string kosong
	user.Password = ""

	//Mengembalikan profil pengguna dengan kode status 200
	c.JSON(200, user)
}
