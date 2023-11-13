package auth

import (
	"errors"
	"time"

	jwt "github.com/golang-jwt/jwt"
)

// JwtWrapper menampung signing key dan penerbitnya
// JwtWrapper adalah tipe data yang menampung secret key, penerbit, dan masa berlaku
type JwtWrapper struct {
	SecretKey         string //Kunci yang digunakan untuk signing JWT token
	Issuer            string //yang mengeluarkan JWT Token
	ExpirationMinutes int64  //Waktu dalam menit token jwt akan berlaku
	ExpirationHours   int64  //Waktu dalam jam token jwt akan berlaku
}

// JwtClaim menambahkan email sebagai bentuk klaim ke token
type JwtClaim struct {
	Email string
	jwt.StandardClaims
}

// GenerateToken membuat token jwt
func (j *JwtWrapper) GenerateToken(email string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return
}

// RefreshToken membuat penyegaran kembali token jwt
func (j *JwtWrapper) RefreshToken(email string) (signedToken string, err error) {
	claims := &JwtClaim{
		Email: email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Local().Add(time.Minute * time.Duration(j.ExpirationMinutes)).Unix(),
			Issuer:    j.Issuer,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString([]byte(j.SecretKey))
	if err != nil {
		return
	}
	return
}

// ValidateToken validates the jwt token
// ValidateToken takes a signed JWT token as an argument and returns the JwtClaim and an error
func (j *JwtWrapper) ValidateToken(signedToken string) (claims *JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return
	}
	claims, ok := token.Claims.(*JwtClaim)
	if !ok {
		err = errors.New("klaim tidak bisa diparsing")
		return
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		err = errors.New("JWT sudah kadaluarsa")
		return
	}
	return
}
