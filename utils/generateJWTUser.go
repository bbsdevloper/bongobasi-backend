package utils






import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GenerateToken(phone string) (string, error) {
	// Define the token claims, including the user ID
	claims := jwt.MapClaims{
		"phone": phone,
		"exp":    time.Now().Add(time.Hour * 24).Unix(), // Set token expiration time
	}

	// Create the JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Sign the token with the secret key
	tokenString, err := token.SignedString([]byte("your-secret-key"))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
