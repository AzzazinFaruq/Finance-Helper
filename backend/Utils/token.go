package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

func getJWTSecret() []byte {
	secret := os.Getenv("JWT_SECRET")
	if secret == "" {
		// Untuk development, sebaiknya jangan panic, tapi berikan default atau log error
		// Namun untuk produksi (klinik), wajib ada di environment
		return []byte("default_secret_key_klinik_123") 
	}
	return []byte(secret)
}

// GenerateJWT membuat token baru untuk user
func GenerateJWT(userID int64) (string, error) {
	// Untuk aplikasi klinik, kita bisa menambahkan role jika diperlukan
	claims := jwt.MapClaims{
		"sub": userID,
		"exp": time.Now().Add(time.Hour * 24).Unix(), // Token berlaku 24 jam
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	
	tokenString, err := token.SignedString(getJWTSecret())
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT memvalidasi string token dan mengembalikan objek token
func ValidateJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Pastikan metode signingnya HMAC (HS256)
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return getJWTSecret(), nil
	})

	if err != nil {
		return nil, err
	}

	return token, nil
}