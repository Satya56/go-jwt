package middlewares

import (
	"go-jwt/auth"
	"strings"

	"github.com/gin-gonic/gin"
)

// Authz adalah middleware yang menvalidasi token dan memberikan otorisasi kepada pengguna
// Authz menggunakan gin.context sebagai argumen atau parameter dan mengembalikan gin.HandlerFunc
// Fungsi bertanggung jawab dalam menvalidasi token yang berada di Authorization header
// dan memberikan otorisasi kepada pengguna jika token pengguna valid
func Authz() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Mengambil Authorization header dari request
		clientToken := c.Request.Header.Get("Authorization")
		if clientToken == "" {
			//Jika Authorization kosong maka akan mengembalikan kode status 403
			c.JSON(403, "Authorization header kosong")
			c.Abort()
			return
		}
		//memecah authorization header untuk mendapatkan token
		extractedToken := strings.Split(clientToken, "Bearer ")
		if len(extractedToken) == 2 {
			//Trim the token
			clientToken = strings.TrimSpace(extractedToken[1])
		} else {
			//jika token tidak sesuai format yang disyaratkan, mengembalikan nilai status kode 400
			c.JSON(400, "Format Authorization Token tidak sesuai ketentuan")
			c.Abort()
			return
		}
		//Menginisialisasi JwtWrapper dengan secret key dan penerbit token
		jwtWrapper := auth.JwtWrapper{
			SecretKey: "verysecretkey",
			Issuer:    "AuthService",
		}
		//Validasi token
		claims, err := jwtWrapper.ValidateToken(clientToken)
		if err != nil {
			//jika token tidak valid, maka akan mengembalikan nilai kode status 401
			c.JSON(401, err.Error())
			c.Abort()
			return
		}
		//Menetapkan claims di dalam context
		c.Set("email", claims.Email)
		//Melanjutkan ke header selanjutnya
		c.Next()
	}
}
