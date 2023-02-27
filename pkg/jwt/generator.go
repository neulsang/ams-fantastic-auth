package jwt

import (
	"time"

	"github.com/golang-jwt/jwt"
)

// GenerateNewToken func for generate a new token.
func GenerateNewToken(expiresIn time.Duration, privateKey, userEmail string) (string, error) {
	// Set secret key from .env file.
	secret := privateKey

	// Set expires minutes count for secret key from .env file.
	minutesCount := expiresIn

	// Create a new claims.
	claims := jwt.MapClaims{}

	now := time.Now().UTC()

	// Set public claims:
	claims["sub"] = userEmail
	claims["exp"] = now.Add(minutesCount).Unix()
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
