package middleware

import (
	"ams-fantastic-auth/internal/config"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/response"
	"fmt"
	"strings"
	"time"

	"ams-fantastic-auth/pkg/jwttool"

	"github.com/gofiber/fiber/v2"
)

type JWTAuth struct {
	db          *database.Database
	accessToken *jwttool.Token
	refreshToen *jwttool.Token
}

func New(db *database.Database, cfgJWT *config.JWT) *JWTAuth {
	// new tokens
	return &JWTAuth{
		db:          db,
		accessToken: jwttool.New(cfgJWT.AccessTokenSecretKey, cfgJWT.AccessTokenPublicKey, time.Duration(cfgJWT.AccessTokenExpiredIn), cfgJWT.AccessTokenMaxage),
		refreshToen: jwttool.New(cfgJWT.RefreshTokenSecretKey, cfgJWT.RefreshTokenPublicKey, time.Duration(cfgJWT.RefreshTokenExpiredIn), cfgJWT.RefreshTokenMaxage),
	}
}

func (j *JWTAuth) AccessToken() *jwttool.Token {
	return j.accessToken
}

func (j *JWTAuth) RefreshToken() *jwttool.Token {
	return j.refreshToen
}

// BearerAuthReq Middleware to handle bearer token authentication
func (j *JWTAuth) BearerAuthReq(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		tokenString = c.Cookies("access_token")
	}

	if len(tokenString) <= 0 {
		return response.AuthError(c, fiber.StatusUnauthorized, response.AuthUnauthorizedErrorCode, "")
	}

	mapClaims, validErr := j.accessToken.ValidToken(tokenString)
	if validErr != nil {
		return response.AuthError(c, fiber.StatusUnauthorized, response.AuthInvalidTokenErrorCode, validErr.Error())
	}

	user, selectErr := database.SelectUserByEmail(j.db.Connect(), fmt.Sprint(mapClaims["sub"]))
	if selectErr != nil {
		return response.AuthError(c, fiber.StatusInternalServerError, response.AuthSelectFailErrorCode, selectErr.Error())
	}

	if user.Email != mapClaims["sub"] {
		return response.AuthError(c, fiber.StatusUnauthorized, response.AuthNotFoundOwnUserErrorCode, "")
	}

	c.Locals("user", *user)

	// Continue with the request chain if the token is valid
	return c.Next()
}
