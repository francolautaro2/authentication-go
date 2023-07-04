package token

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/joho/godotenv"
)

// Struct of claims for jwt token
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Get key secret in .env file
func GetKeySecret() []byte {
	godotenv.Load()
	var secretEnv = os.Getenv("SECRET_KEY")
	var secreteKey = []byte(secretEnv)
	return secreteKey
}

// Generate token function
func GenerateToken(username string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(GetKeySecret())
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Middleware for validate token
func ValidateToken(h http.HandlerFunc) http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		// Get cookie
		c, err := r.Cookie("Token")
		if err != nil {
			fmt.Fprintf(w, "Error on get cookie value")
		}

		// Get token value of cookie
		tokenString := c.Value

		// Validate value of cookie
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(tkn *jwt.Token) (interface{}, error) {
			return GetKeySecret(), nil
		})

		// Errors handlers on jwt validate processing
		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				http.Error(w, "Token de autorización inválido", http.StatusUnauthorized)
				return
			} else {
				http.Error(w, "Session expired, please log in again", http.StatusBadRequest)
			}
			return
		}

		// Verify if jwt token is valid
		if !token.Valid {
			http.Error(w, "Token de autorización inválido", http.StatusUnauthorized)
			return
		}

		// return for other handler
		h(w, r)
	}

}
