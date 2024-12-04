package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}


func CreateToken(username string, duration time.Duration) (string,int64,error) {
	secretKey := os.Getenv("SECRET_KEY")
    // Create a new JWT token with claims
    expirationTime := time.Now().Add(duration).Unix()
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": username,                    // Subject (user identifier)
		"iss": "todo-app",                  // Issuer
		// "aud": getRole(username),           // Audience (user role)
		"exp": expirationTime, // Expiration time
		"iat": time.Now().Unix(),                 // Issued at
	})

	tokenString, err := claims.SignedString([]byte(secretKey))
    if err != nil {
        return "",0,err
    }

    // Print information about the created token
	fmt.Printf("Token claims added: %+v\n", claims)
	return tokenString,expirationTime,nil
}