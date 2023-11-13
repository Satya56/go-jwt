package database

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// GlobalDB adalah obyek global yang akan digunakan pada package yang berbeda
var GlobalDB *gorm.DB

// InitDatabase membuat inisialisasi koneksi ke database
func InitDatabase() (err error) {
	// Membaca berkas environment
	config, err := godotenv.Read()
	if err != nil {
		log.Fatal("Gagal membaca berkas.env")
	}
	// Membuat data source name dari berkas .env
	dsn := fmt.Sprintf(
		"%s:%s@(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config["DBUSER"],
		config["DBPASS"],
		config["DBHOST"],
		config["DBNAME"],
	)
	// Melakukan koneksi ke database lalu menyimpannya ke dalam variabel GlobalDB
	GlobalDB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	fmt.Print(dsn)
	if err != nil {
		return
	}
	return
}
