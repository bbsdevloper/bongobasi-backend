package middleware

import (
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request header or query parameter
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			tokenString = r.URL.Query().Get("token")
		}
		secretKey := os.Getenv("JWT_SECRET")

		// Parse and validate the token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(secretKey), nil
		})
		if err != nil || !token.Valid {
			// Return unauthorized status if token is invalid
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		// Proceed to the next handler if token is valid
		next.ServeHTTP(w, r)
	})
}
