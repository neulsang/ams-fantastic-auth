package server

import (
	"time"

	"ams-fantastic-auth/internal/config"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func New(cfg *config.Server) *fiber.App {
	app := fiber.New(fiber.Config{
		ReadTimeout: time.Second * time.Duration(cfg.ReadTimeout),
	})

	app.Use(
		// Add CORS to each route.
		cors.New(cors.Config{
			AllowOrigins: cfg.Origin,
		}),
		// Add simple logger.
		logger.New(),
	)
	return app
}
