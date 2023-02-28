package middleware

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func Fiber(a *fiber.App) {
	a.Use(
		// Add CORS to each route.
		cors.New(cors.Config{
			AllowOrigins:     "*",
			AllowHeaders:     "Origin, Content-Type, Accept",
			AllowMethods:     "GET, POST, DELETE, PATCH",
			AllowCredentials: true,
		}),
		// Add simple logger.
		logger.New(),
	)
}
