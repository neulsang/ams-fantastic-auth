package middleware

import (
	"ams-fantastic-auth/internal/configs"
	"ams-fantastic-auth/internal/database"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/basicauth"
)

// BasicAuthReq middleware
func BasicAuthReq() func(*fiber.Ctx) error {
	cfg := basicauth.Config{
		Users: map[string]string{
			"admin": "test001", // TODO: fixed admin 계정?!
		},
	}
	err := basicauth.New(cfg)
	return err
}

// BasicAuthExtReq middleware
func BasicAuthExtReq() func(*fiber.Ctx) error {
	cfg := basicauth.Config{
		Authorizer: func(id, pass string) bool {
			db, newErr := database.New(configs.Database())
			if newErr != nil {
				return false
			}

			user, selectErr := database.SelectUser(db, id)
			if selectErr != nil {
				return false
			}

			if user == nil {
				return false
			}
			if id == user.ID && pass == user.Password {
				return true
			}
			return false
		},
	}
	err := basicauth.New(cfg)
	return err
}
