package jwt

import (
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateNewToken func for generate a new token.
func GenerateNewToken(expiresIn time.Duration, privateKey, userID string) (string, error) {
	// Set secret key from .env file.
	secret := privateKey

	// Set expires minutes count for secret key from .env file.
	minutesCount, _ := strconv.Atoi(os.Getenv("JWT_SECRET_KEY_EXPIRE_MINUTES_COUNT"))

	// Create a new claims.
	claims := jwt.MapClaims{}

	now := time.Now().UTC()

	// Set public claims:
	claims["sub"] = userID
	claims["exp"] = now.Add(time.Minute * time.Duration(minutesCount)).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	t, err := token.SignedString([]byte(secret))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return t, nil
}
