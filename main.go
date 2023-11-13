package main

import (
	"go-jwt/controllers"
	"go-jwt/database"
	"go-jwt/middlewares"
	"go-jwt/models"
	"log"

	"github.com/gin-gonic/gin"
)

// fungsi main pada program
// fungsi main menginisialisasi basis data, persiapan router, dan menyalakan server.
func main() {
	//inisialisasi basis data
	err := database.InitDatabase()
	if err != nil {
		// mencatatkan error dan mematikan server
		log.Fatalln("Tidak bisa membuat basis data", err)
	}

	//Migrasi otomatis model pengguna
	// Fungsi AutoMigrate() melakukan migrasi skema basis data dan memperbaharui basis data secara otomatis
	database.GlobalDB.AutoMigrate(&models.User{})

	//Menyiapkan router
	r := setupRouter()

	//menyalakan Server
	r.Run(":8080")
}

// setupRouter menyiapkan router dan menambahkan jalur router.
func setupRouter() *gin.Engine {
	//Membuat router baru
	r := gin.Default()

	//Membuat route untuk selamat datang
	r.GET("/", func(c *gin.Context) {
		c.String(200, "Selamat Datang!!")
	})

	//Membuat grup baru untuk API
	api := r.Group("/api")
	{
		//Membuat grup baru untuk public routes
		public := api.Group("/public")
		{
			//Menambahkan route untuk login
			public.POST("/login", controllers.Login)

			//Membuat route untuk pendaftaran
			public.POST("/signup", controllers.Signup)
		}

		//Menambahkan grup protected
		protected := api.Group("/protected").Use(middlewares.Authz())
		{
			//Menambahkan route signup
			protected.GET("/profile", controllers.Profile)
		}
	}

	//Mengembalikan router
	return r
}
