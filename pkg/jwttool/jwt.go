package jwttool

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

type Token struct {
	secretKey string
	publicKey string
	ExpiredIn time.Duration
	Maxage    int
}

func New(secretKey, publcKey string, expiredIn time.Duration, maxage int) *Token {
	return &Token{
		secretKey: secretKey,
		publicKey: publcKey,
		ExpiredIn: expiredIn,
		Maxage:    maxage,
	}
}

// GenerateNewToken func for generate a new token.
func (t *Token) GenerateNewToken(userEmail string) (string, error) {
	// Create a new claims.
	claims := jwt.MapClaims{}
	now := time.Now().UTC()

	// Set public claims:
	claims["sub"] = userEmail
	claims["exp"] = now.Add(t.ExpiredIn * time.Minute).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	// Create a new JWT access token with claims.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate token.
	newToken, err := token.SignedString([]byte(t.secretKey))
	if err != nil {
		// Return error, it JWT token generation failed.
		return "", err
	}

	return newToken, nil
}

func (j *Token) ValidToken(tokenString string) (jwt.MapClaims, error) {
	tokenByte, parseErr := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(j.secretKey), nil
	})

	if parseErr != nil {
		return nil, fmt.Errorf("invalidate token(%v)", parseErr)
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return nil, errors.New("invalid token claim")
	}
	return claims, nil
}
