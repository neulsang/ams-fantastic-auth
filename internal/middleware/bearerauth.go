package middleware

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"
	"ams-fantastic-auth/internal/response"
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

// BearerAuthReq Middleware to handle bearer token authentication
func BearerAuthReq(c *fiber.Ctx) error {
	var tokenString string
	authorization := c.Get("Authorization")

	if strings.HasPrefix(authorization, "Bearer ") {
		tokenString = strings.TrimPrefix(authorization, "Bearer ")
	} else if c.Cookies("access_token") != "" {
		tokenString = c.Cookies("access_token")
	}

	if len(tokenString) <= 0 {
		return response.NewError(c, fiber.StatusUnauthorized, "You are not logged in")
	}

	accessJwtSercret := "AmsAccessJwtSecret"
	tokenByte, parseErr := jwt.Parse(tokenString, func(jwtToken *jwt.Token) (interface{}, error) {
		if _, ok := jwtToken.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", jwtToken.Header["alg"])
		}

		return []byte(accessJwtSercret), nil
	})

	if parseErr != nil {
		return response.NewError(c, fiber.StatusUnauthorized, fmt.Sprintf("invalidate token(%v)", parseErr))
	}

	claims, ok := tokenByte.Claims.(jwt.MapClaims)
	if !ok || !tokenByte.Valid {
		return response.NewError(c, fiber.StatusUnauthorized, "invalid token claim")
	}

	db, newErr := database.New(configs.Database())
	if newErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, newErr.Error())
	}

	user, selectErr := database.SelectUserByEmail(db, fmt.Sprint(claims["sub"]))
	if selectErr != nil {
		return response.NewError(c, fiber.StatusInternalServerError, selectErr.Error())
	}

	if user.Email != claims["sub"] {
		return response.NewError(c, fiber.StatusUnauthorized, "the user belonging to this token no logger exists")
	}

	c.Locals("user", *user)

	// Continue with the request chain if the token is valid
	return c.Next()
}
