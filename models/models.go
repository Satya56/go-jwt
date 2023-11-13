package models

import (
	"go-jwt/database"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Struct user di dalam basis data
type User struct {
	gorm.Model
	ID       int    `gorm:"primaryKey"`
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required" gorm:"unique"`
	Password string `json:"password" binding:"required"`
}

// CreateUserRecord adalah fungsi untuk membuat user ke dalam basis data
// CreateUserRecord membuat data user dengan mengambil pointer User structs
func (user *User) CreateUserRecord() error {
	result := database.GlobalDB.Create(&user)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// HashPassword mengenkripsi password pengguna
// HashPassword menggunakan string sebagai parameter dan mengenkripsinya
func (user *User) HashPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return err
	}
	user.Password = string(bytes)
	return nil
}

// CheckPassword mengecek password pengguna
// CheckPassword menggunakan string sebagai parameter dan melakukan komparasi antara masukkan string dengan password pengguna yang terenkripsi di basis data
func (user *User) CheckPassword(providedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(providedPassword))
	if err != nil {
		return err
	}
	return nil
}
